package customer

import (
	"errors"
	"kang-sayur-backend/infrastructure/helper"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/middleware"
	"kang-sayur-backend/model/domain/customer"
	"kang-sayur-backend/model/domain/verification"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	jwt               = jsonwebtoken.MockContract{Mock: mock.Mock{}}
	verificationRepo  = verification.MockRepository{Mock: mock.Mock{}}
	customerRepo      = customer.MockRepository{Mock: mock.Mock{}}
	middlewareService = CustomerMiddleware{
		JWT:                    &jwt,
		VerificationRepository: &verificationRepo,
		Customer:               &customerRepo,
	}
)

func TestProcess(t *testing.T) {
	t.Run("empty token", func(t *testing.T) {
		result := middlewareService.Process("")
		expected := middleware.MiddlewareResponse{
			Message: "silahkan login terlebih dahulu",
			Status:  401,
			Is_pass: false,
			Reason:  "token kosong",
		}

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("error while parsing token", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "kesalahan saat parsing token",
			Status:  500,
			Is_pass: false,
			Reason:  "error parsing token",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{}, errors.New(mock.Anything)).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("customer not registered", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "pelanggan tidak terdaftar",
			Status:  500,
			Is_pass: false,
			Reason:  "data pelanggan tidak ada di database",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		customerRepo.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{}).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("customer not verified", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "pelanggan tidak terverifikasi",
			Status:  500,
			Is_pass: false,
			Reason:  "data pelanggan belum tidak terverifikasi di database",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		customerRepo.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{
			ID:                 mock.Anything,
			Name:               mock.Anything,
			Phone:              mock.Anything,
			Email:              mock.Anything,
			Password:           mock.Anything,
			VerificationStatus: "not verified",
			DomicileID:         mock.Anything,
		}).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("success middleware operation", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "customer terauhtorisasi",
			Status:  200,
			Is_pass: true,
			Reason:  "terauthorisasi dengan sukses",
			Claim: struct {
				Id    string "json:\"id,omitempty\""
				Email string "json:\"email,omitempty\""
			}{
				Id:    mock.Anything,
				Email: mock.Anything,
			},
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		customerRepo.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{
			ID:                 mock.Anything,
			Name:               mock.Anything,
			Phone:              mock.Anything,
			Email:              mock.Anything,
			Password:           mock.Anything,
			VerificationStatus: "verified",
			DomicileID:         mock.Anything,
		}).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})
}
