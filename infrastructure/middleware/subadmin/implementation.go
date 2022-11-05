package subadmin

import (
	"kang-sayur-backend/infrastructure/injector"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/middleware"
	"kang-sayur-backend/model/domain/permission"
	subadmin "kang-sayur-backend/model/domain/sub_admin"
	"strings"
)

type subadminMiddleware struct {
	JWT                  jsonwebtoken.JWTContract
	SubAdminRepository   subadmin.SubAdminRepository
	PermissionRepository permission.PermissionRepository
}

// Process implements middleware.Contract
func (sm *subadminMiddleware) Process(token string) *middleware.MiddlewareResponse {
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
	payload, err := sm.JWT.ParseToken(token)
	if err != nil {
		return &middleware.MiddlewareResponse{
			Message: "kesalahan saat parsing token",
			Status:  500,
			Is_pass: false,
			Reason:  "error parsing token",
		}
	}

	email := payload["email"].(string)
	id := payload["id"].(string)

	// retrieve sub admin by its email
	subadmin, err := sm.SubAdminRepository.ReadByEmail(email)
	if err != nil {
		return &middleware.MiddlewareResponse{
			Message: "kesalahan saat pengambilan data pengguna",
			Status:  500,
			Is_pass: false,
			Reason:  "error retrieve user data",
		}
	}

	// check if sub admin registered or not
	if subadmin.ID == "" || subadmin.Email == "" {
		return &middleware.MiddlewareResponse{
			Message: "sub admin tidak terdaftar",
			Status:  401,
			Is_pass: false,
			Reason:  "data sub admin tidak ada di database",
		}
	}

	// check if sub admin account blocked/supended or ot
	if strings.ToLower(subadmin.Status) == "blocked" {
		return &middleware.MiddlewareResponse{
			Message: "hak akses anda ditahan",
			Status:  401,
			Is_pass: false,
			Reason:  "status sub admin blocked",
		}
	}

	// retrieve permission
	permissions, err := sm.PermissionRepository.ReadOne(id)
	if err != nil {
		return &middleware.MiddlewareResponse{
			Message: "kesalahan saat pengambilan hak akses",
			Status:  500,
			Is_pass: false,
			Reason:  "error retrieve permission data",
		}
	}

	if permissions.ID == "" {
		return &middleware.MiddlewareResponse{
			Message: "hak akses belum ada",
			Status:  404,
			Is_pass: false,
			Reason:  "permission not set",
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
			Id:    id,
			Email: email,
		},
		Permission: permissions,
	}
}

func Middleware(infrastructure *injector.Infrastructures) middleware.Contract {
	return &subadminMiddleware{
		JWT:                  *infrastructure.JsonWebToken(),
		SubAdminRepository:   infrastructure.Repositories().SubAdmin,
		PermissionRepository: infrastructure.Repositories().Permission,
	}
}
