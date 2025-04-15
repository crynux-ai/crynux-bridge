package apikey

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CreateAPIKeyInput struct {
	ClientID string `path:"client_id" json:"client_id" description:"Client id" validate:"required"`
}

type CreateAPIKeyInputWithSignature struct {
	CreateAPIKeyInput
	Timestamp int64  `form:"timestamp" json:"timestamp" description:"Signature timestamp" validate:"required"`
	Signature string `form:"signature" json:"signature" description:"Signature" validate:"required"`
}

type CreateAPIKeyOutput struct {
	APIKey    string `json:"api_key" description:"Generated API key"`
	ExpiresAt int64  `json:"expires_at" description:"API key expiration time"`
}

type CreateAPIKeyResponse struct {
	response.Response
	Data *CreateAPIKeyOutput `json:"data"`
}

func CreateAPIKey(c *gin.Context, in *CreateAPIKeyInputWithSignature) (*CreateAPIKeyResponse, error) {
	match, address, err := tools.ValidateSignature(in.CreateAPIKeyInput, in.Timestamp, in.Signature)

	if err != nil || !match {

		if err != nil {
			log.Debugln("error in sig validate: " + err.Error())
		}

		validationErr := response.NewValidationErrorResponse("signature", "Invalid signature")
		return nil, validationErr
	}

	if in.ClientID != address {
		validationErr := response.NewValidationErrorResponse("client_id", "Invalid client id")
		return nil, validationErr
	}

	_, err = tools.CreateClientIfNotExist(c, config.GetDB(), in.ClientID)

	if err != nil {
		log.Debugln("error in create client: " + err.Error())
		return nil, response.NewExceptionResponse(err)
	}

	// generate a new API key
	apikey, expiresAt, err := tools.GenerateAPIKey(c.Request.Context(), config.GetDB(), in.ClientID)
	if err != nil {
		log.Debugln("error in generate apikey: " + err.Error())
		return nil, response.NewExceptionResponse(err)
	}
	return &CreateAPIKeyResponse{
		Data: &CreateAPIKeyOutput{
			APIKey:    apikey,
			ExpiresAt: expiresAt,
		},
	}, nil
}
