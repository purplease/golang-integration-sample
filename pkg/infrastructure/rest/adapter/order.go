// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	orderService "github.com/purplease/golang-integration-sample/pkg/app/usecases/order"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/logger"
	orderRepository "github.com/purplease/golang-integration-sample/pkg/infrastructure/repository/order"
	orderController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/order"
	sdksql "github.com/purplease/golang-integration-sample/pkg/infrastructure/sql"
)

// OrderAdapter is a function that returns a order controller
func OrderAdapter(db *sdksql.DB, logger *logger.Logger) *orderController.Controller {
	mRepository := orderRepository.Repository{Store: db, Logger: logger}
	service := orderService.Service{OrderRepository: &mRepository}
	return &orderController.Controller{OrderService: service}
}
