// Package diner provides the use case for diner
package diner

import (
	"context"

	dinerDomain "github.com/purplease/golang-integration-sample/pkg/domain/diner"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/repository"
)

// Service is a struct that contains the repository implementation for diner use case
type Service struct {
	DinerRepository repository.Diners
}

// GetAll is a function that returns all diners
func (s *Service) GetAll(ctx context.Context, page int64, limit int64) (*PaginationResultDiner, error) {
	all, err := s.DinerRepository.GetAll(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	return &PaginationResultDiner{
		Data:       all.Data,
		Total:      all.Total,
		Limit:      all.Limit,
		Current:    all.Current,
		NextCursor: all.NextCursor,
		PrevCursor: all.PrevCursor,
		NumPages:   all.NumPages,
	}, nil
}

// GetByID is a function that returns a diner by id
func (s *Service) GetByID(ctx context.Context, id int64) (*dinerDomain.Diner, error) {
	return s.DinerRepository.GetByID(ctx, id)
}

// Create is a function that creates a diner
func (s *Service) Create(ctx context.Context, diner *NewDiner) (*dinerDomain.Diner, error) {
	dinerModel := diner.toDomainMapper()
	return s.DinerRepository.Create(ctx, dinerModel)
}

// Delete is a function that deletes a diner by id
func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.DinerRepository.Delete(ctx, id)
}
