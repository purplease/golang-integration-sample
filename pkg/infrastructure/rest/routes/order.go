// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	orderController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/order"
)

// OrderRoutes is a function that contains all order routes
func OrderRoutes(router *gin.RouterGroup, controller *orderController.Controller) {

	routerOrder := router.Group("/orders")
	{
		routerOrder.POST("/", controller.NewOrder)
		routerOrder.GET("/:id", controller.GetOrdersByDinerID)
		routerOrder.DELETE("/:id", controller.DeleteOrder)
	}

}
