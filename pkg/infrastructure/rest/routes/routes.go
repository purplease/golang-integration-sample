// Package routes contains all routes of the application
package routes

import (
	// swaggerFiles for documentation
	"github.com/gin-gonic/gin"
	_ "github.com/purplease/golang-integration-sample/docs"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/logger"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/adapter"
	errorsController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/errors"
	sdksql "github.com/purplease/golang-integration-sample/pkg/infrastructure/sql"
)

//	@title			Golang REST APIs
//	@version		2.0
//	@description	Documentation's Golang REST APIs
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Rajesh Kumar Biswas
//	@contact.url	http://github.com/Raj63
//	@contact.email	biswas.rajesh63@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// ApplicationV1Router is a function that contains all routes of the application
//
//	@host		localhost:8080
//	@BasePath	/v1
func ApplicationV1Router(router *gin.Engine, db *sdksql.DB, logger *logger.Logger) {
	// the application errors will be processed here before returning to the caller
	router.Use(errorsController.Handler)

	routerV1 := router.Group("/v1")
	{
		MenuRoutes(routerV1, adapter.MenuAdapter(db, logger))
		DinerRoutes(routerV1, adapter.DinerAdapter(db, logger))
		OrderRoutes(routerV1, adapter.OrderAdapter(db, logger))
	}
}
