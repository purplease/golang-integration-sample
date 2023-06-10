// Package menu provides the use case for menu
package menu

import (
	domainMenu "github.com/purplease/golang-integration-sample/pkg/domain/menu"
)

func (n *NewMenu) toDomainMapper() *domainMenu.Menu {
	return &domainMenu.Menu{
		Name:        n.Name,
		Description: n.Description,
		Price:       n.Price,
	}
}
