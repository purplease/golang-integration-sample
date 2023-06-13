// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	verifyController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/verify"
)

// VerifyRoutes is a function that contains all Verify routes
func VerifyRoutes(router *gin.RouterGroup, controller *verifyController.Controller) {

	routerVerify := router.Group("/otp")
	{
		routerVerify.POST("/create", controller.NewCreateOTP)
		routerVerify.POST("/validate", controller.NewValidateOTP)
	}

}
