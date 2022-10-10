package customer

import (
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/middleware"
	"kang-sayur-backend/model/domain/customer"
	"kang-sayur-backend/model/domain/verification"
)

type CustomerMiddleware struct {
	JWT                    jsonwebtoken.JWTContract
	VerificationRepository verification.VerificationRepository
	Customer               customer.CustomerRepository
}

// Process implements middleware.Contract
func (cm *CustomerMiddleware) Process(token string) *middleware.MiddlewareResponse {
	// check token kosong
	if token == "" {
		return &middleware.MiddlewareResponse{
			Message: "silahkan login terlebih dahulu",
			Status:  401,
			Is_pass: false,
			Reason:  "token kosong",
		}
	}

	// verify token
	payload, err := cm.JWT.ParseToken(token)
	if err != nil {
		return &middleware.MiddlewareResponse{
			Message: "kesalahan saat parsing token",
			Status:  500,
			Is_pass: false,
			Reason:  "error parsing token",
		}
	}

	email := payload["email"].(string)

	// check if customer is registered or not
	customer := cm.Customer.ReadByEmail(email)

	if customer.ID == "" || customer.Email == "" {
		return &middleware.MiddlewareResponse{
			Message: "pelanggan tidak terdaftar",
			Status:  500,
			Is_pass: false,
			Reason:  "data pelanggan tidak ada di database",
		}
	}

	if customer.VerificationStatus == "not verified" {
		return &middleware.MiddlewareResponse{
			Message: "pelanggan tidak terverifikasi",
			Status:  500,
			Is_pass: false,
			Reason:  "data pelanggan belum tidak terverifikasi di database",
		}
	}

	return &middleware.MiddlewareResponse{
		Message: "customer terauhtorisasi",
		Status:  200,
		Is_pass: true,
		Reason:  "terauthorisasi dengan sukses",
		Claim: struct {
			Id    string "json:\"id,omitempty\""
			Email string "json:\"email,omitempty\""
		}{
			Id:    customer.ID,
			Email: customer.Email,
		},
	}
}

func Middleware(verification verification.VerificationRepository, customer customer.CustomerRepository, jwt jsonwebtoken.JWTContract) middleware.Contract {
	return &CustomerMiddleware{
		JWT:                    jwt,
		VerificationRepository: verification,
		Customer:               customer,
	}
}
