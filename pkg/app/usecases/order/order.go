// Package order provides the use case for order
package order

import (
	"context"

	orderDomain "github.com/purplease/golang-integration-sample/pkg/domain/order"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/repository"
)

// Service is a struct that contains the repository implementation for order use case
type Service struct {
	OrderRepository repository.Orders
}

// GetByID is a function that returns a order by diner ID
func (s *Service) GetByID(ctx context.Context, dinerID int64) ([]orderDomain.Response, error) {
	return s.OrderRepository.GetByID(ctx, dinerID)
}

// Create is a function that creates a order
func (s *Service) Create(ctx context.Context, order *NewOrder) (*orderDomain.Request, error) {
	orderModel := order.toDomainMapper()
	return s.OrderRepository.Create(ctx, orderModel)
}

// Delete is a function that deletes a order by id
func (s *Service) Delete(ctx context.Context, id int) error {
	return s.OrderRepository.Delete(ctx, id)
}
