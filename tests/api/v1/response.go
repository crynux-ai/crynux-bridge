package v1

import (
	"crynux_bridge/api/v1/inference_tasks"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/models"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertValidationErrorResponse(t *testing.T, r *httptest.ResponseRecorder, fieldName, fieldMessage string) {

	assert.Equal(t, 400, r.Code, "wrong http status code")

	validationResponse := &response.ValidationErrorResponse{}

	err := json.Unmarshal(r.Body.Bytes(), validationResponse)
	assert.Equal(t, nil, err, "json unmarshal error")

	assert.Equal(t, "validation_error", validationResponse.GetErrorType(), "wrong validation message type")
	assert.Equal(t, fieldName, validationResponse.GetFieldName(), "wrong field name")
	assert.Equal(t, fieldMessage, validationResponse.GetFieldMessage(), "wrong field message")
}

func AssertExceptionResponse(t *testing.T, r *httptest.ResponseRecorder, message string) {
	assert.Equal(t, 500, r.Code, "wrong http status code")

	exceptionResponse := &response.ExceptionResponse{}

	err := json.Unmarshal(r.Body.Bytes(), exceptionResponse)
	assert.Equal(t, nil, err, "json unmarshal error")

	assert.Equal(t, message, exceptionResponse.GetMessage(), "wrong response message")
}

func AssertTaskResponse(t *testing.T, r *httptest.ResponseRecorder, task *models.InferenceTask) {
	assert.Equal(t, r.Code, 200, "wrong http status code")

	taskResponse := &inference_tasks.TaskResponse{}

	responseBytes := r.Body.Bytes()

	err := json.Unmarshal(responseBytes, taskResponse)
	assert.Equal(t, nil, err, "json unmarshal error")

	assert.Equal(t, "success", taskResponse.GetMessage(), "wrong message: "+string(responseBytes))
	assert.Equal(t, task.TaskId, taskResponse.Data.TaskId, "wrong task id")
	assert.Equal(t, task.TaskArgs, taskResponse.Data.TaskArgs, "wrong task args")
	assert.Equal(t, false, taskResponse.Data.CreatedAt.IsZero(), "wrong task created at")
	assert.Equal(t, false, taskResponse.Data.UpdatedAt.IsZero(), "wrong task updated at")
}

func AssertEmptySuccessResponse(t *testing.T, r *httptest.ResponseRecorder) {
	assert.Equal(t, 200, r.Code, "wrong http status code")

	emptyResponse := &response.Response{}

	responseBytes := r.Body.Bytes()

	err := json.Unmarshal(responseBytes, emptyResponse)
	assert.Equal(t, err, nil, "json unmarshal error")

	assert.Equal(t, "success", emptyResponse.GetMessage(), "wrong message: "+string(responseBytes))
}
