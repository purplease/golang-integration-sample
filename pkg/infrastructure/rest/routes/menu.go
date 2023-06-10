// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	menuController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/menu"
)

// MenuRoutes is a function that contains all menu routes
func MenuRoutes(router *gin.RouterGroup, controller *menuController.Controller) {

	routerMenu := router.Group("/menus")
	{
		routerMenu.POST("/", controller.NewMenu)
		routerMenu.GET("/:id", controller.GetMenusByID)
		routerMenu.GET("/top", controller.GetTopMenus)
		routerMenu.GET("/", controller.GetAllMenus)
		routerMenu.DELETE("/:id", controller.DeleteMenu)
	}

}
