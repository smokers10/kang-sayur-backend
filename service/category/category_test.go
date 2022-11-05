package service

import (
	"errors"
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/model/domain/category"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/category"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	repository = category.MockRepository{Mock: mock.Mock{}}
	service    = categoryService{
		repository: &repository,
	}
)

func TestCreate(t *testing.T) {
	t.Run("error create", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat menyimpan data",
			Status:  500,
		}

		repository.Mock.On("Create", mock.Anything).Return(&category.Category{}, errors.New(mock.Anything)).Once()

		result := service.Create(request_body.CreateCategory{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})

	t.Run("success create", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "kategori berhasil disimpan",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Create", mock.Anything).Return(&category.Category{}, nil).Once()

		result := service.Create(request_body.CreateCategory{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})
}

func TestDelete(t *testing.T) {
	t.Run("error delete", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat pengapusan data",
			Status:  500,
		}

		repository.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		result := service.Delete(request_body.UpdateOrDeleteCategory{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})

	t.Run("success delete", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "kategori berhasil dihapus",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Delete", mock.Anything).Return(nil).Once()

		result := service.Delete(request_body.UpdateOrDeleteCategory{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})
}

func TestRead(t *testing.T) {
	t.Run("error read", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}

		repository.Mock.On("Read").Return([]category.Category{}, errors.New(mock.Anything)).Once()

		result := service.Read()

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})

	t.Run("success read", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "kategori berhasil diambil",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Read").Return([]category.Category{}, nil).Once()

		result := service.Read()

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("error read", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}

		repository.Mock.On("Update", mock.Anything).Return(errors.New(mock.Anything)).Once()

		result := service.Update(request_body.UpdateOrDeleteCategory{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})

	t.Run("success read", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "kategori berhasil diupdate",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Update", mock.Anything).Return(nil).Once()

		result := service.Update(request_body.UpdateOrDeleteCategory{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})
}
