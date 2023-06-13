// Package message contains the message controller
package message

// NewMessageRequest is a struct that contains the new message request information
type NewMessageRequest struct {
	PhoneNumber string `json:"phone_number" example:"+919901170563" binding:"required"`
	Body        string `json:"body" example:"Something" binding:"required"`
}
