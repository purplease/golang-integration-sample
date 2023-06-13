// Package verify contains the verify controller
package verify

// NewCreateOTPRequest is a struct that contains the new verify OTP request information
type NewCreateOTPRequest struct {
	To      string `json:"to" example:"+919901170563, rajesh.biswas@purplease.com" binding:"required"`
	Channel string `json:"channel" example:"sms, whatsapp, call, email" binding:"required"`
}

// NewValidateOTPRequest is a struct that contains the new validate OTP request information
type NewValidateOTPRequest struct {
	To   string `json:"to,omitempty" example:"+919901170563, rajesh.biswas@purplease.com" validate:"required"`
	Code string `json:"code,omitempty" example:"123456" validate:"required"`
}
