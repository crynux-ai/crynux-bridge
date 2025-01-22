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
var sdFinetuneLoraTaskSchema *jsonschema.Schema

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

	num := taskArgsMap["task_config"].(map[string]interface{})["num_images"].(float64)
	return int(num), nil
}

func GetSDTaskConfigBaseModel(taskArgs string) (string, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return "", nil
	}

	var baseModel string
	switch v := taskArgsMap["base_model"].(type) {
	case string:
		baseModel = v
	case map[string]interface{}:
		name, ok := v["name"].(string)
		if !ok {
			return "", errors.New("sd task config is invalid")
		}
		baseModel = name
	default:
		return "", errors.New("sd task config is invalid")
	}
	return baseModel, nil
}

func getSDTaskConfigModelIDs(taskArgs string) ([]string, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return nil, nil
	}

	modelIDs := make([]string, 0)

	var baseModel string
	switch v := taskArgsMap["base_model"].(type) {
	case string:
		baseModel = v
	case map[string]interface{}:
		_, ok := v["name"]
		if !ok {
			return nil, errors.New("sd task config is invalid")
		}
		name, ok := v["name"].(string)
		if !ok {
			return nil, errors.New("sd task config is invalid")
		}
		baseModel = name
		if originVariant, ok := v["variant"]; ok {
			variant, ok1 := originVariant.(string)
			if ok1 {
				baseModel = baseModel + "+" + variant
			}
		}
	default:
		return nil, errors.New("sd task config is invalid")
	}
	modelIDs = append(modelIDs, "base:"+baseModel)

	if _, ok := taskArgsMap["lora"]; ok {
		lora, ok := taskArgsMap["lora"].(map[string]interface{})
		if !ok {
			return nil, errors.New("sd task config is invalid")
		}
		_, ok = lora["model"]
		if !ok {
			return nil, errors.New("sd task config is invalid")
		}
		model, ok := lora["model"].(string)
		if !ok {
			return nil, errors.New("sd task config is invalid")
		}
		modelIDs = append(modelIDs, "lora:"+model)
	}

	if _, ok := taskArgsMap["controlnet"]; ok {
		controlnet, ok := taskArgsMap["controlnet"].(map[string]interface{})
		if !ok {
			return nil, errors.New("sd task config is invalid")
		}
		_, ok = controlnet["model"]
		if !ok {
			return nil, errors.New("sd task config is invalid")
		}
		controlnetModel, ok := controlnet["model"].(string)
		if !ok {
			return nil, errors.New("sd task config is invalid")
		}
		if originVariant, ok := controlnet["variant"]; ok {
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
		return nil, nil
	}

	_, ok := taskArgsMap["model"]
	if !ok {
		return nil, errors.New("gpt task config is invalid")
	}
	model, ok := taskArgsMap["model"].(string)
	if !ok {
		return nil, errors.New("gpt task config is invalid")
	}
	modelIDs := []string{"base:" + model}
	return modelIDs, nil
}

func getSDFTTaskConfigModelIDs(taskArgs string) ([]string, error) {
	var taskArgsMap map[string]interface{}

	err := json.Unmarshal([]byte(taskArgs), &taskArgsMap)
	if err != nil {
		return nil, nil
	}

	_, ok := taskArgsMap["model"]
	if !ok {
		return nil, errors.New("sd finetune task config is invalid")
	}
	var baseModel string
	model, ok := taskArgsMap["model"].(map[string]interface{})
	if !ok {
		return nil, errors.New("sd finetune task config is invalid")
	}
	_, ok = model["name"]
	if !ok {
		return nil, errors.New("sd finetune task config is invalid")
	}
	name, ok := model["name"].(string)
	if !ok {
		return nil, errors.New("sd task config is invalid")
	}
	baseModel = name
	if originVariant, ok := model["variant"]; ok {
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
