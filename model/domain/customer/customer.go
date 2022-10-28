package customer

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/customer"
	verification_request_body "kang-sayur-backend/model/web/request_body/verification"
)

type Customer struct {
	ID                 string `json:"id" bson:"_id"`
	Name               string `json:"name" bson:"name"`
	Phone              string `json:"phone" bson:"phone"`
	Email              string `json:"email" bson:"email"`
	Password           string `json:"password" bson:"password"`
	VerificationStatus string `json:"verification_status" bson:"verification_status"`
	DomicileID         string `json:"domicile_id" bson:"domicile_id"`
}

type CustomerService interface {
	Register(body *request_body.Register) *response.HTTPResponse

	Login(body *request_body.Login) *response.HTTPResponse

	ViewProfile(body *request_body.ViewProfile) *response.HTTPResponse

	UpdateProfile(body *request_body.UpdateProfile) *response.HTTPResponse

	ForgotPassword(body *request_body.ForgotPassword) *response.HTTPResponse

	ResetPassword(body *request_body.ResetPassword) *response.HTTPResponse

	ReadAll() *response.HTTPResponse

	RequestVerification(body *verification_request_body.RequestVerification) *response.HTTPResponse

	VerifyVerification(body *verification_request_body.Verify) *response.HTTPResponse
}

type CustomerRepository interface {
	Create(data *request_body.Register) (*Customer, error)

	ReadByEmail(email string) *Customer

	ReadByID(id string) *Customer

	UpdateProfile(data *request_body.UpdateProfile) error

	UpdatePassword(password string, customer_id string) error

	VerifyVerification(customer_id string) error

	Read() []Customer
}
