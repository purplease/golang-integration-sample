// Package message provides the use case for message
package message

import (
	"time"

	domainMessage "github.com/purplease/golang-integration-sample/pkg/domain/message"
	"github.com/purplease/golang-integration-sample/pkg/integration/twilio/model"
)

func domaintoModelMapper(req *domainMessage.SMSRequest) *model.MessageRequest {
	return &model.MessageRequest{
		PhoneNumber: req.PhoneNumber,
		Body:        req.Body,
	}
}

func modelToDomainMapper(res *model.MessageResponse) *domainMessage.Response {
	return &domainMessage.Response{
		Body:      *res.Body,
		Status:    *res.Status,
		CreatedAt: stringToTimeMapper(res.SentAt),
	}
}

func stringToTimeMapper(req *string) time.Time {
	return time.Now() // TODO: convert the value to time
}
