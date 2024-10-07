package relay

import (
	"bytes"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type GetTaskResultInput struct {
	ImageNum string `json:"image_num"`
	TaskId   uint64 `json:"task_id"`
}

type UploadTaskParamsInput struct {
	TaskArgs string `json:"task_args"`
	TaskId   uint64 `json:"task_id"`
}

type UploadTaskPramsWithSignature struct {
	UploadTaskParamsInput
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

type UploadResultInput struct {
	TaskId uint64 `form:"task_id" json:"task_id"`
}

func UploadTask(task *models.InferenceTask) error {

	appConfig := config.GetConfig()

	params := &UploadTaskParamsInput{
		TaskArgs: task.TaskArgs,
		TaskId:   task.TaskId,
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return err
	}

	paramsWithSig := &UploadTaskPramsWithSignature{
		UploadTaskParamsInput: *params,
		Timestamp:             timestamp,
		Signature:             signature,
	}

	postJson, err := json.Marshal(paramsWithSig)
	if err != nil {
		return err
	}

	body := bytes.NewReader(postJson)
	reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks"

	r, _ := http.NewRequest("POST", reqUrl, body)
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Duration(3) * time.Second,
	}

	response, err := client.Do(r)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {

		responseBytes, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}

		return errors.New("upload task params error: " + string(responseBytes))
	}

	return nil
}

func DownloadTaskResult(task *models.InferenceTask) error {

	appConfig := config.GetConfig()

	taskFolder := path.Join(
		appConfig.DataDir.InferenceTasks,
		strconv.FormatUint(uint64(task.ID), 10))

	if err := os.MkdirAll(taskFolder, 0700); err != nil {
		return err
	}

	taskIdStr := strconv.FormatUint(task.TaskId, 10)

	var numImages int
	if task.TaskType == models.TaskTypeSD {
		var err error
		numImages, err = models.GetTaskConfigNumImages(task.TaskArgs)
		if err != nil {
			return err
		}
	} else {
		numImages = 1
	}

	var fileExt string
	if task.TaskType == models.TaskTypeSD {
		fileExt = ".png"
	} else {
		fileExt = ".json"
	}

	startTime := time.Now()
	for i := numImages - 1; i >= 0; i-- {
		iStr := strconv.Itoa(i)

		getResultInput := &GetTaskResultInput{
			ImageNum: strconv.Itoa(i),
			TaskId:   task.TaskId,
		}

		timestamp, signature, err := SignData(getResultInput, appConfig.Blockchain.Account.PrivateKey)
		if err != nil {
			return err
		}

		timestampStr := strconv.FormatInt(timestamp, 10)

		queryStr := "?timestamp=" + timestampStr + "&signature=" + signature
		reqUrl := appConfig.Relay.BaseURL + "/v1/inference_tasks/" + taskIdStr + "/results/" + iStr
		reqUrl = reqUrl + queryStr

		filename := path.Join(taskFolder, iStr+fileExt)

		log.Debugln("Downloading result: " + reqUrl)

		resp, err := http.Get(reqUrl)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			respBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			return errors.New(string(respBytes))
		}

		file, err := os.Create(filename)
		if err != nil {
			if err := resp.Body.Close(); err != nil {
				return err
			}

			return err
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			if err := resp.Body.Close(); err != nil {
				return err
			}

			if err := file.Close(); err != nil {
				return err
			}

			return err
		}

		if err := resp.Body.Close(); err != nil {
			return err
		}

		if err := file.Close(); err != nil {
			return err
		}

	}
	endTime := time.Now()
	timeCost := endTime.Sub(startTime).Seconds()
	log.Infof("RelayGetResult: time cost %f seconds", timeCost)

	log.Debugln("All results downloaded!")

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
