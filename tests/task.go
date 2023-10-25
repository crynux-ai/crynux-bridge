package tests

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"ig_server/api/v1/response"
	"ig_server/config"
	"ig_server/models"
	"testing"
)

const FullTaskArgsJson string = `{
	"base_model": "runwayml/stable-diffusion-v1-5",
	"prompt": "best quality, ultra high res, photorealistic++++, 1girl, off-shoulder sweater, smiling, faded ash gray messy bun hair+, border light, depth of field, looking at viewer, closeup",
	"negative_prompt": "paintings, sketches, worst quality+++++, low quality+++++, normal quality+++++, lowres, normal quality, monochrome++, grayscale++, skin spots, acnes, skin blemishes, age spot, glans",
	"controlnet": {
		"preprocess": {
			"method": "canny",
			"args": {
				"low_threshold": 100,
				"high_threshold": 200
			}
		},
		"model": "lllyasviel/sd-controlnet-canny",
		"weight": 80,
		"image_dataurl": "image/png,base64:FFFFFF"
	},
	"refiner": {
		"model": "stabilityai/stable-diffusion-xl-refiner-1.0",
		"denoising_cutoff": 80,
		"steps": 25
	},
	"lora": {
		"model": "https://civitai.com/api/download/models/34562",
		"weight": 80
	},
	"vae": "stabilityai/sd-vae-ft-mse",
	"textual_inversion": "sd-concepts-library/cat-toy",
	"task_config": {
		"image_width": 512,
		"image_height": 512,
		"num_images": 9,
		"seed": 5123333,
		"steps": 30,
		"safety_checker": false,
		"cfg": 7
	}
}`

func NewClient() (*models.Client, error) {

	uuidV4, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	return &models.Client{ClientId: uuidV4.String()}, nil
}

func NewTask() (*models.InferenceTask, error) {

	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	task := &models.InferenceTask{
		Client:   *client,
		TaskArgs: FullTaskArgsJson,
	}

	if err := config.GetDB().Create(task).Error; err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return task, nil
}

func AssertTaskStatus(t *testing.T, taskID uint, status models.TaskStatus) *models.InferenceTask {

	task := &models.InferenceTask{}

	err := config.GetDB().Omit("pose").First(task, taskID).Error
	assert.Equal(t, nil, err, "error find task in db")

	assert.Equal(t, status, task.Status, "task status mismatch")

	return task
}
