package verification

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/verification"
)

type Verification struct {
	ID               string `json:"id" bson:"_id"`
	Token            string `json:"token" bson:"token"`
	CustomerID       string `json:"customer_id" bson:"customer_id"`
	VerificationCode string `json:"Verification_code" bson:"Verification_code"`
}

type VerificationService interface {
	Request(body *request_body.RequestVerification) *response.HTTPResponse

	Verify(body *request_body.Verify) *response.HTTPResponse
}

type VerificationRepository interface {
	Upsert(token string, customer_id string, verification_code string) error

	ReadOne(token string, customer_id string) (*Verification, error)

	Delete(customer_id string) error
}
