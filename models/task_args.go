package models

import (
	"crynux_bridge/config"
	"encoding/json"
	"errors"
	"net/url"

	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
)

var sdInferenceTaskSchema *jsonschema.Schema
var gptInferenceTaskSchema *jsonschema.Schema

func ValidateTaskArgsJsonStr(jsonStr string, taskType ChainTaskType) (validationError, err error) {
	if taskType == TaskTypeSD {
		return validateSDTaskArgs(jsonStr)
	} else {
		return validateGPTTaskArgs(jsonStr)
	}
}

func validateSDTaskArgs(jsonStr string) (validationError, err error) {
	if sdInferenceTaskSchema == nil {
		schemaJson := config.GetConfig().TaskSchema.StableDiffusionInference

		if !isValidUrl(schemaJson) {
			return nil, errors.New("invalid URL for task json schema")
		}

		sdInferenceTaskSchema, err = jsonschema.Compile(schemaJson)

		if err != nil {
			return nil, err
		}
	}

	var v interface{}
	if err := json.Unmarshal([]byte(jsonStr), &v); err != nil {
		return nil, err
	}

	return sdInferenceTaskSchema.Validate(v), nil
}

func validateGPTTaskArgs(jsonStr string) (validationError, err error) {
	if gptInferenceTaskSchema == nil {
		schemaJson := config.GetConfig().TaskSchema.GPTInference

		if !isValidUrl(schemaJson) {
			return nil, errors.New("invalid URL for task json schema")
		}

		gptInferenceTaskSchema, err = jsonschema.Compile(schemaJson)

		if err != nil {
			return nil, err
		}
	}

	var v interface{}
	if err := json.Unmarshal([]byte(jsonStr), &v); err != nil {
		return nil, err
	}

	return gptInferenceTaskSchema.Validate(v), nil
}

func GetTaskConfigNumImages(taskArgs string) (int, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return 0, err
	}

	num := taskArgsMap["task_config"].(map[string]interface{})["num_images"].(float64)
	return int(num), nil
}

func GetTaskConfigBaseModel(taskArgs string) (string, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return "", nil
	}

	baseModel := taskArgsMap["base_model"].(string)
	return baseModel, nil
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
