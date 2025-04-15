package apikey

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type DeleteAPIKeyInput struct {
	ClientID string `path:"client_id" json:"client_id" description:"Client id" validate:"required"`
}

type DeleteAPIKeyInputWithSignature struct {
	DeleteAPIKeyInput
	Timestamp int64  `form:"timestamp" json:"timestamp" description:"Signature timestamp" validate:"required"`
	Signature string `form:"signature" json:"signature" description:"Signature" validate:"required"`
}

func DeleteAPIKey(c *gin.Context, in *DeleteAPIKeyInputWithSignature) (*response.Response, error) {
	match, address, err := tools.ValidateSignature(in.DeleteAPIKeyInput, in.Timestamp, in.Signature)

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
	// delete the API key
	if err := tools.DeleteAPIKey(c.Request.Context(), config.GetDB(), in.ClientID); err != nil {
		return nil, response.NewExceptionResponse(err)
	}
	return &response.Response{}, nil
}
