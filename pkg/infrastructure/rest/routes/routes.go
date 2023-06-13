// Package routes contains all routes of the application
package routes

import (
	// swaggerFiles for documentation
	"github.com/gin-gonic/gin"
	_ "github.com/purplease/golang-integration-sample/docs"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/config"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/logger"
	"github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/adapter"
	errorsController "github.com/purplease/golang-integration-sample/pkg/infrastructure/rest/controllers/errors"
	sdksql "github.com/purplease/golang-integration-sample/pkg/infrastructure/sql"
	twilioConfig "github.com/purplease/golang-integration-sample/pkg/integration/twilio/config"
)

type DI struct {
	Router *gin.Engine
	DB     *sdksql.DB
	Logger *logger.Logger
	Config config.AppConfig
}

//	@title			Golang Integration Sample
//	@version		2.0
//	@description	Documentation's Golang Integration Sample
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Rajesh Kumar Biswas
//	@contact.url	https://github.com/purplease-rajesh
//	@contact.email	rajesh.biswas@purplease.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// ApplicationV1Router is a function that contains all routes of the application
//
//	@host		localhost:8080
//	@BasePath	/v1
func ApplicationV1Router(di DI) {
	// the application errors will be processed here before returning to the caller
	di.Router.Use(errorsController.Handler)

	routerV1 := di.Router.Group("/v1")
	{
		MenuRoutes(routerV1, adapter.MenuAdapter(di.DB, di.Logger))
		DinerRoutes(routerV1, adapter.DinerAdapter(di.DB, di.Logger))
		OrderRoutes(routerV1, adapter.OrderAdapter(di.DB, di.Logger))
		if di.Config.TwilioConfig.Enabled {
			twilioConfig := twilioConfig.Config{
				AccountSID:          di.Config.TwilioConfig.AccountSID,
				AuthToken:           di.Config.TwilioConfig.AuthToken,
				VerifyServiceSID:    di.Config.TwilioConfig.VerifyServiceSID,
				WhatsappPhoneNumber: di.Config.TwilioConfig.WhatsappPhoneNumber,
				PhoneNumberSID:      di.Config.TwilioConfig.PhoneNumberSID,
				SMSPhoneNumber:      di.Config.TwilioConfig.SMSPhoneNumber,
			}

			MessageRoutes(routerV1, adapter.MessageAdapter(twilioConfig, di.Logger))
			VerifyRoutes(routerV1, adapter.VerifyAdapter(twilioConfig, di.Logger))
		}

	}
}
