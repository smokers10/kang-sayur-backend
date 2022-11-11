package service

import (
	"errors"
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/model/domain/feedback"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/feedback"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	repository = feedback.MockRepository{Mock: mock.Mock{}}
	service    = feedbackService{
		repository: &repository,
	}
)

func TestCreate(t *testing.T) {
	t.Run("error create", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat menyimpan data",
			Status:  500,
		}

		repository.Mock.On("Create", mock.Anything).Return(&feedback.Feedback{}, errors.New(mock.Anything)).Once()

		result := service.Create(&request_body.Create{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success create", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "timbal balik berhasil disimpan",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Create", mock.Anything).Return(&feedback.Feedback{
			ID:         mock.Anything,
			CustomerID: mock.Anything,
			ProductID:  mock.Anything,
			Body:       mock.Anything,
			Rating:     mock.Anything,
		}, nil).Once()

		result := service.Create(&request_body.Create{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, &helper.Options{DataChecking: true})
	})
}

func TestDelete(t *testing.T) {
	t.Run("error delete", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat pengapusan data",
			Status:  500,
		}

		repository.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		result := service.Delete(&request_body.UpdateOrDelete{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success create", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "timbal balik berhasil dihapus",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Delete", mock.Anything).Return(nil).Once()

		result := service.Delete(&request_body.UpdateOrDelete{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}

func TestRead(t *testing.T) {
	t.Run("error read", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}

		repository.Mock.On("Read", mock.Anything).Return([]feedback.Feedback{}, errors.New(mock.Anything)).Once()

		result := service.Read(&request_body.ReadOnProduct{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success read", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "timbal balik berhasil diambil",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Read", mock.Anything).Return([]feedback.Feedback{
			{
				ID:         mock.Anything,
				CustomerID: mock.Anything,
				ProductID:  mock.Anything,
				Body:       mock.Anything,
				Rating:     mock.Anything,
			},
		}, nil).Once()

		result := service.Read(&request_body.ReadOnProduct{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, &helper.Options{DataChecking: true})
	})
}

func TestReadOne(t *testing.T) {
	t.Run("error read one", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}

		repository.Mock.On("ReadOne", mock.Anything).Return(&feedback.Feedback{}, errors.New(mock.Anything)).Once()

		result := service.ReadOne(&request_body.ReadOne{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success read one", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "timbal balik berhasil diambil",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("ReadOne", mock.Anything).Return(&feedback.Feedback{
			ID:         mock.Anything,
			CustomerID: mock.Anything,
			ProductID:  mock.Anything,
			Body:       mock.Anything,
			Rating:     mock.Anything,
		}, nil).Once()

		result := service.ReadOne(&request_body.ReadOne{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, &helper.Options{DataChecking: true})
	})
}

func TestUpdate(t *testing.T) {
	t.Run("error update", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}

		repository.Mock.On("Update", mock.Anything).Return(errors.New(mock.Anything)).Once()

		result := service.Update(&request_body.UpdateOrDelete{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success update", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "timbal balik berhasil diupdate",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Update", mock.Anything).Return(nil).Once()

		result := service.Update(&request_body.UpdateOrDelete{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}
