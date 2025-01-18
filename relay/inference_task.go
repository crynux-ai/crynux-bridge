package relay

import (
	"context"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"errors"
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
	Index            int64  `json:"index"`
	TaskIDCommitment string `json:"task_id_commitment"`
}

type UploadTaskParamsInput struct {
	TaskArgs         string `json:"task_args"`
	TaskIDCommitment string `json:"task_id_commitment"`
}

type UploadResultInput struct {
	TaskId uint64 `form:"task_id" json:"task_id"`
}

func UploadTask(ctx context.Context, task *models.InferenceTask) error {

	appConfig := config.GetConfig()

	params := &UploadTaskParamsInput{
		TaskArgs:         task.TaskArgs,
		TaskIDCommitment: task.TaskIDCommitment,
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("task_id_commitment", task.TaskIDCommitment)
	form.Add("task_args", task.TaskArgs)
	form.Add("timestamp", strconv.FormatInt(timestamp, 10))
	form.Add("signature", signature)
	body := strings.NewReader(form.Encode())

	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/" + task.TaskIDCommitment

	r, _ := http.NewRequestWithContext(ctx, "POST", reqUrl, body)
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}

	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {

		responseBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New("upload task params error: " + string(responseBytes))
	}

	log.Infof("Relay: upload task params of %s", task.TaskIDCommitment)
	return nil
}

func DownloadTaskResult(ctx context.Context, task *models.InferenceTask, index int64, dst io.Writer) error {
	appConfig := config.GetConfig()
	getResultInput := &GetTaskResultInput{
		Index:            index,
		TaskIDCommitment: task.TaskIDCommitment,
	}

	timestamp, signature, err := SignData(getResultInput, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	query := url.Values{}
	query.Add("timestamp", strconv.FormatInt(timestamp, 10))
	query.Add("signature", signature)
	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/" + task.TaskIDCommitment + "/results/" + strconv.FormatInt(index, 10)
	req, _ := http.NewRequestWithContext(ctx, "GET", reqUrl, nil)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(respBytes))
	}

	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		return err
	}
	log.Infof("Relay: get result %d of task %s", index, task.TaskIDCommitment)

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
