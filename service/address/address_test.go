package service

import (
	"errors"
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/model/domain/address"
	response "kang-sayur-backend/model/web"
	"testing"

	request_body "kang-sayur-backend/model/web/request_body/address"

	"github.com/stretchr/testify/mock"
)

var (
	repository = address.MockRepository{Mock: mock.Mock{}}
	service    = addressService{
		repository: &repository,
	}
)

func TestCreate(t *testing.T) {
	t.Run("error create", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}

		repository.Mock.On("Create", mock.Anything).Return(&address.Address{}, errors.New(mock.Anything)).Once()

		result := service.Create(&request_body.CreateAddress{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success create", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "alamat berhasil disimpan",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Create", mock.Anything).Return(&address.Address{}, nil).Once()

		result := service.Create(&request_body.CreateAddress{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}

func TestRead(t *testing.T) {
	t.Run("error read", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}

		repository.Mock.On("Read", mock.Anything).Return([]address.Address{}, errors.New(mock.Anything)).Once()

		result := service.Read(&request_body.DeleteOrReadAddress{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success read", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "alamat berhasil diambil",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Read", mock.Anything).Return([]address.Address{
			{
				ID:         mock.Anything,
				Name:       mock.Anything,
				Address:    mock.Anything,
				Status:     mock.Anything,
				CustomerID: mock.Anything,
			},
		}, nil).Once()

		result := service.Read(&request_body.DeleteOrReadAddress{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}

func TestDelete(t *testing.T) {
	t.Run("error delete", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan penghapusan data",
			Status:  500,
		}

		repository.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		result := service.Delete(&request_body.DeleteOrReadAddress{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success delete", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "alamat berhasil dihapus",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Delete", mock.Anything).Return(nil).Once()

		result := service.Delete(&request_body.DeleteOrReadAddress{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("error update", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat update",
			Status:  500,
		}

		repository.Mock.On("Update", mock.Anything).Return(&address.Address{}, errors.New(mock.Anything)).Once()

		result := service.Update(&request_body.UpdateAddress{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success update", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "alamat berhasil diupdate",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Update", mock.Anything).Return(&address.Address{}, nil).Once()

		result := service.Update(&request_body.UpdateAddress{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}
