package service

import (
	"errors"
	mailer "kang-sayur-backend/infrastructure/SMTP"
	"kang-sayur-backend/infrastructure/encryption"
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/infrastructure/identifier"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/model/domain/admin"
	response "kang-sayur-backend/model/web"
	requestbody "kang-sayur-backend/model/web/request_body/admin"
	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	adminRepository = admin.MockRepository{Mock: mock.Mock{}}
	jwt             = jsonwebtoken.MockContract{Mock: mock.Mock{}}
	encrypt         = encryption.MockContract{Mock: mock.Mock{}}
	smtp            = mailer.MockContract{Mock: mock.Mock{}}
	uuid            = identifier.MockContract{Mock: mock.Mock{}}
	service         = adminService{
		adminRepository: &adminRepository,
		jwt:             &jwt,
		encrypt:         &encrypt,
		mailer:          &smtp,
		identifier:      &uuid,
	}
)

func TestRequestLogin(t *testing.T) {
	t.Run("admin not registered", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "admin tidak terdaftar",
			Status:  404,
		}

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{}).Once()
		result := service.RequestLogin(requestbody.LoginRequest{Email: mock.Anything})
		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})

	t.Run("failed to update password", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat request",
			Status:  500,
		}

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{
			ID:       mock.Anything,
			Name:     mock.Anything,
			Email:    mock.Anything,
			Password: mock.Anything,
		}).Once()

		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		encrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()

		adminRepository.Mock.On("UpdatePassword", mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		result := service.RequestLogin(requestbody.LoginRequest{Email: mock.Anything})
		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})

	t.Run("failed to send access", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message: "kesalahan saat kirim akses",
			Status:  500,
		}

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{
			ID:       mock.Anything,
			Name:     mock.Anything,
			Email:    mock.Anything,
			Password: mock.Anything,
		}).Once()

		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		encrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		adminRepository.Mock.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		smtp.Mock.On("Send", []string{mock.Anything}, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		result := service.RequestLogin(requestbody.LoginRequest{Email: mock.Anything})
		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})

	t.Run("request success", func(t *testing.T) {
		expected := response.HTTPResponse{
			Message:   "request diterima",
			Status:    200,
			IsSuccess: true,
		}

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{
			ID:       mock.Anything,
			Name:     mock.Anything,
			Email:    mock.Anything,
			Password: mock.Anything,
		}).Once()

		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		encrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		adminRepository.Mock.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		smtp.Mock.On("Send", []string{mock.Anything}, mock.Anything, mock.Anything).Return(nil).Once()

		result := service.RequestLogin(requestbody.LoginRequest{Email: mock.Anything})
		helper.UnitTesting().CommonAssertion(t, (*helper.Expected)(&expected), &result, helper.DefaultOption)
	})
}

func TestLogin(t *testing.T) {
	t.Run("admin not registered", func(t *testing.T) {
		expected := helper.Expected{
			Message: "admin tidak terdaftar",
			Status:  404,
		}

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{}).Once()

		result := service.Login(requestbody.Login{Email: mock.Anything, Password: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &expected, &result, helper.DefaultOption)
	})

	t.Run("password not match", func(t *testing.T) {
		expected := helper.Expected{
			Message: "password salah",
			Status:  401,
		}

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{
			ID:    mock.Anything,
			Email: mock.Anything,
		}).Once()

		encrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(false).Once()

		result := service.Login(requestbody.Login{Email: mock.Anything, Password: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &expected, &result, helper.DefaultOption)
	})

	t.Run("error when generate JWT token", func(t *testing.T) {
		expected := helper.Expected{
			Message: "kesalahan authorisasi",
			Status:  500,
		}

		adminRepository.Mock.On("CheckEmail", mock.Anything).Return(&admin.Admin{
			ID:    mock.Anything,
			Email: mock.Anything,
		}).Once()

		encrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		jwt.Mock.On("Sign", mock.Anything).Return(mock.Anything, errors.New(mock.Anything)).Once()

		result := service.Login(requestbody.Login{Email: mock.Anything, Password: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &expected, &result, helper.DefaultOption)
	})
}
