// Package verify provides the use case for verify
package verify

import (
	"context"
	"fmt"

	verifyDomain "github.com/purplease/golang-integration-sample/pkg/domain/verify"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio"
)

// Service is a struct that contains the repository implementation for verify use case
type Service struct {
	VerifyIntegration twilio.Verifys
}

// CreateOTP is a function that creates a verify OTP
func (s *Service) CreateOTP(ctx context.Context, verify *verifyDomain.CreateOTPRequest) (*verifyDomain.Response, error) {
	verifyModel := otpDomainToModelMapper(verify)
	respVerify, err := s.VerifyIntegration.CreateOTP(ctx, verifyModel)
	if err != nil {
		return nil, fmt.Errorf("error createOTP: %w", err)
	}
	return modelToDomainMapper(respVerify), nil
}

// ValidateOTP is a function that validates a SMS/whatsapp OTP
func (s *Service) ValidateOTP(ctx context.Context, verify *verifyDomain.ValidateOTPRequest) (*verifyDomain.Response, error) {
	verifyModel := ValidateDomaintoModelMapper(verify)
	respVerify, err := s.VerifyIntegration.ValidateOTP(ctx, verifyModel)
	if err != nil {
		return nil, fmt.Errorf("error createSMS: %w", err)
	}
	return modelToDomainMapper(respVerify), nil
}
