package admin

import (
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/middleware"
	"kang-sayur-backend/model/domain/admin"
)

type AdminMiddleware struct {
	AdminRepository admin.AdminRepository
	JWT             jsonwebtoken.JWTContract
}

// Process implements middleware.Contract
func (am *AdminMiddleware) Process(token string) *middleware.MiddlewareResponse {
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
	payload, err := am.JWT.ParseToken(token)
	if err != nil {
		return &middleware.MiddlewareResponse{
			Message: "kesalahan saat parsing token",
			Status:  500,
			Is_pass: false,
			Reason:  "error parsing token",
		}
	}

	email := payload["email"].(string)

	// check admin by email
	admin := am.AdminRepository.CheckEmail(email)

	if admin.ID == "" || admin.Email == "" {
		return &middleware.MiddlewareResponse{
			Message: "admin tidak terdaftar",
			Status:  404,
			Is_pass: false,
			Reason:  "admin tidak ada pada database",
		}
	}

	return &middleware.MiddlewareResponse{
		Message: "admin terauhtorisasi",
		Status:  200,
		Is_pass: true,
		Reason:  "terauthorisasi dengan sukses",
		Claim: struct {
			Id    string "json:\"id,omitempty\""
			Email string "json:\"email,omitempty\""
		}{
			Id:    admin.ID,
			Email: admin.Email,
		},
	}
}

func Middleware(repo *admin.AdminRepository, jwt *jsonwebtoken.JWTContract) middleware.Contract {
	return &AdminMiddleware{
		AdminRepository: *repo,
		JWT:             *jwt,
	}
}
