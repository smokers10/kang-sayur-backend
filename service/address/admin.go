package service

import (
	encryption "kang-sayur-backend/infrastructure/encryption"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/model/domain/admin"
	subadmin "kang-sayur-backend/model/domain/sub_admin"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/admin"
)

type adminService struct {
	adminRepository    admin.AdminRepository
	subAdminRepository subadmin.SubAdminRepository
	jwt                jsonwebtoken.JWTContract
	encryption         encryption.EncryptionContract
}

// Login implements admin.AdminService
func (*adminService) Login(body request_body.Login) response.HTTPResponse {
	// check if admin exists or not
	panic("unimplemented")
}

// RequestLogin implements admin.AdminService
func (*adminService) RequestLogin(body request_body.LoginRequest) response.HTTPResponse {
	panic("unimplemented")
}

func AdminService(adminRepository *admin.AdminRepository, subAdminRepository *subadmin.SubAdminRepository, jwt *jsonwebtoken.JWTContract, encryption *encryption.EncryptionContract) admin.AdminService {
	return &adminService{
		adminRepository:    *adminRepository,
		subAdminRepository: *subAdminRepository,
		jwt:                *jwt,
		encryption:         *encryption,
	}
}
