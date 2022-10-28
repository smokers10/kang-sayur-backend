package service

import (
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/address"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/address"
)

type addressService struct {
	repository address.AddressRepository
}

// Create implements address.AddressService
func (*addressService) Create(body *request_body.CreateAddress) *response.HTTPResponse {
	panic("unimplemented")
}

// Delelete implements address.AddressService
func (*addressService) Delelete(body *request_body.DeleteOrReadAddress) *response.HTTPResponse {
	panic("unimplemented")
}

// Read implements address.AddressService
func (*addressService) Read(body *request_body.DeleteOrReadAddress) *response.HTTPResponse {
	panic("unimplemented")
}

// Update implements address.AddressService
func (*addressService) Update(body *request_body.UpdateAddress) *response.HTTPResponse {
	panic("unimplemented")
}

func AddressService(infra *infrastructures.Infrastructures) address.AddressService {
	return &addressService{repository: infra.Repositories().Address}
}
