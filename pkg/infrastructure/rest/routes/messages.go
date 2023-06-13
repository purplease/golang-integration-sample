// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	messageController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/message"
)

// MessageRoutes is a function that contains all Message routes
func MessageRoutes(router *gin.RouterGroup, controller *messageController.Controller) {

	routerMessage := router.Group("/messages")
	{
		routerMessage.POST("/sms", controller.NewSMSMessage)
		routerMessage.POST("/whatsapp", controller.NewWhatsappMessage)
	}

}
