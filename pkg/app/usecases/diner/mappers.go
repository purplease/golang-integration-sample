// Package diner provides the use case for diner
package diner

import (
	domainDiner "github.com/purplease/golang-integration-sample/pkg/domain/diner"
)

func (n *NewDiner) toDomainMapper() *domainDiner.Diner {
	return &domainDiner.Diner{
		Name:        n.Name,
		TableNumber: n.TableNumber,
	}
}
