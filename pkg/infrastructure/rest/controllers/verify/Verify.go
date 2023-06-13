// Package verify contains the verify controller
package verify

import (
	domainErrors "github.com/purplease/golang-integration-sample/pkg/domain/errors"
	domainVerify "github.com/purplease/golang-integration-sample/pkg/domain/verify"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the verify service
type Controller struct {
	VerifyService domainVerify.Service
}

// NewCreateOTP godoc
//
//	@Tags			verifys
//	@Summary		Create New SMS/whatsapp Verify OTP
//	@Description	Create new SMS/whatsapp verify OTP from external system like Twilio
//	@Accept			json
//	@Produce		json
//	@Param			data	body		NewCreateOTPRequest	true	"body data"
//	@Success		201		{object}	domainVerify.Response
//	@Failure		400		{object}	VerifyResponse
//	@Failure		500		{object}	VerifyResponse
//	@Router			/otp/create [post]
func (c *Controller) NewCreateOTP(ctx *gin.Context) {
	var request NewCreateOTPRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newVerify := domainVerify.CreateOTPRequest{
		To:      request.To,
		Channel: request.Channel,
	}
	var result *domainVerify.Response
	var err error

	result, err = c.VerifyService.CreateOTP(ctx.Request.Context(), &newVerify)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

// NewValidateOTP godoc
//
//	@Tags			verifys
//	@Summary		Validates a SMS/whatsapp OTP
//	@Description	Validate a SMS/whatsapp OTP from external system like Twilio
//	@Accept			json
//	@Produce		json
//	@Param			data	body		NewValidateOTPRequest	true	"body data"
//	@Success		201		{object}	domainVerify.Response
//	@Failure		400		{object}	VerifyResponse
//	@Failure		500		{object}	VerifyResponse
//	@Router			/otp/validate [post]
func (c *Controller) NewValidateOTP(ctx *gin.Context) {
	var request NewValidateOTPRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newVerify := domainVerify.ValidateOTPRequest{
		To:   request.To,
		Code: request.Code,
	}
	var result *domainVerify.Response
	var err error

	result, err = c.VerifyService.ValidateOTP(ctx.Request.Context(), &newVerify)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}
