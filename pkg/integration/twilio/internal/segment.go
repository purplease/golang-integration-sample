package internal

import (
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/model"
	"github.com/twilio/twilio-go/rest"
)

// SegmentIntegration handles the integration with Twilio Segment
type SegmentIntegration struct {
	TwilioClient *Client
}

// NewSegmentIntegration creates a new instance of the Twilio Segment integration
func NewSegmentIntegration(twilioClient *Client) *SegmentIntegration {
	return &SegmentIntegration{
		TwilioClient: twilioClient,
	}
}

// SendData sends data to Twilio Segment
func (s *SegmentIntegration) SendData(data *model.SegmentData) error {
	// Implement the logic to send data to Twilio Segment
	return nil
}

// SendSMS sends an SMS message using Twilio
func (s *SegmentIntegration) SendSMS(to, from, body string) (*rest.Message, error) {
	return s.TwilioClient.SendSMS(to, from, body)
}

// MakeVoiceCallWithIVR makes a voice call with IVR using Twilio
func (s *SegmentIntegration) MakeVoiceCallWithIVR(to, from, url string) (*rest.Call, error) {
	return s.TwilioClient.MakeVoiceCallWithIVR(to, from, url)
}
