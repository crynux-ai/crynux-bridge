package relay

import (
	"context"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crynux_bridge/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	ethParams "github.com/ethereum/go-ethereum/params"
)

type GetTaskByCommitmentInput struct {
	TaskIDCommitment string `json:"task_id_commitment"`
}

type CreateTaskInput struct {
	TaskIDCommitment string   `json:"task_id_commitment"`
	MinVram          uint64   `json:"min_vram"`
	Nonce            string   `json:"nonce"`
	RequiredGpu      string   `json:"required_gpu"`
	RequiredGpuVram  uint64   `json:"required_gpu_vram"`
	TaskArgs         string   `json:"task_args"`
	TaskModelIds     []string `json:"task_model_ids"`
	TaskSize         uint64   `json:"task_size"`
	TaskType         int      `json:"task_type"`
	TaskVersion      string   `json:"task_version"`

	TaskFee uint64 `json:"task_fee"`
}

type ValidateTaskInput struct {
	PublicKey         string   `json:"public_key"`
	TaskID            string   `json:"task_id"`
	TaskIDCommitments []string `json:"task_id_commitments"`
	VrfProof          string   `json:"vrf_proof"`
}

type CancelTaskInput struct {
	TaskIDCommitment string `json:"task_id_commitment"`
	AbortReason      int    `json:"abort_reason"`
}

type CheckBalanceInput struct {
	PublicKey string `json:"public_key"`
}

// Parse the "data" field of relay response, and store it in parsedData
func parseRelayResponseData(resp *http.Response, parsedData any) error {
	if resp.StatusCode != 200 {
		return errors.New("response status code is not 200")
	}

	// read response body
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// parse JSON
	content := make(map[string]interface{})
	if err := json.Unmarshal(respBytes, &content); err != nil {
		return err
	}

	// get "data" field
	data, ok := content["data"]
	if !ok {
		return errors.New("data field is missing")
	}

	// check parsedData, make sure it is a non-nil pointer
	val := reflect.ValueOf(parsedData)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		return errors.New("parsedData must be a non-nil pointer")
	}

	// parse "data" field according to its type
	switch v := data.(type) {
	case string:
		if strPtr, ok := parsedData.(*string); ok {
			*strPtr = v
		} else {
			return errors.New("parsedData must be *string for string data")
		}

	case float64:
		if strPtr, ok := parsedData.(*string); ok {
			*strPtr = fmt.Sprintf("%.0f", v)
		} else {
			return errors.New("parsedData must be *string for numeric data")
		}

	case map[string]interface{}:
		// target type must be a pointer to a struct
		jsonBytes, err := json.Marshal(v) // convert map to JSON
		if err != nil {
			return err
		}
		if err := json.Unmarshal(jsonBytes, parsedData); err != nil {
			return err
		}

	default:
		return errors.New("data field has an unsupported type")
	}

	return nil
}

func GetTaskByCommitment(ctx context.Context, taskIDCommitment string) (*models.RelayTask, error) {
	appConfig := config.GetConfig()

	params := &GetTaskByCommitmentInput{
		TaskIDCommitment: taskIDCommitment,
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return nil, err
	}

	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/" + taskIDCommitment

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	req, _ := http.NewRequestWithContext(timeoutCtx, "GET", reqUrl, nil)
	query := req.URL.Query()
	query.Add("timestamp", strconv.FormatInt(timestamp, 10))
	query.Add("signature", signature)
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay: get task by taskIDCommitment %s error: %v", taskIDCommitment, err)
		return nil, err
	}

	relayTask := new(models.RelayTask)
	err = parseRelayResponseData(resp, relayTask)
	if err != nil {
		log.Errorf("Relay: get task by taskIDCommitment %s error: %v", taskIDCommitment, err)
		return nil, err
	}

	log.Infof("Relay: get task %s", taskIDCommitment)
	return relayTask, nil
}

func CreateTask(ctx context.Context, task *models.InferenceTask) error {
	appConfig := config.GetConfig()

	params := &CreateTaskInput{
		TaskIDCommitment: task.TaskIDCommitment,
		MinVram:          task.MinVram,
		Nonce:            task.Nonce,
		RequiredGpu:      task.RequiredGPU,
		RequiredGpuVram:  task.RequiredGPUVram,
		TaskArgs:         task.TaskArgs,
		TaskModelIds:     task.TaskModelIDs,
		TaskSize:         task.TaskSize,
		TaskType:         int(task.TaskType),
		TaskVersion:      task.TaskVersion,
		TaskFee:          task.TaskFee,
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("min_vram", strconv.FormatUint(task.MinVram, 10))
	form.Add("nonce", task.Nonce)
	form.Add("required_gpu", task.RequiredGPU)
	form.Add("required_gpu_vram", strconv.FormatUint(task.RequiredGPUVram, 10))
	form.Add("task_args", task.TaskArgs)
	form.Add("task_model_ids", strings.Join(task.TaskModelIDs, ","))
	form.Add("task_size", strconv.FormatUint(task.TaskSize, 10))
	form.Add("task_type", strconv.Itoa(int(task.TaskType)))
	form.Add("task_version", task.TaskVersion)
	form.Add("task_fee", strconv.FormatUint(task.TaskFee, 10))
	form.Add("timestamp", strconv.FormatInt(timestamp, 10))
	form.Add("signature", signature)
	body := strings.NewReader(form.Encode())

	taskIDCommitment := task.TaskIDCommitment
	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/" + taskIDCommitment

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(timeoutCtx, "POST", reqUrl, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay: create task %s error: %v", taskIDCommitment, err)
		return err
	}

	log.Infof("Relay: create task %s", taskIDCommitment)
	return nil
}

// No matter single task or task group
// func ValidateSingleTask() and ValidateTaskGroup()
func ValidateTask(ctx context.Context, tasks []*models.InferenceTask) error {
	appConfig := config.GetConfig()

	taskIDCommitments := make([]string, len(tasks))
	taskID := tasks[0].TaskID
	vrfProof := tasks[0].VRFProof
	for i, task := range tasks {
		taskIDCommitments[i] = task.TaskIDCommitment
		if task.TaskID != taskID {
			return errors.New("task ID mismatch")
		}
		if task.VRFProof != vrfProof {
			return errors.New("VRF proof mismatch")
		}
	}

	privateKey := appConfig.Blockchain.Account.PrivateKey
	publicKey, err := utils.GetPubKeyFromPrivKey(privateKey)
	if err != nil {
		log.Errorf("Relay: validate task %s error: %v", taskIDCommitments, err)
		return err
	}

	params := &ValidateTaskInput{
		PublicKey:         publicKey,
		TaskID:            taskID,
		TaskIDCommitments: taskIDCommitments,
		VrfProof:          vrfProof,
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("public_key", publicKey)
	form.Add("task_id", taskID)
	form.Add("task_id_commitments", strings.Join(taskIDCommitments, ","))
	form.Add("vrf_proof", vrfProof)
	form.Add("timestamp", strconv.FormatInt(timestamp, 10))
	form.Add("signature", signature)
	body := strings.NewReader(form.Encode())

	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/validate"

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(timeoutCtx, "POST", reqUrl, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay: validate task %s error: %v", taskIDCommitments, err)
		return err
	}

	log.Infof("Relay: validate task %s success", taskIDCommitments)
	return nil
}

func CancelTask(ctx context.Context, task *models.InferenceTask, abortReason models.TaskAbortReason) error {

	appConfig := config.GetConfig()

	taskIDCommitment := task.TaskIDCommitment

	params := &CancelTaskInput{
		TaskIDCommitment: taskIDCommitment,
		AbortReason:      int(abortReason),
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("task_id_commitment", taskIDCommitment)
	form.Add("abort_reason", strconv.Itoa(int(abortReason)))
	form.Add("timestamp", strconv.FormatInt(timestamp, 10))
	form.Add("signature", signature)
	body := strings.NewReader(form.Encode())

	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/cancel"

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(timeoutCtx, "POST", reqUrl, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay: cancel task %s error: %v", taskIDCommitment, err)
		return err
	}

	log.Infof("Relay: cancel task %s success", taskIDCommitment)
	return nil
}

func CheckBalanceForTaskCreator(ctx context.Context) error {

	appConfig := config.GetConfig()

	privateKey := appConfig.Blockchain.Account.PrivateKey
	publicKey, err := utils.GetPubKeyFromPrivKey(privateKey)
	log.Infof("Relay: CheckBalanceForTaskCreator, publicKey: %s", publicKey)
	if err != nil {
		log.Errorf("Relay: CheckBalanceForTaskCreator error: %v", err)
		return err
	}

	params := &CheckBalanceInput{
		PublicKey: publicKey,
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("public_key", publicKey)
	form.Add("timestamp", strconv.FormatInt(timestamp, 10))
	form.Add("signature", signature)
	body := strings.NewReader(form.Encode())

	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/balanceOf"

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(timeoutCtx, "POST", reqUrl, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay: CheckBalanceForTaskCreator error: %v", err)
		return err
	}

	// get and check balance
	balanceStr := new(string)
	err = parseRelayResponseData(resp, balanceStr)
	if err != nil {
		log.Errorf("Relay: CheckBalanceForTaskCreator error: %v", err)
	}
	log.Infof("Relay: CheckBalanceForTaskCreator, balance: %s", *balanceStr)

	balance, ok := new(big.Int).SetString(*balanceStr, 10)
	if !ok {
		return errors.New("failed to convert balance string to big.Int")
	}

	ethThreshold := new(big.Int).Mul(big.NewInt(500), big.NewInt(ethParams.Ether))
	if balance.Cmp(ethThreshold) != 1 {
		return errors.New("not enough ETH left")
	}

	return nil
}
