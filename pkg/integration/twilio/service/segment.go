package service

import (
	"context"

	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/model"
)

// CreateSMS exposes twilio service to send SMS
func (s *IntegrationService) CreateSMS(ctx context.Context, req *model.MessageRequest) (*model.MessageResponse, error) {
	return s.TwilioClient.SendSMS(req)
}

// CreateWhatsapp exposes twilio service to send Whatsapp message
func (s *IntegrationService) CreateWhatsapp(ctx context.Context, req *model.MessageRequest) (*model.MessageResponse, error) {
	return s.TwilioClient.SendWhatsapp(req)
}

// CreateOTP exposes twilio service to send OTP message to Whatsapp or SMS
func (s *IntegrationService) CreateOTP(ctx context.Context, req *model.OTPRequest) (*model.VerifyResponse, error) {
	return s.TwilioClient.SendOTPVerification(req)
}

// ValidateOTP exposes twilio service to validate a OTP message
func (s *IntegrationService) ValidateOTP(ctx context.Context, req *model.ValidateOTPRequest) (*model.VerifyResponse, error) {
	return s.TwilioClient.ValidateOTPVerification(req)
}
