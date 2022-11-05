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
func (as *addressService) Create(body *request_body.CreateAddress) *response.HTTPResponse {
	inserted, err := as.repository.Create(body)
	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "alamat berhasil disimpan",
		Status:    200,
		IsSuccess: true,
		Data:      inserted,
	}
}

// Delelete implements address.AddressService
func (as *addressService) Delete(body *request_body.DeleteOrReadAddress) *response.HTTPResponse {
	if err := as.repository.Delete(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan penghapusan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "alamat berhasil dihapus",
		Status:    200,
		IsSuccess: true,
	}
}

// Read implements address.AddressService
func (as *addressService) Read(body *request_body.DeleteOrReadAddress) *response.HTTPResponse {
	addresses, err := as.repository.Read(body)

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "alamat berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      addresses,
	}
}

// Update implements address.AddressService
func (as *addressService) Update(body *request_body.UpdateAddress) *response.HTTPResponse {
	address, err := as.repository.Update(body)

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "alamat berhasil diupdate",
		Status:    200,
		IsSuccess: true,
		Data:      address,
	}
}

func AddressService(infra *infrastructures.Infrastructures) address.AddressService {
	return &addressService{repository: infra.Repositories().Address}
}
