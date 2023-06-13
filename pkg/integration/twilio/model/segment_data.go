package model

import "time"

// SegmentData represents the data to be exchanged between Twilio and Outreach
type SegmentData struct {
	// Define the structure of the data exchanged between Twilio and Outreach
}

type Contact struct {
}

// MessageResponse is a struct that contains the message information
type MessageResponse struct {
	Body   *string
	Status *string
	SentAt *string
}

// MessageRequest is a struct that contains the message information
type MessageRequest struct {
	PhoneNumber string
	Body        string
	From        string
}

// Channel can be `email`, `sms`, `whatsapp`, `call`, `sna` or `auto`.
type Channel string

var (
	EMAIL    Channel = "email"
	SMS      Channel = "sms"
	WHATSAPP Channel = "whatsapp"
	CALL     Channel = "call"
)

func ToChannel(channel string) Channel {
	switch channel {
	case "email":
		return EMAIL
	case "whatsapp":
		return WHATSAPP
	case "call":
		return CALL
	}
	return SMS
}

type OTPRequest struct {
	To      string
	Channel Channel
}

type ValidateOTPRequest struct {
	To   string
	Code string
}

type VerifyResponse struct {
	Status           *string
	SendCodeAttempts *[]interface{}
	CreatedAt        *time.Time
}
