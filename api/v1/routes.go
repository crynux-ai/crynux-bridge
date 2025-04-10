package v1

import (
	"crynux_bridge/api/v1/application"
	"crynux_bridge/api/v1/count"
	"crynux_bridge/api/v1/inference_tasks"
	"crynux_bridge/api/v1/models"
	"crynux_bridge/api/v1/network"
	"crynux_bridge/api/v1/openrouter"
	"crynux_bridge/api/v1/response"

	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
)

func InitRoutes(r *fizz.Fizz) {

	v1g := r.Group("v1", "ApiV1", "API version 1")

	tasksGroup := v1g.Group("inference_tasks", "Inference tasks", "Inference tasks related APIs")

	tasksGroup.POST("", []fizz.OperationOption{
		fizz.Summary("Create an inference task"),
		fizz.Response("400", "validation errors", response.ValidationErrorResponse{}, nil, nil),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(inference_tasks.CreateTask, 200))

	tasksGroup.GET("/:client_id/:client_task_id", []fizz.OperationOption{
		fizz.Summary("Get task details by task id"),
		fizz.Response("400", "validation errors", response.ValidationErrorResponse{}, nil, nil),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(inference_tasks.GetTaskById, 200))

	tasksGroup.GET("/:client_id/:client_task_id/images/:index", []fizz.OperationOption{
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

	networkGroup := v1g.Group("network", "Network", "Network status")
	networkGroup.GET("nodes", []fizz.OperationOption{}, tonic.Handler(network.GetNodeStats, 200))

	applicationGroup := v1g.Group("application", "Application", "Application related APIs")
	applicationGroup.GET("/wallet/balance", []fizz.OperationOption{
		fizz.Summary("Get the balance of the application wallet"),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(application.GetWalletBalance, 200))

	countGroup := v1g.Group("count", "Count", "Task count related APIs")
	countGroup.GET("/task", []fizz.OperationOption{
		fizz.Summary("Count the task in the recent period"),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(count.CountTask, 200))

	// for openrouter, api: /completions and /chat/completions
	openrouterGroup := v1g.Group("openrouter", "OpenRouter", "OpenRouter related APIs")

	openrouterGroup.POST("/completions", []fizz.OperationOption{
		fizz.Summary("Api for openrouter, /completions"),
		fizz.Response("400", "validation errors", response.ValidationErrorResponse{}, nil, nil),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(openrouter.Completions, 200))

	openrouterGroup.POST("/chat/completions", []fizz.OperationOption{
		fizz.Summary("Api for openrouter, /chat/completions"),
		fizz.Response("400", "validation errors", response.ValidationErrorResponse{}, nil, nil),
		fizz.Response("500", "exception", response.ExceptionResponse{}, nil, nil),
	}, tonic.Handler(openrouter.ChatCompletions, 200))
}
