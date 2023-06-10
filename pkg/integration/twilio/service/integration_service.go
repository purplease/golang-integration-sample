package service

import (
	"fmt"
	"log"
	"net/http"

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
	return &IntegrationService{
		Config: cfg,
	}
}

// Run starts the integration process
func (s *IntegrationService) Run() error {

	// Create Twilio client
	s.TwilioClient = internal.NewTwilioClient(s.Config.TwilioAccountSID, s.Config.TwilioAuthToken)

	// Create Outreach client
	s.OutreachClient = outreach.NewOutreachClient(s.Config.OutreachAPIKey, s.Config.OutreachAPIBaseURL)

	// Set up IVR webhook handler
	http.HandleFunc("/ivr-webhook", s.IVRWebhookHandler)

	// Configure Twilio Voice URL for the phone number
	twilioPhoneNumberSID := "your-twilio-phone-number-sid" // Replace with the SID of your Twilio phone number
	voiceURL := "https://your-domain.com/ivr-webhook"      // Replace with the publicly accessible URL of your /ivr-webhook endpoint

	err := s.TwilioClient.ConfigurePhoneNumberVoiceURL(twilioPhoneNumberSID, voiceURL)
	if err != nil {
		return err
	}

	// Start the server
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Start integration process
	err = s.processIntegration()
	if err != nil {
		return err
	}

	return nil
}

// processIntegration handles the data flow between Twilio and Outreach
func (s *IntegrationService) processIntegration() error {
	// Implement the logic to handle the data flow between Twilio and Outreach
	return nil
}

// IVRWebhookHandler handles the webhook request from Twilio during the IVR call
func (s *IntegrationService) IVRWebhookHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the IVR response from the request
	ivrResponse := r.FormValue("Digits") // Assuming the response is captured as "Digits"

	fmt.Printf("ivr response %s", ivrResponse)
	// Save the IVR response against the person
	// Implement your logic here to save the response, e.g., update a person's record in a database

	// Optionally, you can return a response to Twilio to control the call flow
	// For example, you can use TwiML (Twilio Markup Language) to instruct Twilio on how to proceed with the call
	response := []byte(`
		<Response>
			<Hangup/>
		</Response>
	`)
	w.Header().Set("Content-Type", "application/xml")
	w.Write(response)
}
