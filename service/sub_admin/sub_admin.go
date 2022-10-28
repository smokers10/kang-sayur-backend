package service

import (
	"kang-sayur-backend/infrastructure/injector"
	subadmin "kang-sayur-backend/model/domain/sub_admin"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/sub_admin"
)

type subAdminService struct {
	repository subadmin.SubAdminRepository
}

// Login implements subadmin.SubAdminService
func (*subAdminService) Login(body *request_body.Login) response.HTTPResponse {
	panic("unimplemented")
}

// Create implements subadmin.SubAdminService
func (*subAdminService) Create(body *request_body.Create) response.HTTPResponse {
	panic("unimplemented")
}

// Delete implements subadmin.SubAdminService
func (*subAdminService) Delete(body *request_body.Delete) response.HTTPResponse {
	panic("unimplemented")
}

// Read implements subadmin.SubAdminService
func (*subAdminService) Read() response.HTTPResponse {
	panic("unimplemented")
}

// SetPermission implements subadmin.SubAdminService
func (*subAdminService) SetPermission(body *request_body.SetPermission) response.HTTPResponse {
	panic("unimplemented")
}

// Update implements subadmin.SubAdminService
func (*subAdminService) Update(body *request_body.Update) response.HTTPResponse {
	panic("unimplemented")
}

// UpdateStatus implements subadmin.SubAdminService
func (*subAdminService) UpdateStatus(body *request_body.UpdateStatus) response.HTTPResponse {
	panic("unimplemented")
}

func SubAdminService(infra injector.Infrastructures) subadmin.SubAdminService {
	return &subAdminService{repository: infra.Repositories().SubAdmin}
}
