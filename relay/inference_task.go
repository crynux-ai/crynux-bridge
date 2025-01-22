package relay

import (
	"context"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type GetTaskResultInput struct {
	Index            string `json:"index"`
	TaskIDCommitment string `json:"task_id_commitment"`
}

type UploadTaskParamsInput struct {
	TaskArgs         string `json:"task_args"`
	TaskIDCommitment string `json:"task_id_commitment"`
}

type UploadResultInput struct {
	TaskId uint64 `form:"task_id" json:"task_id"`
}

type RelayError struct {
	StatusCode   int
	Method       string
	URL          string
	ErrorMessage string
}

func (e RelayError) Error() string {
	return fmt.Sprintf("RelayError: %s %s error code %d, %s", e.Method, e.URL, e.StatusCode, e.ErrorMessage)
}

func processRelayResponse(resp *http.Response) error {
	method := resp.Request.Method
	url := resp.Request.URL.RequestURI()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	} else if resp.StatusCode == 400 {
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		content := make(map[string]interface{})
		if err := json.Unmarshal(respBytes, &content); err != nil {
			return err
		}
		if data, ok := content["data"]; ok {
			msgBytes, err := json.Marshal(data)
			if err != nil {
				return err
			}
			msg := string(msgBytes)
			return RelayError{
				StatusCode:   resp.StatusCode,
				Method:       method,
				URL:          url,
				ErrorMessage: msg,
			}
		}
		if message, ok := content["message"]; ok {
			if msg, ok1 := message.(string); ok1 {
				return RelayError{
					StatusCode:   resp.StatusCode,
					Method:       method,
					URL:          url,
					ErrorMessage: msg,
				}
			}
		}
		return RelayError{
			StatusCode:   resp.StatusCode,
			Method:       method,
			URL:          url,
			ErrorMessage: string(respBytes),
		}
	} else {
		return RelayError{
			StatusCode:   resp.StatusCode,
			Method:       method,
			URL:          url,
			ErrorMessage: resp.Status,
		}
	}
}

func UploadTask(ctx context.Context, taskIDCommitment, taskArgs string) error {

	appConfig := config.GetConfig()

	params := &UploadTaskParamsInput{
		TaskArgs:         taskArgs,
		TaskIDCommitment: taskIDCommitment,
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("task_id_commitment", taskIDCommitment)
	form.Add("task_args", taskArgs)
	form.Add("timestamp", strconv.FormatInt(timestamp, 10))
	form.Add("signature", signature)
	body := strings.NewReader(form.Encode())

	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/" + taskIDCommitment

	ctx1, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	r, _ := http.NewRequestWithContext(ctx1, "POST", reqUrl, body)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay, upload task params of %s error: %v", taskIDCommitment, err)
		return err
	}

	log.Infof("Relay: upload task params of %s", taskIDCommitment)
	return nil
}

func DownloadTaskResult(ctx context.Context, taskIDCommitment string, index uint64, dst io.Writer) error {
	appConfig := config.GetConfig()
	getResultInput := &GetTaskResultInput{
		Index:            strconv.FormatUint(index, 10),
		TaskIDCommitment: taskIDCommitment,
	}

	timestamp, signature, err := SignData(getResultInput, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	ctx1, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/" + taskIDCommitment + "/results/" + strconv.FormatUint(index, 10)
	req, _ := http.NewRequestWithContext(ctx1, "GET", reqUrl, nil)
	query := req.URL.Query()
	query.Add("timestamp", strconv.FormatInt(timestamp, 10))
	query.Add("signature", signature)
	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay, get result of %s error: %v", taskIDCommitment, err)
		return err
	}

	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		return err
	}
	log.Infof("Relay: get result %d of task %s", index, taskIDCommitment)

	return nil
}

func UploadTaskResult(taskId uint64, taskType models.ChainTaskType, resultFiles []io.Reader) error {

	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)

	appConfig := config.GetConfig()

	uploadResultInput := &UploadResultInput{
		TaskId: taskId,
	}

	timestamp, signature, err := SignData(uploadResultInput, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	go func() {
		log.Debugln("writing form fields in go routine...")

		err = prepareUploadResultForm(resultFiles, taskType, writer, timestamp, signature)
		if err != nil {
			log.Errorln("error preparing the result uploading form")
			log.Errorln(err)

			if err2 := pw.CloseWithError(err); err2 != nil {
				log.Errorln("error closing the pipe")
				log.Errorln(err2)
			}
		}

		if err = writer.Close(); err != nil {
			log.Errorln("error closing the multipart form writer")
			log.Errorln(err)
		}

		if err = pw.Close(); err != nil {
			log.Errorln("error closing the pipe writer")
			log.Errorln(err)
		}

		log.Debugln("writing form fields completed")
	}()

	return callUploadResultApi(taskId, writer, pr)
}

func callUploadResultApi(taskId uint64, writer *multipart.Writer, body io.Reader) error {
	taskIdStr := strconv.FormatUint(taskId, 10)

	appConfig := config.GetConfig()

	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/" + taskIdStr + "/results"

	req, err := http.NewRequest("POST", reqUrl, body)
	if err != nil {
		log.Errorln("error creating upload result request")
		return err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := http.Client{}

	log.Debugln("uploading results...")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln("error making upload result request")
		log.Errorln(err)

		var urlErr *url.Error
		errors.As(err, &urlErr)

		return urlErr.Unwrap()
	}

	log.Debugln("upload result api finished")

	if resp.StatusCode != 200 {
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(respBytes))
	}

	return nil
}

func prepareUploadResultForm(
	resultFiles []io.Reader,
	taskType models.ChainTaskType,
	writer *multipart.Writer,
	timestamp int64,
	signature string) error {
	timestampStr := strconv.FormatInt(timestamp, 10)

	if err := writer.WriteField("timestamp", timestampStr); err != nil {
		log.Errorln("error writing timestamp fields to multipart form")
		return err
	}
	if err := writer.WriteField("signature", signature); err != nil {
		log.Errorln("error writing signature fields to multipart form")
		return err
	}

	var fileExt string
	if taskType == models.TaskTypeSD {
		fileExt = ".png"
	} else {
		fileExt = ".json"
	}

	for i := 0; i < len(resultFiles); i++ {
		part, err := writer.CreateFormFile("images", "image_"+strconv.Itoa(i)+fileExt)
		if err != nil {
			log.Errorln("error creating form file field " + strconv.Itoa(i))
			return err
		}

		if _, err := io.Copy(part, resultFiles[i]); err != nil {
			log.Errorln("error copying image to the form field " + strconv.Itoa(i))
			return err
		}
	}

	return nil
}
