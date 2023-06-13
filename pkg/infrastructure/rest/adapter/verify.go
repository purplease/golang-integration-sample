// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	verifyUsecase "github.com/purplease/golang-integration-sample/pkg/app/usecases/verify"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/logger"
	verifyController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/verify"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/config"
	verifyIntegrationService "github.com/purplease/golang-integration-sample/pkg/integration/twilio/service"
)

// VerifyAdapter is a function that returns a verify controller
func VerifyAdapter(cfg config.Config, logger *logger.Logger) *verifyController.Controller {
	msgService := verifyIntegrationService.NewIntegrationService(cfg)
	service := &verifyUsecase.Service{VerifyIntegration: msgService}
	return &verifyController.Controller{VerifyService: service}
}
