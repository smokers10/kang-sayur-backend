package subadmin

import (
	"errors"
	"kang-sayur-backend/infrastructure/helper"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/middleware"
	"kang-sayur-backend/model/domain/permission"
	subadmin "kang-sayur-backend/model/domain/sub_admin"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	jwt                  = jsonwebtoken.MockContract{Mock: mock.Mock{}}
	subAdminRepository   = subadmin.MockRepository{Mock: mock.Mock{}}
	permissionRepository = permission.MockRepository{Mock: mock.Mock{}}
	middlewareService    = subadminMiddleware{
		JWT:                  &jwt,
		SubAdminRepository:   &subAdminRepository,
		PermissionRepository: &permissionRepository,
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

	t.Run("error when retrieve sub admin data", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "kesalahan saat pengambilan data pengguna",
			Status:  500,
			Is_pass: false,
			Reason:  "error retrieve user data",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		subAdminRepository.Mock.On("ReadByEmail", mock.Anything).Return(&subadmin.SubAdmin{}, errors.New(mock.Anything)).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("subadmin not registered", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "sub admin tidak terdaftar",
			Status:  401,
			Is_pass: false,
			Reason:  "data sub admin tidak ada di database",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		subAdminRepository.Mock.On("ReadByEmail", mock.Anything).Return(&subadmin.SubAdmin{}, nil).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("subadmin status is blocked", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "hak akses anda ditahan",
			Status:  401,
			Is_pass: false,
			Reason:  "status sub admin blocked",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		subAdminRepository.Mock.On("ReadByEmail", mock.Anything).Return(&subadmin.SubAdmin{
			ID:     mock.Anything,
			Email:  mock.Anything,
			Status: "Blocked",
		}, nil).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("error when retrieve sub admin permission", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "kesalahan saat pengambilan hak akses",
			Status:  500,
			Is_pass: false,
			Reason:  "error retrieve permission data",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		subAdminRepository.Mock.On("ReadByEmail", mock.Anything).Return(&subadmin.SubAdmin{
			ID:     mock.Anything,
			Email:  mock.Anything,
			Status: mock.Anything,
		}, nil).Once()

		permissionRepository.Mock.On("ReadOne", mock.Anything).Return(&permission.Permission{}, errors.New(mock.Anything)).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})

	t.Run("pemission not set", func(t *testing.T) {
		expected := middleware.MiddlewareResponse{
			Message: "hak akses belum ada",
			Status:  404,
			Is_pass: false,
			Reason:  "permission not set",
		}

		jwt.Mock.On("ParseToken", mock.Anything).Return(map[string]interface{}{
			"email": mock.Anything,
			"id":    mock.Anything,
		}, nil).Once()

		subAdminRepository.Mock.On("ReadByEmail", mock.Anything).Return(&subadmin.SubAdmin{
			ID:     mock.Anything,
			Email:  mock.Anything,
			Status: mock.Anything,
		}, nil).Once()

		permissionRepository.Mock.On("ReadOne", mock.Anything).Return(&permission.Permission{}, nil).Once()

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

		subAdminRepository.Mock.On("ReadByEmail", mock.Anything).Return(&subadmin.SubAdmin{
			ID:     mock.Anything,
			Email:  mock.Anything,
			Status: mock.Anything,
		}, nil).Once()

		permissionRepository.Mock.On("ReadOne", mock.Anything).Return(&permission.Permission{
			ID:             mock.Anything,
			MajBarang:      true,
			MajPenjualan:   true,
			MajPengguna:    true,
			MajPenggunaAdm: true,
			MajKeuangan:    true,
			SubAdminID:     mock.Anything,
		}, nil).Once()

		result := middlewareService.Process(mock.Anything)

		helper.UnitTesting().MiddlewareAssertion(t, &expected, result)
	})
}
