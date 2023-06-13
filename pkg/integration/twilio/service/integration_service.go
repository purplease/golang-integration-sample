package service

import (
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/config"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/internal"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/outreach"
)

// IntegrationService handles the integration between Twilio and Outreach
type IntegrationService struct {
	TwilioClient   *internal.Client
	OutreachClient *outreach.Client
	Config         config.Config
}

// NewIntegrationService creates a new instance of the IntegrationService
func NewIntegrationService(cfg config.Config) *IntegrationService {
	s := &IntegrationService{Config: cfg}
	// Create Twilio client
	s.TwilioClient = internal.NewTwilioClient(s.Config)

	// Create Outreach client
	s.OutreachClient = outreach.NewOutreachClient(s.Config.OutreachAPIKey, s.Config.OutreachAPIBaseURL)

	return s
}

// Run starts the integration process
func (s *IntegrationService) Run() error {
	// Set up IVR webhook handler
	// http.HandleFunc("/ivr-webhook", s.IVRWebhookHandler)

	// Configure Twilio Voice URL for the phone number
	// twilioPhoneNumberSID := "PNbe238eecbc69d1a03459af9b7911826f" // "+14302456361" "+14155238886" // Replace with the SID of your Twilio phone number
	// voiceURL := "https://your-domain.com/ivr-webhook"            // Replace with the publicly accessible URL of your /ivr-webhook endpoint

	// err := s.TwilioClient.ConfigurePhoneNumberVoiceURL(twilioPhoneNumberSID, voiceURL)
	// if err != nil {
	// 	return err
	// }

	return nil
}
