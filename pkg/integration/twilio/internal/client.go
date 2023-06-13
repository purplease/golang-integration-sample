package internal

import (
	"fmt"

	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/config"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/model"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	verifyApi "github.com/twilio/twilio-go/rest/verify/v2"
)

// Client represents the Twilio API client
type Client struct {
	AccountSID          string
	AuthToken           string
	VerifyServiceSID    string
	WhatsappPhoneNumber string
	SMSPhoneNumber      string
}

// NewTwilioClient creates a new instance of the Twilio client
func NewTwilioClient(cfg config.Config) *Client {
	return &Client{
		AccountSID:          cfg.AccountSID,
		AuthToken:           cfg.AuthToken,
		VerifyServiceSID:    cfg.VerifyServiceSID,
		WhatsappPhoneNumber: cfg.WhatsappPhoneNumber,
		SMSPhoneNumber:      cfg.SMSPhoneNumber,
	}
}

// SendSMS sends an SMS message using the Twilio API
func (c *Client) SendSMS(req *model.MessageRequest) (*model.MessageResponse, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.AuthToken,
		Username: c.AccountSID,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(req.PhoneNumber)
	params.SetFrom(c.SMSPhoneNumber)
	params.SetBody(req.Body)

	msg, err := client.Api.CreateMessage(params)
	if err != nil {
		return nil, fmt.Errorf("error sending SMS message: %v", err)
	}
	return &model.MessageResponse{
		Body:   msg.Body,
		Status: msg.Status,
		SentAt: msg.DateSent,
	}, nil
}

// SendWhatsapp sends an whatsapp message using the Twilio API
func (c *Client) SendWhatsapp(req *model.MessageRequest) (*model.MessageResponse, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.AuthToken,
		Username: c.AccountSID,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("whatsapp:" + req.PhoneNumber)
	params.SetFrom("whatsapp:" + c.WhatsappPhoneNumber)
	params.SetBody(req.Body)

	msg, err := client.Api.CreateMessage(params) // TODO: handle whatsapp flow
	if err != nil {
		return nil, fmt.Errorf("error sending Whatsapp message: %v", err)
	}
	return &model.MessageResponse{
		Body:   msg.Body,
		Status: msg.Status,
		SentAt: msg.DateSent,
	}, nil
}

// SendOTPVerification sends an whatsapp/SMS OTP message using the Twilio API
func (c *Client) SendOTPVerification(req *model.OTPRequest) (*model.VerifyResponse, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.AuthToken,
		Username: c.AccountSID,
	})
	params := &verifyApi.CreateVerificationParams{}
	params.SetTo(req.To)
	params.SetChannel(string(req.Channel))

	resp, err := client.VerifyV2.CreateVerification(c.VerifyServiceSID, params)
	if err != nil {
		return nil, err
	}

	return &model.VerifyResponse{
		Status:           resp.Status,
		SendCodeAttempts: resp.SendCodeAttempts,
		CreatedAt:        resp.DateCreated,
	}, nil
}

// ValidateOTPVerification validates an whatsapp/SMS OTP message using the Twilio API
func (c *Client) ValidateOTPVerification(req *model.ValidateOTPRequest) (*model.VerifyResponse, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.AuthToken,
		Username: c.AccountSID,
	})
	params := &verifyApi.CreateVerificationCheckParams{}
	params.SetTo(req.To)
	params.SetCode(req.Code)

	resp, err := client.VerifyV2.CreateVerificationCheck(c.VerifyServiceSID, params)
	if err != nil {
		return nil, fmt.Errorf("error creating verifier: %v", err)
	}
	return &model.VerifyResponse{
		Status:    resp.Status,
		CreatedAt: resp.DateCreated,
	}, nil
}

// // MakeVoiceCallWithIVR makes a voice call with Interactive Voice Response using the Twilio API
// func (c *Client) MakeVoiceCallWithIVR(to, from, url string) (*rest.Call, error) {
// 	client := twilio.NewRestClientWithParams(twilio.ClientParams{
// 		Password: c.AuthToken,
// 		Username: c.AccountSID,
// 	})
// 	_ = client
// 	//return client.Calls.Create(to, from, url, nil)
// 	return nil, nil
// }

// // ConfigurePhoneNumberVoiceURL configures the Voice URL for a Twilio phone number
// func (c *Client) ConfigurePhoneNumberVoiceURL(phoneNumberSID, voiceURL string) error {
// 	client := twilio.NewRestClientWithParams(twilio.ClientParams{
// 		Password: c.AuthToken,
// 		Username: c.AccountSID,
// 	})
// 	_, err := client.Api.CreateIncomingPhoneNumber(
// 		&openapi.CreateIncomingPhoneNumberParams{
// 			VoiceUrl:   &voiceURL,
// 			AddressSid: &phoneNumberSID,
// 		})

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
