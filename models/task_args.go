package models

import (
	"crynux_bridge/config"
	"encoding/json"
	"errors"
	"net/url"
	"reflect"

	"github.com/santhosh-tekuri/jsonschema/v5"
	_ "github.com/santhosh-tekuri/jsonschema/v5/httploader"
)

var sdInferenceTaskSchema *jsonschema.Schema
var gptInferenceTaskSchema *jsonschema.Schema
var sdFinetuneLoraTaskSchema *jsonschema.Schema

// IsNil checks if an interface{} value is nil
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Chan, reflect.Func:
		return rv.IsNil()
	default:
		return false
	}
}

func ValidateTaskArgsJsonStr(jsonStr string, taskType ChainTaskType) (validationError, err error) {
	if taskType == TaskTypeSD {
		return validateSDTaskArgs(jsonStr)
	} else if taskType == TaskTypeLLM {
		return validateGPTTaskArgs(jsonStr)
	} else { 
		return validateSDFinetuneLoraTaskArgs(jsonStr)
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

func validateSDFinetuneLoraTaskArgs(jsonStr string) (validationError, err error) {
	if sdFinetuneLoraTaskSchema == nil {
		schemaJson := config.GetConfig().TaskSchema.StableDiffusionFinetuneLora

		if !isValidUrl(schemaJson) {
			return nil, errors.New("invalid URL for task json schema")
		}

		sdFinetuneLoraTaskSchema, err = jsonschema.Compile(schemaJson)

		if err != nil {
			return nil, err
		}
	}

	var v interface{}
	if err := json.Unmarshal([]byte(jsonStr), &v); err != nil {
		return nil, err
	}
	return sdFinetuneLoraTaskSchema.Validate(v), nil
}

func GetTaskConfigNumImages(taskArgs string) (int, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return 0, err
	}

	taskConfig, ok := taskArgsMap["task_config"]
	if !ok || IsNil(taskConfig) {
		return 0, errors.New("task_config is missing or null")
	}

	taskConfigMap, ok := taskConfig.(map[string]interface{})
	if !ok {
		return 0, errors.New("task_config is not an object")
	}

	numImages, ok := taskConfigMap["num_images"]
	if !ok || IsNil(numImages) {
		return 0, errors.New("num_images is missing or null")
	}

	num, ok := numImages.(float64)
	if !ok {
		return 0, errors.New("num_images is not a number")
	}

	return int(num), nil
}

func GetSDTaskConfigBaseModel(taskArgs string) (string, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return "", err
	}

	baseModelValue, ok := taskArgsMap["base_model"]
	if !ok || IsNil(baseModelValue) {
		return "", errors.New("base_model is missing or null")
	}

	var baseModel string
	switch v := baseModelValue.(type) {
	case string:
		baseModel = v
	case map[string]interface{}:
		nameValue, ok := v["name"]
		if !ok || IsNil(nameValue) {
			return "", errors.New("base_model.name is missing or null")
		}
		name, ok := nameValue.(string)
		if !ok {
			return "", errors.New("base_model.name is not a string")
		}
		baseModel = name
	default:
		return "", errors.New("base_model has invalid type")
	}
	return baseModel, nil
}

func getSDTaskConfigModelIDs(taskArgs string) ([]string, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return nil, err
	}

	modelIDs := make([]string, 0)

	baseModelValue, ok := taskArgsMap["base_model"]
	if !ok || IsNil(baseModelValue) {
		return nil, errors.New("base_model is missing or null")
	}

	var baseModel string
	switch v := baseModelValue.(type) {
	case string:
		baseModel = v
	case map[string]interface{}:
		nameValue, ok := v["name"]
		if !ok || IsNil(nameValue) {
			return nil, errors.New("base_model.name is missing or null")
		}
		name, ok := nameValue.(string)
		if !ok {
			return nil, errors.New("base_model.name is not a string")
		}
		baseModel = name
		if originVariant, ok := v["variant"]; ok && !IsNil(originVariant) {
			variant, ok1 := originVariant.(string)
			if ok1 {
				baseModel = baseModel + "+" + variant
			}
		}
	default:
		return nil, errors.New("base_model has invalid type")
	}
	modelIDs = append(modelIDs, "base:"+baseModel)

	if loraValue, ok := taskArgsMap["lora"]; ok && !IsNil(loraValue) {
		lora, ok := loraValue.(map[string]interface{})
		if !ok {
			return nil, errors.New("lora is not an object")
		}
		modelValue, ok := lora["model"]
		if !ok || IsNil(modelValue) {
			return nil, errors.New("lora.model is missing or null")
		}
		model, ok := modelValue.(string)
		if !ok {
			return nil, errors.New("lora.model is not a string")
		}
		modelIDs = append(modelIDs, "lora:"+model)
	}

	if controlnetValue, ok := taskArgsMap["controlnet"]; ok && !IsNil(controlnetValue) {
		controlnet, ok := controlnetValue.(map[string]interface{})
		if !ok {
			return nil, errors.New("controlnet is not an object")
		}
		modelValue, ok := controlnet["model"]
		if !ok || IsNil(modelValue) {
			return nil, errors.New("controlnet.model is missing or null")
		}
		controlnetModel, ok := modelValue.(string)
		if !ok {
			return nil, errors.New("controlnet.model is not a string")
		}
		if originVariant, ok := controlnet["variant"]; ok && !IsNil(originVariant) {
			variant, ok1 := originVariant.(string)
			if ok1 {
				controlnetModel = controlnetModel + "+" + variant
			}
		}
		modelIDs = append(modelIDs, "controlnet:"+controlnetModel)
	}

	return modelIDs, nil
}

func getGPTTaskConfigModelIDs(taskArgs string) ([]string, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return nil, err
	}

	modelValue, ok := taskArgsMap["model"]
	if !ok || IsNil(modelValue) {
		return nil, errors.New("model is missing or null")
	}
	model, ok := modelValue.(string)
	if !ok {
		return nil, errors.New("model is not a string")
	}
	modelIDs := []string{"base:" + model}
	return modelIDs, nil
}

func getSDFTTaskConfigModelIDs(taskArgs string) ([]string, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return nil, err
	}

	modelValue, ok := taskArgsMap["model"]
	if !ok || IsNil(modelValue) {
		return nil, errors.New("model is missing or null")
	}

	var baseModel string
	model, ok := modelValue.(map[string]interface{})
	if !ok {
		return nil, errors.New("model is not an object")
	}

	nameValue, ok := model["name"]
	if !ok || IsNil(nameValue) {
		return nil, errors.New("model.name is missing or null")
	}
	name, ok := nameValue.(string)
	if !ok {
		return nil, errors.New("model.name is not a string")
	}
	baseModel = name
	if originVariant, ok := model["variant"]; ok && !IsNil(originVariant) {
		variant, ok1 := originVariant.(string)
		if ok1 {
			baseModel = baseModel + "+" + variant
		}
	}
	modelIDs := []string{"base:" + baseModel}
	return modelIDs, nil
}

func GetTaskConfigModelIDs(taskArgs string, taskType ChainTaskType) ([]string, error) {
	if taskType == TaskTypeSD {
		return getSDTaskConfigModelIDs(taskArgs)
	} else if taskType == TaskTypeLLM {
		return getGPTTaskConfigModelIDs(taskArgs)
	} else {
		return getSDFTTaskConfigModelIDs(taskArgs)
	}
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
