package service

import (
	"errors"
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/model/domain/cart"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/cart"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	repository = cart.MockRepository{Mock: mock.Mock{}}
	service    = cartService{
		repository: &repository,
	}
)

func TestAdd(t *testing.T) {
	t.Run("error add", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "Kesalahan saat penyimpanan data",
			Status:  500,
		}

		repository.Mock.On("Create").Return(&cart.Cart{}, errors.New(mock.Anything)).Once()

		result := service.Add(&request_body.BasicAction{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success add", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "berhasil ditambahkan ke keranjang",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Create").Return(&cart.Cart{
			ID:         mock.Anything,
			CustomerID: mock.Anything,
			ProductID:  mock.Anything,
			Quantity:   1,
		}, nil).Once()

		result := service.Add(&request_body.BasicAction{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, &helper.Options{DataChecking: true})
	})
}

func TestDelete(t *testing.T) {
	t.Run("error delete", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "Kesalahan saat menghapus data",
			Status:  500,
		}

		repository.Mock.On("Delete").Return(errors.New(mock.Anything)).Once()

		result := service.Delete(&request_body.BasicAction{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success delete", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "Barang berhasil dihapus dari keranjang",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("Delete").Return(nil).Once()

		result := service.Delete(&request_body.BasicAction{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}

func TestGetCart(t *testing.T) {
	t.Run("error get cart", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "Kesalahan saat pengambilan data",
			Status:  500,
		}

		repository.Mock.On("ReadCart", mock.Anything).Return([]cart.Cart{}, errors.New(mock.Anything)).Once()

		result := service.GetCart(&request_body.ReadCart{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success get cart", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "keranjang berhasil diambil",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("ReadCart", mock.Anything).Return([]cart.Cart{}, nil).Once()

		result := service.GetCart(&request_body.ReadCart{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}

func TestUpdateQuantity(t *testing.T) {
	t.Run("error update quantity", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat update quantitas",
			Status:  500,
		}

		repository.Mock.On("UpdateQuantity", mock.Anything).Return(errors.New(mock.Anything)).Once()

		result := service.UpdateQuantity(&request_body.UpdateQuantity{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})

	t.Run("success update quantity", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "Quantitas berhasil diupdate",
			Status:    200,
			IsSuccess: true,
		}

		repository.Mock.On("UpdateQuantity", mock.Anything).Return(nil).Once()

		result := service.UpdateQuantity(&request_body.UpdateQuantity{})

		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), result, helper.DefaultOption)
	})
}
