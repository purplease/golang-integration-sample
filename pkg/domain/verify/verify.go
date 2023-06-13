// Package verify contains the business logic for the verify entity
package verify

import (
	"context"
	"time"
)

// Response is a struct that contains the verify information
type Response struct {
	Body      string    `json:"body" example:"Sent Successfully"`
	Status    string    `json:"status" example:"ok"`
	CreatedAt time.Time `json:"created_at,omitempty" `
}

// CreateOTPRequest is a struct that contains the OTP information
type CreateOTPRequest struct {
	To      string `json:"to" example:"+919901170563"`
	Channel string `json:"channel" example:"whatsapp"`
}

// ValidateOTPRequest is a struct that contains the verification of OTP information
type ValidateOTPRequest struct {
	To   string `json:"to" example:"+919901170563"`
	Code string `json:"code" example:"123456"`
}

// Service is a interface that contains the methods for the verify service
type Service interface {
	CreateOTP(context.Context, *CreateOTPRequest) (*Response, error)
	ValidateOTP(context.Context, *ValidateOTPRequest) (*Response, error)
}
