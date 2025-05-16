package apikey

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type CreateAPIKeyInput struct {
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

func CreateAPIKey(c *gin.Context, in *CreateAPIKeyInput) (*CreateAPIKeyResponse, error) {
	match, address, err := tools.ValidateSignature(map[string]interface{}{}, in.Timestamp, in.Signature)

	if err != nil || !match {

		if err != nil {
			log.Debugln("error in sig validate: " + err.Error())
		}

		validationErr := response.NewValidationErrorResponse("signature", "Invalid signature")
		return nil, validationErr
	}

	appConfig := config.GetConfig()
	if address != appConfig.Blockchain.Account.Address {
		validationErr := response.NewValidationErrorResponse("signature", "Invalid signer")
		return nil, validationErr
	}

	clientID := uuid.New().String()

	// generate a new API key
	apikey, expiresAt, err := tools.GenerateAPIKey(c.Request.Context(), config.GetDB(), clientID)
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
