package customer

import (
	mailer "kang-sayur-backend/infrastructure/SMTP"
	"kang-sayur-backend/infrastructure/encryption"
	"kang-sayur-backend/infrastructure/identifier"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/model/domain/customer"
	forgotpassword "kang-sayur-backend/model/domain/forgot_password"
	"kang-sayur-backend/model/domain/verification"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/customer"
	verification_request_body "kang-sayur-backend/model/web/request_body/verification"
)

type customerService struct {
	customerRepository       customer.CustomerRepository
	forgotPasswordRepository forgotpassword.ForgotPasswordRepository
	verificationRepository   verification.VerificationRepository
	jwt                      jsonwebtoken.JWTContract
	bcrypt                   encryption.EncryptionContract
	mailer                   mailer.Contract
	uuid                     identifier.IdentifierContract
}

// ForgotPassword implements customer.CustomerService
func (cs *customerService) ForgotPassword(body *request_body.ForgotPassword) *response.HTTPResponse {
	// retrieve customer by its email
	customer := cs.customerRepository.ReadByEmail(body.Email)

	// check if customer exists or not
	if customer.ID == "" || customer.Email == "" {
		return &response.HTTPResponse{
			Message: "pengguna tidak terdaftar",
			Status:  404,
		}
	}

	// create token
	return nil
}

// ReadAll implements customer.CustomerService
func (cs *customerService) ReadAll() *response.HTTPResponse {
	panic("unimplemented")
}

// Register implements customer.CustomerService
func (cs *customerService) Register(body *request_body.Register) *response.HTTPResponse {
	panic("unimplemented")
}

// RequestVerification implements customer.CustomerService
func (cs *customerService) RequestVerification(body *verification_request_body.RequestVerification) *response.HTTPResponse {
	panic("unimplemented")
}

// ResetPassword implements customer.CustomerService
func (cs *customerService) ResetPassword(body *request_body.ResetPassword) *response.HTTPResponse {
	panic("unimplemented")
}

// UpdateProfile implements customer.CustomerService
func (cs *customerService) UpdateProfile(body *request_body.UpdateProfile) *response.HTTPResponse {
	panic("unimplemented")
}

// ValidateAccount implements customer.CustomerService
func (cs *customerService) ValidateAccount(body *request_body.ValidateAccount) *response.HTTPResponse {
	panic("unimplemented")
}

// VerifyVerification implements customer.CustomerService
func (cs *customerService) VerifyVerification(body *verification_request_body.Verify) *response.HTTPResponse {
	panic("unimplemented")
}

// ViewProfile implements customer.CustomerService
func (cs *customerService) ViewProfile(body *request_body.ViewProfile) *response.HTTPResponse {
	panic("unimplemented")
}

// login implements customer.CustomerService
func (cs *customerService) Login(body *request_body.Login) *response.HTTPResponse {
	panic("unimplemented")
}

func CustomerService(cr *customer.CustomerRepository, fpr *forgotpassword.ForgotPasswordRepository, vr *verification.VerificationRepository, jwt *jsonwebtoken.JWTContract, bcrypt *encryption.EncryptionContract, smtp *mailer.Contract, uuid *identifier.IdentifierContract) customer.CustomerService {
	return &customerService{
		customerRepository:       *cr,
		forgotPasswordRepository: *fpr,
		verificationRepository:   *vr,
		jwt:                      *jwt,
		bcrypt:                   *bcrypt,
		mailer:                   *smtp,
		uuid:                     *uuid,
	}
}
