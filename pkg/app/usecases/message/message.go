// Package message provides the use case for message
package message

import (
	"context"
	"fmt"

	messageDomain "github.com/purplease/golang-integration-sample/pkg/domain/message"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio"
)

// Service is a struct that contains the repository implementation for message use case
type Service struct {
	MessageIntegration twilio.Messages
}

// CreateSMS is a function that creates a message
func (s *Service) CreateSMS(ctx context.Context, message *messageDomain.SMSRequest) (*messageDomain.Response, error) {
	messageModel := domaintoModelMapper(message)
	respMessage, err := s.MessageIntegration.CreateSMS(ctx, messageModel)
	if err != nil {
		return nil, fmt.Errorf("error createSMS message: %w", err)
	}
	return modelToDomainMapper(respMessage), nil
}

// CreateWhatsapp is a function that creates a whatsapp message
func (s *Service) CreateWhatsapp(ctx context.Context, message *messageDomain.SMSRequest) (*messageDomain.Response, error) {
	messageModel := domaintoModelMapper(message)
	respMessage, err := s.MessageIntegration.CreateWhatsapp(ctx, messageModel)
	if err != nil {
		return nil, fmt.Errorf("error createWhatsapp message: %w", err)
	}
	return modelToDomainMapper(respMessage), nil
}
