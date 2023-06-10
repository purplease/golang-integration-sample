// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	menuService "github.com/purplease/golang-integration-sample/pkg/app/usecases/menu"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/logger"
	menuRepository "github.com/purplease/golang-integration-sample/pkg/infrastructure/repository/menu"
	menuController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/menu"
	sdksql "github.com/purplease/golang-integration-sample/pkg/infrastructure/sql"
)

// MenuAdapter is a function that returns a menu controller
func MenuAdapter(db *sdksql.DB, logger *logger.Logger) *menuController.Controller {
	mRepository := menuRepository.Repository{Store: db, Logger: logger}
	service := menuService.Service{MenuRepository: &mRepository}
	return &menuController.Controller{MenuService: service}
}
