package internal

import (
	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// Client represents the Twilio API client
type Client struct {
	AccountSID string
	AuthToken  string
}

// NewTwilioClient creates a new instance of the Twilio client
func NewTwilioClient(accountSID, authToken string) *Client {
	return &Client{
		AccountSID: accountSID,
		AuthToken:  authToken,
	}
}

// SendSMS sends an SMS message using the Twilio API
func (c *Client) SendSMS(to, from, body string) (*rest.Message, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.AuthToken,
		Username: c.AccountSID,
	})
	return client.Api.CreateMessage(&openapi.CreateMessageParams{To: &to, From: &from, Body: &body})
}

// MakeVoiceCallWithIVR makes a voice call with Interactive Voice Response using the Twilio API
func (c *Client) MakeVoiceCallWithIVR(to, from, url string) (*rest.Call, error) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.AuthToken,
		Username: c.AccountSID,
	})
	_ = client
	//return client.Calls.Create(to, from, url, nil)
	return nil, nil
}

// ConfigurePhoneNumberVoiceURL configures the Voice URL for a Twilio phone number
func (c *Client) ConfigurePhoneNumberVoiceURL(phoneNumberSID, voiceURL string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.AuthToken,
		Username: c.AccountSID,
	})
	_, err := client.Api.CreateIncomingPhoneNumber(
		&openapi.CreateIncomingPhoneNumberParams{
			VoiceUrl:   &voiceURL,
			AddressSid: &phoneNumberSID,
		})

	if err != nil {
		return err
	}

	return nil
}
