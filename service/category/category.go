package service

import (
	"fmt"
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/category"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/category"
)

type categoryService struct {
	repository category.CategoryRepository
}

// Create implements category.CategoryService
func (cs *categoryService) Create(body request_body.CreateCategory) response.HTTPResponse {
	inserted, err := cs.repository.Create(body)

	if err != nil {
		return response.HTTPResponse{
			Message: "kesalahan saat menyimpan data",
			Status:  500,
		}
	}

	return response.HTTPResponse{
		Message:   "kategori berhasil disimpan",
		Status:    200,
		IsSuccess: true,
		Data:      inserted,
	}
}

// Delete implements category.CategoryService
func (cs *categoryService) Delete(body request_body.UpdateOrDeleteCategory) response.HTTPResponse {
	if err := cs.repository.Delete(body); err != nil {
		return response.HTTPResponse{
			Message: "kesalahan saat pengapusan data",
			Status:  500,
		}
	}

	return response.HTTPResponse{
		Message:   "kategori berhasil dihapus",
		Status:    200,
		IsSuccess: true,
	}
}

// Read implements category.CategoryService
func (cs *categoryService) Read() response.HTTPResponse {
	categories, err := cs.repository.Read()

	if err != nil {
		fmt.Println(err)
		return response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return response.HTTPResponse{
		Message:   "kategori berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      categories,
	}
}

// Update implements category.CategoryService
func (cs *categoryService) Update(body request_body.UpdateOrDeleteCategory) response.HTTPResponse {
	if err := cs.repository.Update(body); err != nil {
		return response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}
	}

	return response.HTTPResponse{
		Message:   "kategori berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

func CategoryService(infra infrastructures.Infrastructures) category.CategoryService {
	return &categoryService{repository: infra.Repositories().Category}
}
