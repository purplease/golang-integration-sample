// Package verify provides the use case for verify
package verify

import (
	"time"

	domainVerify "github.com/purplease/golang-integration-sample/pkg/domain/verify"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/model"
)

func otpDomainToModelMapper(req *domainVerify.CreateOTPRequest) *model.OTPRequest {
	return &model.OTPRequest{
		To:      req.To,
		Channel: model.ToChannel(req.Channel),
	}
}

func ValidateDomaintoModelMapper(req *domainVerify.ValidateOTPRequest) *model.ValidateOTPRequest {
	return &model.ValidateOTPRequest{
		To:   req.To,
		Code: req.Code,
	}
}

func modelToDomainMapper(res *model.VerifyResponse) *domainVerify.Response {
	return &domainVerify.Response{
		Status:    *res.Status,
		CreatedAt: *res.CreatedAt,
	}
}

func stringToTimeMapper(req *string) time.Time {
	return time.Now() // TODO: convert the value to time
}
