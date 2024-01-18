package models

import (
	"crynux_bridge/config"
	"encoding/json"
	"net/url"

	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
)

var sdInferenceTaskSchema *jsonschema.Schema

func ValidateTaskArgsJsonStr(jsonStr string) (validationError, err error) {

	if sdInferenceTaskSchema == nil {
		schemaJson := config.GetConfig().TaskSchema.StableDiffusionInference

		if !isValidUrl(schemaJson) {
			sdInferenceTaskSchema, err = jsonschema.CompileString("stable_diffusion_task.json", schemaJson)
		} else {
			sdInferenceTaskSchema, err = jsonschema.Compile(schemaJson)
		}

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

func GetTaskConfigNumImages(taskArgs string) (int, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return 0, err
	}

	num := taskArgsMap["task_config"].(map[string]interface{})["num_images"].(float64)
	return int(num), nil
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
