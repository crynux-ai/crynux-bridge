package apikey

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ChangeUseLimitInput struct {
	ClientID string `path:"client_id" json:"client_id" description:"Client id" validate:"required"`
	UseLimit int64  `json:"use_limit" description:"Use limit" validate:"required"`
}

type ChangeUseLimitInputWithSignature struct {
	ChangeUseLimitInput
	Timestamp int64  `form:"timestamp" json:"timestamp" description:"Signature timestamp" validate:"required"`
	Signature string `form:"signature" json:"signature" description:"Signature" validate:"required"`
}

func ChangeUseLimit(c *gin.Context, in *ChangeUseLimitInputWithSignature) (*response.Response, error) {
	match, address, err := tools.ValidateSignature(in.ChangeUseLimitInput, in.Timestamp, in.Signature)

	if err != nil || !match {

		if err != nil {
			log.Debugln("error in sig validate: " + err.Error())
		}

		validationErr := response.NewValidationErrorResponse("signature", "Invalid signature")
		return nil, validationErr
	}
	appConfig := config.GetConfig()
	if address != appConfig.Blockchain.Account.Address {
		validationErr := response.NewValidationErrorResponse("client_id", "Invalid signer")
		return nil, validationErr
	}

	if err := tools.ChangeUseLimit(c.Request.Context(), config.GetDB(), in.ClientID, in.UseLimit); err != nil {
		log.Debugln("error in change use limit: " + err.Error())
		return nil, response.NewExceptionResponse(err)
	}

	return &response.Response{}, nil
}
