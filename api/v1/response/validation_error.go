package response

type ValidationErrorResponse struct {
	ErrorResponse

	Data struct {
		FieldName string `json:"field_name" description:"The name of the field that fails the validation"`
		Message   string `json:"message" description:"The error type of the validation that fails"`
	} `json:"data" description:"The validation error detail"`
}

func (ver *ValidationErrorResponse) SetFieldName(fieldName string) {
	ver.Data.FieldName = fieldName
}

func (ver *ValidationErrorResponse) SetFieldMessage(message string) {
	ver.Data.Message = message
}

func (ver *ValidationErrorResponse) GetFieldName() string {
	return ver.Data.FieldName
}

func (ver *ValidationErrorResponse) GetFieldMessage() string {
	return ver.Data.Message
}

func NewValidationErrorResponse(fieldName, fieldMessage string) *ValidationErrorResponse {
	r := &ValidationErrorResponse{}
	r.SetErrorType("validation_error")
	r.SetFieldName(fieldName)
	r.SetFieldMessage(fieldMessage)
	return r
}
