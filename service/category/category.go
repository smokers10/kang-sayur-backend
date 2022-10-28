package service

import (
	infrastructures "kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/category"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/category"
)

type categoryService struct {
	repository category.CategoryRepository
}

// Create implements category.CategoryService
func (*categoryService) Create(body request_body.CreateCategory) response.HTTPResponse {
	panic("unimplemented")
}

// Delete implements category.CategoryService
func (*categoryService) Delete(body request_body.UpdateOrDeleteCategory) response.HTTPResponse {
	panic("unimplemented")
}

// Read implements category.CategoryService
func (*categoryService) Read() response.HTTPResponse {
	panic("unimplemented")
}

// Update implements category.CategoryService
func (*categoryService) Update(body request_body.UpdateOrDeleteCategory) response.HTTPResponse {
	panic("unimplemented")
}

func CategoryService(infra infrastructures.Infrastructures) category.CategoryService {
	return &categoryService{repository: infra.Repositories().Category}
}
