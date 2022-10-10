package admin

import (
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/middleware"
	"kang-sayur-backend/model/domain/admin"
)

type adminMiddleware struct {
	adminRepository admin.AdminRepository
	jwt             jsonwebtoken.JWTContract
}

// Process implements middleware.Contract
func (am *adminMiddleware) Process(token string) *middleware.MiddlewareResponse {
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
	payload, err := am.jwt.ParseToken(token)
	if err != nil {
		return &middleware.MiddlewareResponse{
			Message: "kesalahan saat parsing token",
			Status:  500,
			Reason:  "error",
		}
	}

	email := payload["email"].(string)

	// check admin by email
	admin := am.adminRepository.CheckEmail(email)

	if admin.ID == "" || admin.Email == "" {
		return &middleware.MiddlewareResponse{
			Message: "admin tidak terdaftar",
			Status:  404,
			Is_pass: false,
			Reason:  "admin tidak terdaftar",
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

func AdminMiddleware(repo *admin.AdminRepository, jwt *jsonwebtoken.JWTContract) middleware.Contract {
	return &adminMiddleware{
		adminRepository: *repo,
		jwt:             *jwt,
	}
}
