// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	dinerController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/diner"
)

// DinerRoutes is a function that contains all diner routes
func DinerRoutes(router *gin.RouterGroup, controller *dinerController.Controller) {

	routerDiner := router.Group("/diners")
	{
		routerDiner.POST("/", controller.NewDiner)
		routerDiner.GET("/:id", controller.GetDinersByID)
		routerDiner.GET("/", controller.GetAllDiners)
		routerDiner.DELETE("/:id", controller.DeleteDiner)
	}

}
