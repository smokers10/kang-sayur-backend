package service

import (
	"fmt"
	"kang-sayur-backend/infrastructure/encryption"
	"kang-sayur-backend/infrastructure/injector"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/model/domain/permission"
	subadmin "kang-sayur-backend/model/domain/sub_admin"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/sub_admin"
)

type subAdminService struct {
	repository           subadmin.SubAdminRepository
	permissionRepository permission.PermissionRepository
	jwt                  jsonwebtoken.JWTContract
	bcrypt               encryption.EncryptionContract
}

// Login implements subadmin.SubAdminService
func (sas *subAdminService) Login(body *request_body.Login) *response.HTTPResponse {
	subadmin, err := sas.repository.ReadByEmail(body.Email)

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan check data",
			Status:  500,
		}
	}

	// if sub admin not registered
	if subadmin.ID == "" || subadmin.Email == "" {
		return &response.HTTPResponse{
			Message: "sub admin tidak terdaftar",
			Status:  500,
		}
	}

	// compare password
	if !sas.bcrypt.Compare(body.Password, subadmin.Password) {
		return &response.HTTPResponse{
			Message: "password salah",
			Status:  401,
		}
	}

	// sign jwt
	token, err := sas.jwt.Sign(map[string]interface{}{
		"id":    subadmin.ID,
		"email": subadmin.Email,
	})

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan pembuatan token",
			Status:  500,
		}
	}

	// hide password
	subadmin.Password = ""

	return &response.HTTPResponse{
		Message:   "",
		Status:    200,
		IsSuccess: true,
		Data:      subadmin,
		Token:     token,
	}
}

// Create implements subadmin.SubAdminService
func (sas *subAdminService) Create(body *request_body.Create) *response.HTTPResponse {
	subadmin, err := sas.repository.ReadByEmail(body.Email)

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan check data",
			Status:  500,
		}
	}

	// if sub admin already registered
	if subadmin.ID != "" || subadmin.Email != "" {
		return &response.HTTPResponse{
			Message: "sub admin telah teregistrasi",
			Status:  500,
		}
	}

	// secure password
	body.Password = sas.bcrypt.Hash(body.Password)

	// store sub admin
	if err := sas.repository.Create(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "sub admin berhasil dibuat",
		Status:    200,
		IsSuccess: true,
	}
}

// Delete implements subadmin.SubAdminService
func (sas *subAdminService) Delete(body *request_body.Delete) *response.HTTPResponse {
	if err := sas.repository.Delete(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat pengapusan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "Sub admin berhasil dihapus",
		Status:    200,
		IsSuccess: true,
	}
}

// Read implements subadmin.SubAdminService
func (sas *subAdminService) Read() *response.HTTPResponse {
	subadmins, err := sas.repository.Read()

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "sub admin berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      subadmins,
	}
}

// SetPermission implements subadmin.SubAdminService
func (sas *subAdminService) SetPermission(body *request_body.SetPermission) *response.HTTPResponse {
	if err := sas.permissionRepository.Upsert(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "Hak akses berhasil ditetapkan",
		Status:    200,
		IsSuccess: true,
	}
}

// Update implements subadmin.SubAdminService
func (sas *subAdminService) Update(body *request_body.Update) *response.HTTPResponse {
	if err := sas.repository.Update(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "sub admin berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

// UpdateStatus implements subadmin.SubAdminService
func (sas *subAdminService) UpdateStatus(body *request_body.UpdateStatus) *response.HTTPResponse {
	if err := sas.repository.UpdateStatus(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "status sub admin berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

func SubAdminService(infra injector.Infrastructures) subadmin.SubAdminService {
	return &subAdminService{
		repository:           infra.Repositories().SubAdmin,
		permissionRepository: infra.Repositories().Permission,
		jwt:                  *infra.JsonWebToken(),
		bcrypt:               *infra.Encryption(),
	}
}
