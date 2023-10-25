package v1

import (
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"ig_server/api/v1/inference_tasks"
	"ig_server/api/v1/models"
	"ig_server/api/v1/response"
)

func InitRoutes(r *fizz.Fizz) {

	v1g := r.Group("v1", "ApiV1", "API version 1")

	tasksGroup := v1g.Group("inference_tasks", "Inference tasks", "Inference tasks related APIs")

	tasksGroup.POST("", []fizz.OperationOption{
		fizz.Summary("Create an inference task"),
		fizz.Response("400", "validation errors", response.ValidationErrorResponse{}, nil, nil),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(inference_tasks.CreateTask, 200))

	tasksGroup.GET("/:client_id/:task_id", []fizz.OperationOption{
		fizz.Summary("Get task details by task id"),
		fizz.Response("400", "validation errors", response.ValidationErrorResponse{}, nil, nil),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(inference_tasks.GetTaskById, 200))

	tasksGroup.GET("/:client_id/:task_id/images/:image_num", []fizz.OperationOption{
		fizz.Summary("Get task details by task id"),
		fizz.Response("400", "validation errors", response.ValidationErrorResponse{}, nil, nil),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(inference_tasks.GetTaskImage, 200))

	modelsGroup := v1g.Group("models", "Models", "Models related APIs")

	modelsGroup.GET("base", []fizz.OperationOption{
		fizz.Summary("Get the list of the base models"),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(models.GetBaseModels, 200))

	modelsGroup.GET("lora", []fizz.OperationOption{
		fizz.Summary("Get the list of the lora models"),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(models.GetLoraModels, 200))
}
