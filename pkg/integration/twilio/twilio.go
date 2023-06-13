package twilio

import (
	"context"

	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/model"
)

// Messages is a interface that contains the methods for the message service
type Messages interface {
	CreateSMS(context.Context, *model.MessageRequest) (*model.MessageResponse, error)
	CreateWhatsapp(context.Context, *model.MessageRequest) (*model.MessageResponse, error)
}

type Verifys interface {
	CreateOTP(context.Context, *model.OTPRequest) (*model.VerifyResponse, error)
	ValidateOTP(context.Context, *model.ValidateOTPRequest) (*model.VerifyResponse, error)
}
