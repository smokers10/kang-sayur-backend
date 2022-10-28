package service

import (
	"fmt"
	"kang-sayur-backend/infrastructure/encryption"
	"kang-sayur-backend/infrastructure/identifier"
	infrastructures "kang-sayur-backend/infrastructure/injector"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/mailer"
	"kang-sayur-backend/model/domain/admin"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/admin"
)

type adminService struct {
	adminRepository admin.AdminRepository
	jwt             jsonwebtoken.JWTContract
	encrypt         encryption.EncryptionContract
	mailer          mailer.Contract
	identifier      identifier.IdentifierContract
}

// Login implements admin.AdminService
func (s *adminService) Login(body request_body.Login) response.HTTPResponse {
	// retrieve data admin
	admin := s.adminRepository.CheckEmail(body.Email)

	// check admin existance on database
	if admin.ID == "" || admin.Email == "" {
		return response.HTTPResponse{
			Message: "admin tidak terdaftar",
			Status:  404,
		}
	}

	// compare admin password
	if !s.encrypt.Compare(body.Password, admin.Password) {
		return response.HTTPResponse{
			Message: "password salah",
			Status:  401,
		}
	}

	// generate jwt token
	payload := map[string]interface{}{
		"id":    admin.ID,
		"email": admin.Email,
	}
	token, err := s.jwt.Sign(payload)
	if err != nil {
		return response.HTTPResponse{
			Message: "kesalahan authorisasi",
			Status:  500,
		}
	}

	// empty admin password
	admin.Password = ""

	return response.HTTPResponse{
		Message:   "login berhasil",
		Status:    200,
		IsSuccess: true,
		Data:      admin,
		Token:     token,
	}
}

// RequestLogin implements admin.AdminService
func (s *adminService) RequestLogin(body request_body.LoginRequest) response.HTTPResponse {
	admin := s.adminRepository.CheckEmail(body.Email)
	if admin.ID == "" {
		return response.HTTPResponse{
			Message: "admin tidak terdaftar",
			Status:  404,
		}
	}

	// create password
	unsafePassword := s.identifier.GenerateID()
	safePassword := s.encrypt.Hash(unsafePassword)

	// update password
	if err := s.adminRepository.UpdatePassword(admin.ID, safePassword); err != nil {
		fmt.Print(err)
		return response.HTTPResponse{
			Message: "kesalahan saat request",
			Status:  500,
		}
	}

	// send unsafe password to handler email
	if err := s.mailer.Send([]string{body.Email}, "akses admin", AdminPasswordEmail(unsafePassword)); err != nil {
		return response.HTTPResponse{
			Message: "kesalahan saat kirim akses",
			Status:  500,
		}
	}

	// success response
	return response.HTTPResponse{
		Message:   "request diterima",
		Status:    200,
		IsSuccess: true,
	}
}

func AdminService(infra *infrastructures.Infrastructures) admin.AdminService {
	return &adminService{
		adminRepository: infra.Repositories().Admin,
		jwt:             *infra.JsonWebToken(),
		encrypt:         *infra.Encryption(),
		mailer:          *infra.Mailer(),
		identifier:      *infra.Identifier(),
	}
}
