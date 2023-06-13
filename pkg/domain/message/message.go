// Package message contains the business logic for the message entity
package message

import (
	"context"
	"time"
)

// Response is a struct that contains the message information
type Response struct {
	Body      string    `json:"body" example:"Sent Successfully"`
	Status    string    `json:"status" example:"ok"`
	CreatedAt time.Time `json:"created_at,omitempty" `
}

// SMSRequest is a struct that contains the message information
type SMSRequest struct {
	PhoneNumber string `json:"phone_number" example:"+919901170563"`
	Body        string `json:"body" example:"Sent Successfully"`
}

// Service is a interface that contains the methods for the message service
type Service interface {
	// Get(context.Context, string) (*Response, error)
	// GetAll(context.Context) ([]*Response, error)
	CreateSMS(context.Context, *SMSRequest) (*Response, error)
	CreateWhatsapp(context.Context, *SMSRequest) (*Response, error)
}
