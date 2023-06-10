// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	dinerService "github.com/purplease/golang-integration-sample/pkg/app/usecases/diner"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/logger"
	dinerRepository "github.com/purplease/golang-integration-sample/pkg/infrastructure/repository/diner"
	dinerController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/diner"
	sdksql "github.com/purplease/golang-integration-sample/pkg/infrastructure/sql"
)

// DinerAdapter is a function that returns a diner controller
func DinerAdapter(db *sdksql.DB, logger *logger.Logger) *dinerController.Controller {
	mRepository := dinerRepository.Repository{Store: db, Logger: logger}
	service := dinerService.Service{DinerRepository: &mRepository}
	return &dinerController.Controller{DinerService: service}
}
