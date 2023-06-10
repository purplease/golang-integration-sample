package outreach

import (
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/model"
)

// Client represents the Outreach API client
type Client struct {
	APIKey  string
	BaseURL string
}

// NewOutreachClient creates a new instance of the Outreach client
func NewOutreachClient(apiKey, baseURL string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

// CreateContact creates a contact using the Outreach API
func (c *Client) CreateContact(contact *model.Contact) (*model.Contact, error) {
	// Implement the logic to create a contact using the Outreach API
	return nil, nil
}
