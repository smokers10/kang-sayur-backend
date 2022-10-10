package admin

import (
	"errors"
	"kang-sayur-backend/infrastructure/helper"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/middleware"
	"kang-sayur-backend/model/domain/admin"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	adminRepository   = admin.MockRepository{Mock: mock.Mock{}}
	jwt               = jsonwebtoken.MockContract{Mock: mock.Mock{}}
	middlewareService = AdminMiddleware{
		AdminRepository: &adminRepository,
		JWT:             &jwt,
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

	t.Run("admin not registered", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "admin tidak terdaftar",
			Status:  404,
			Is_pass: false,
			Reason:  "admin tidak ada pada database",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{}).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("success middleware operation", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "admin terauhtorisasi",
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

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{ID: mock.Anything, Email: mock.Anything}).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})
}
