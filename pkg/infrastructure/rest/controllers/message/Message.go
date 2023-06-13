// Package message contains the message controller
package message

import (
	domainErrors "github.com/purplease/golang-integration-sample/pkg/domain/errors"
	domainMessage "github.com/purplease/golang-integration-sample/pkg/domain/message"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the message service
type Controller struct {
	MessageService domainMessage.Service
}

// NewSMSMessage godoc
//
//	@Tags			messages
//	@Summary		Create New SMS Message
//	@Description	Create new SMS message on the external system like TWILIO
//	@Accept			json
//	@Produce		json
//	@Param			data	body		NewMessageRequest	true	"body data"
//	@Success		201		{object}	domainMessage.Response
//	@Failure		400		{object}	MessageResponse
//	@Failure		500		{object}	MessageResponse
//	@Router			/messages/sms [post]
func (c *Controller) NewSMSMessage(ctx *gin.Context) {
	var request NewMessageRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newMessage := domainMessage.SMSRequest{
		PhoneNumber: request.PhoneNumber,
		Body:        request.Body,
	}
	var result *domainMessage.Response
	var err error

	result, err = c.MessageService.CreateSMS(ctx.Request.Context(), &newMessage)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// NewWhatsappMessage godoc
//
//	@Tags			messages
//	@Summary		Create New Whatsapp Message
//	@Description	Create new Whatsapp message on the external system like TWILIO
//	@Accept			json
//	@Produce		json
//	@Param			data	body		NewMessageRequest	true	"body data"
//	@Success		201		{object}	domainMessage.Response
//	@Failure		400		{object}	MessageResponse
//	@Failure		500		{object}	MessageResponse
//	@Router			/messages/whatsapp [post]
func (c *Controller) NewWhatsappMessage(ctx *gin.Context) {
	var request NewMessageRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	newMessage := domainMessage.SMSRequest{
		PhoneNumber: request.PhoneNumber,
		Body:        request.Body,
	}
	var result *domainMessage.Response
	var err error

	result, err = c.MessageService.CreateWhatsapp(ctx.Request.Context(), &newMessage)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
