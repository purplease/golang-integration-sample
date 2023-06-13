// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	messageUsecase "github.com/purplease/golang-integration-sample/pkg/app/usecases/message"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/logger"
	messageController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/message"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/config"
	messageIntegrationService "github.com/purplease/golang-integration-sample/pkg/integration/twilio/service"
)

// MessageAdapter is a function that returns a message controller
func MessageAdapter(cfg config.Config, logger *logger.Logger) *messageController.Controller {
	msgService := messageIntegrationService.NewIntegrationService(cfg)
	service := &messageUsecase.Service{MessageIntegration: msgService}
	return &messageController.Controller{MessageService: service}
}
