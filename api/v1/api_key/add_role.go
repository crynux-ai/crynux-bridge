package apikey

import (
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/config"
	"crynux_bridge/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AddRoleInput struct {
	ClientID string      `path:"client_id" json:"client_id" description:"Client id" validate:"required"`
	Role     models.Role `json:"role" enum:"admin,chat" description:"Role to add" validate:"required"`
}

type AddRoleInputWithSignature struct {
	AddRoleInput
	Timestamp int64  `form:"timestamp" json:"timestamp" description:"Signature timestamp" validate:"required"`
	Signature string `form:"signature" json:"signature" description:"Signature" validate:"required"`
}

func AddRole(c *gin.Context, in *AddRoleInputWithSignature) (*response.Response, error) {
	match, address, err := tools.ValidateSignature(in.AddRoleInput, in.Timestamp, in.Signature)

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

	if err := tools.AddAPIKeyRole(c.Request.Context(), config.GetDB(), in.ClientID, in.Role); err != nil {
		log.Debugln("error in add api key role: " + err.Error())
		return nil, response.NewExceptionResponse(err)
	}

	return &response.Response{}, nil
}
