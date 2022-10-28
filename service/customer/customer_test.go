package customer

import (
	"errors"
	codegenerator "kang-sayur-backend/infrastructure/code_generator"
	"kang-sayur-backend/infrastructure/encryption"
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/infrastructure/identifier"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/mailer"
	"kang-sayur-backend/model/domain/customer"
	forgotpassword "kang-sayur-backend/model/domain/forgot_password"
	"kang-sayur-backend/model/domain/verification"
	request_body "kang-sayur-backend/model/web/request_body/customer"
	request_body_verification "kang-sayur-backend/model/web/request_body/verification"

	"testing"

	"github.com/stretchr/testify/mock"
)

var (
	customerRepository       = customer.MockRepository{Mock: mock.Mock{}}
	forgotPasswordRepository = forgotpassword.MockRepository{Mock: mock.Mock{}}
	verificationRepository   = verification.MockRepository{Mock: mock.Mock{}}
	jwt                      = jsonwebtoken.MockContract{Mock: mock.Mock{}}
	bcrypt                   = encryption.MockContract{Mock: mock.Mock{}}
	email                    = mailer.MockContract{Mock: mock.Mock{}}
	uuid                     = identifier.MockContract{Mock: mock.Mock{}}
	codeGen                  = codegenerator.MockContract{Mock: mock.Mock{}}
	service                  = customerService{
		customerRepository:       &customerRepository,
		forgotPasswordRepository: &forgotPasswordRepository,
		verificationRepository:   &verificationRepository,
		jwt:                      &jwt,
		bcrypt:                   &bcrypt,
		mailer:                   &email,
		uuid:                     &uuid,
		codeGenerator:            &codeGen,
	}
)

func TestForgotPassword(t *testing.T) {
	t.Run("customer not registered", func(t *testing.T) {
		expected := helper.Expected{
			Message: "pengguna tidak terdaftar",
			Status:  404,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{}).Once()

		r := service.ForgotPassword(&request_body.ForgotPassword{Email: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &expected, r, helper.DefaultOption)
	})

	t.Run("error when store FP", func(t *testing.T) {
		expected := helper.Expected{
			Message: "kesalahan keamanan",
			Status:  500,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{Email: mock.Anything, ID: mock.Anything}).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		forgotPasswordRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		r := service.ForgotPassword(&request_body.ForgotPassword{Email: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &expected, r, helper.DefaultOption)
	})

	t.Run("error when send FP email", func(t *testing.T) {
		expected := helper.Expected{
			Message: "kesalahan pengiriman akses",
			Status:  500,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{Email: mock.Anything, ID: mock.Anything}).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		forgotPasswordRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		email.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		r := service.ForgotPassword(&request_body.ForgotPassword{Email: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &expected, r, helper.DefaultOption)
	})

	t.Run("success FP operation", func(t *testing.T) {
		expected := helper.Expected{
			Message:   "request forgot password berhasil",
			Status:    200,
			IsSuccess: true,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{Email: mock.Anything, ID: mock.Anything}).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		forgotPasswordRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		email.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		r := service.ForgotPassword(&request_body.ForgotPassword{Email: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &expected, r, &helper.Options{DataChecking: false, TokenChecking: true})
	})
}

func TestReadAll(t *testing.T) {
	t.Run("empty data", func(t *testing.T) {
		e := helper.Expected{
			Message:   "data pelanggan kosong",
			Status:    200,
			IsSuccess: true,
		}

		customerRepository.Mock.On("Read").Return([]customer.Customer{}).Once()

		a := service.ReadAll()

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("data exists", func(t *testing.T) {
		e := helper.Expected{
			Message:   "data pelanggan berhasil di ambil",
			Status:    200,
			IsSuccess: true,
		}

		customerRepository.Mock.On("Read").Return([]customer.Customer{
			{
				ID:                 mock.Anything,
				Name:               mock.Anything,
				Phone:              mock.Anything,
				Email:              mock.Anything,
				Password:           mock.Anything,
				VerificationStatus: mock.Anything,
				DomicileID:         mock.Anything,
			},
		}).Once()

		a := service.ReadAll()

		helper.UnitTesting().CommonAssertion(t, &e, a, &helper.Options{DataChecking: true})
	})
}

func TestRegister(t *testing.T) {
	t.Run("user already registered", func(t *testing.T) {
		e := helper.Expected{
			Message: "data pelanggan sudah terdaftar",
			Status:  409,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{ID: mock.Anything, Email: mock.Anything}).Once()

		a := service.Register(&request_body.Register{
			Name:                 mock.Anything,
			Phone:                mock.Anything,
			Email:                mock.Anything,
			Password:             mock.Anything,
			PasswordConfirmation: mock.Anything,
			DomicileID:           mock.Anything,
		})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("error when create user", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{}).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		customerRepository.Mock.On("Create", mock.Anything).Return(&customer.Customer{}, errors.New(mock.Anything)).Once()

		a := service.Register(&request_body.Register{
			Name:                 mock.Anything,
			Phone:                mock.Anything,
			Email:                mock.Anything,
			Password:             mock.Anything,
			PasswordConfirmation: mock.Anything,
			DomicileID:           mock.Anything,
		})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("error when create verification", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{}).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		customerRepository.Mock.On("Create", mock.Anything).Return(&customer.Customer{}, nil).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		verificationRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.Register(&request_body.Register{
			Name:                 mock.Anything,
			Phone:                mock.Anything,
			Email:                mock.Anything,
			Password:             mock.Anything,
			PasswordConfirmation: mock.Anything,
			DomicileID:           mock.Anything,
		})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("error when send verification", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan pengiriman",
			Status:  500,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{}).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		customerRepository.Mock.On("Create", mock.Anything).Return(&customer.Customer{}, nil).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		verificationRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		email.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.Register(&request_body.Register{
			Name:                 mock.Anything,
			Phone:                mock.Anything,
			Email:                mock.Anything,
			Password:             mock.Anything,
			PasswordConfirmation: mock.Anything,
			DomicileID:           mock.Anything,
		})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("success registration", func(t *testing.T) {
		e := helper.Expected{
			Message:   "registrasi berhasil",
			Status:    200,
			IsSuccess: true,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{}).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		customerRepository.Mock.On("Create", mock.Anything).Return(&customer.Customer{}, nil).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		verificationRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		email.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		a := service.Register(&request_body.Register{
			Name:                 mock.Anything,
			Phone:                mock.Anything,
			Email:                mock.Anything,
			Password:             mock.Anything,
			PasswordConfirmation: mock.Anything,
			DomicileID:           mock.Anything,
		})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})
}

func TestRequestVerification(t *testing.T) {
	t.Run("user not registered", func(t *testing.T) {
		e := helper.Expected{
			Message: "pengguna tidak terdaftar",
			Status:  404,
		}

		customerRepository.Mock.On("ReadByID", mock.Anything).Return(&customer.Customer{}).Once()

		a := service.RequestVerification(&request_body_verification.RequestVerification{CustomerID: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("error when creating verification", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}

		customerRepository.Mock.On("ReadByID", mock.Anything).Return(&customer.Customer{ID: mock.Anything, Email: mock.Anything}).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		verificationRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.RequestVerification(&request_body_verification.RequestVerification{CustomerID: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("error when send verification", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan pengiriman",
			Status:  500,
		}

		customerRepository.Mock.On("ReadByID", mock.Anything).Return(&customer.Customer{ID: mock.Anything, Email: mock.Anything}).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		verificationRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		email.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.RequestVerification(&request_body_verification.RequestVerification{CustomerID: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("success operation", func(t *testing.T) {
		e := helper.Expected{
			Message:   "request verifikasi berhasil",
			Status:    200,
			IsSuccess: true,
		}

		customerRepository.Mock.On("ReadByID", mock.Anything).Return(&customer.Customer{ID: mock.Anything, Email: mock.Anything}).Once()
		codeGen.Mock.On("Generate").Return(mock.Anything).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		uuid.Mock.On("GenerateID").Return(mock.Anything).Once()
		verificationRepository.Mock.On("Upsert", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		email.Mock.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		a := service.RequestVerification(&request_body_verification.RequestVerification{CustomerID: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})
}

func TestResetPassword(t *testing.T) {
	t.Run("fp data not found", func(t *testing.T) {
		e := helper.Expected{
			Message: "sesi reset password tidak valid",
			Status:  440,
		}

		forgotPasswordRepository.Mock.On("ReadOne", mock.Anything).Return(&forgotpassword.ForgotPassword{}).Once()

		a := service.ResetPassword(&request_body.ResetPassword{Token: mock.Anything, Code: mock.Anything, Password: mock.Anything, PasswordConfirmation: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("wrong reset code", func(t *testing.T) {
		e := helper.Expected{
			Message: "kode reset password salah",
			Status:  401,
		}

		forgotPasswordRepository.Mock.On("ReadOne", mock.Anything).Return(&forgotpassword.ForgotPassword{
			ID:     mock.Anything,
			Token:  mock.Anything,
			UserID: mock.Anything,
			Code:   mock.Anything,
		}).Once()
		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(false).Once()

		a := service.ResetPassword(&request_body.ResetPassword{Token: mock.Anything, Code: mock.Anything, Password: mock.Anything, PasswordConfirmation: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("error when update password", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan update password",
			Status:  500,
		}

		forgotPasswordRepository.Mock.On("ReadOne", mock.Anything).Return(&forgotpassword.ForgotPassword{
			ID:     mock.Anything,
			Token:  mock.Anything,
			UserID: mock.Anything,
			Code:   mock.Anything,
		}).Once()
		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		customerRepository.Mock.On("UpdatePassword", mock.Anything, mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.ResetPassword(&request_body.ResetPassword{Token: mock.Anything, Code: mock.Anything, Password: mock.Anything, PasswordConfirmation: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("error when delete reset session password", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan penghapusan sesi reset password",
			Status:  500,
		}

		forgotPasswordRepository.Mock.On("ReadOne", mock.Anything).Return(&forgotpassword.ForgotPassword{
			ID:     mock.Anything,
			Token:  mock.Anything,
			UserID: mock.Anything,
			Code:   mock.Anything,
		}).Once()
		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		customerRepository.Mock.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		forgotPasswordRepository.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.ResetPassword(&request_body.ResetPassword{Token: mock.Anything, Code: mock.Anything, Password: mock.Anything, PasswordConfirmation: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("error when delete reset session password", func(t *testing.T) {
		e := helper.Expected{
			Message:   "reset password berhasil",
			Status:    200,
			IsSuccess: true,
		}

		forgotPasswordRepository.Mock.On("ReadOne", mock.Anything).Return(&forgotpassword.ForgotPassword{
			ID:     mock.Anything,
			Token:  mock.Anything,
			UserID: mock.Anything,
			Code:   mock.Anything,
		}).Once()
		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()
		bcrypt.Mock.On("Hash", mock.Anything).Return(mock.Anything).Once()
		customerRepository.Mock.On("UpdatePassword", mock.Anything, mock.Anything).Return(nil).Once()
		forgotPasswordRepository.Mock.On("Delete", mock.Anything).Return(nil).Once()

		a := service.ResetPassword(&request_body.ResetPassword{Token: mock.Anything, Code: mock.Anything, Password: mock.Anything, PasswordConfirmation: mock.Anything})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})
}

func TestUpdateProfile(t *testing.T) {
	t.Run("error when updating profile", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan update data",
			Status:  500,
		}

		customerRepository.Mock.On("UpdateProfile", mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.UpdateProfile(&request_body.UpdateProfile{
			CustomerID: mock.Anything,
			Name:       mock.Anything,
			Phone:      mock.Anything,
			DomicileID: mock.Anything,
		})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("success update operation", func(t *testing.T) {
		e := helper.Expected{
			Message:   "profile berhasil diupdate",
			Status:    200,
			IsSuccess: true,
		}

		customerRepository.Mock.On("UpdateProfile", mock.Anything).Return(nil).Once()

		a := service.UpdateProfile(&request_body.UpdateProfile{
			CustomerID: mock.Anything,
			Name:       mock.Anything,
			Phone:      mock.Anything,
			DomicileID: mock.Anything,
		})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})
}

func TestVerifyVerification(t *testing.T) {
	t.Run("Verification data not exists", func(t *testing.T) {
		e := helper.Expected{
			Message: "sesi verifikasi tidak valid",
			Status:  440,
		}

		verificationRepository.Mock.On("ReadOne", mock.Anything, mock.Anything).Return(&verification.Verification{}).Once()

		a := service.VerifyVerification(&request_body_verification.Verify{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("wrong verification data", func(t *testing.T) {
		e := helper.Expected{
			Message: "kode verifikasi salah",
			Status:  401,
		}

		verificationRepository.Mock.On("ReadOne", mock.Anything, mock.Anything).Return(&verification.Verification{ID: mock.Anything, Token: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(false).Once()

		a := service.VerifyVerification(&request_body_verification.Verify{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("failed to verify account", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan saat verifikasi",
			Status:  500,
		}

		verificationRepository.Mock.On("ReadOne", mock.Anything, mock.Anything).Return(&verification.Verification{ID: mock.Anything, Token: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		customerRepository.Mock.On("VerifyVerification", mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.VerifyVerification(&request_body_verification.Verify{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("failed to delete verification session", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan saat verifikasi tahap 2",
			Status:  500,
		}

		verificationRepository.Mock.On("ReadOne", mock.Anything, mock.Anything).Return(&verification.Verification{ID: mock.Anything, Token: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		customerRepository.Mock.On("VerifyVerification", mock.Anything).Return(nil).Once()

		verificationRepository.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.VerifyVerification(&request_body_verification.Verify{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("failed to delete verification session", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan saat verifikasi tahap 2",
			Status:  500,
		}

		verificationRepository.Mock.On("ReadOne", mock.Anything, mock.Anything).Return(&verification.Verification{ID: mock.Anything, Token: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		customerRepository.Mock.On("VerifyVerification", mock.Anything).Return(nil).Once()

		verificationRepository.Mock.On("Delete", mock.Anything).Return(errors.New(mock.Anything)).Once()

		a := service.VerifyVerification(&request_body_verification.Verify{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("success verification operation", func(t *testing.T) {
		e := helper.Expected{
			Message:   "akun berhasil diverifikasi",
			Status:    200,
			IsSuccess: true,
		}

		verificationRepository.Mock.On("ReadOne", mock.Anything, mock.Anything).Return(&verification.Verification{ID: mock.Anything, Token: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		customerRepository.Mock.On("VerifyVerification", mock.Anything).Return(nil).Once()

		verificationRepository.Mock.On("Delete", mock.Anything).Return(nil).Once()

		a := service.VerifyVerification(&request_body_verification.Verify{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})
}

func TestViewProfile(t *testing.T) {
	t.Run("customer not found", func(t *testing.T) {
		e := helper.Expected{
			Message: "pengguna tidak ditemukan",
			Status:  404,
		}

		customerRepository.Mock.On("ReadByID", mock.Anything).Return(&customer.Customer{}).Once()

		a := service.ViewProfile(&request_body.ViewProfile{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("customer found", func(t *testing.T) {
		e := helper.Expected{
			Message:   "profile berhasil di ambil",
			Status:    200,
			IsSuccess: true,
		}

		customerRepository.Mock.On("ReadByID", mock.Anything).Return(&customer.Customer{ID: mock.Anything}).Once()

		a := service.ViewProfile(&request_body.ViewProfile{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})
}

func TestLogin(t *testing.T) {
	t.Run("customer not registered", func(t *testing.T) {
		e := helper.Expected{
			Message: "pelanggan tidak terdaftar",
			Status:  409,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{}).Once()

		a := service.Login(&request_body.Login{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("wrong password", func(t *testing.T) {
		e := helper.Expected{
			Message: "password salah",
			Status:  401,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{ID: mock.Anything, Email: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(false).Once()

		a := service.Login(&request_body.Login{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("failed to create/sign JWT token", func(t *testing.T) {
		e := helper.Expected{
			Message: "kesalahan pembuatan akses",
			Status:  500,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{ID: mock.Anything, Email: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		jwt.Mock.On("Sign", mock.Anything).Return(mock.Anything, errors.New(mock.Anything)).Once()

		a := service.Login(&request_body.Login{})

		helper.UnitTesting().CommonAssertion(t, &e, a, helper.DefaultOption)
	})

	t.Run("success login operation", func(t *testing.T) {
		e := helper.Expected{
			Message:   "login berhasil",
			Status:    200,
			IsSuccess: true,
		}

		customerRepository.Mock.On("ReadByEmail", mock.Anything).Return(&customer.Customer{ID: mock.Anything, Email: mock.Anything}).Once()

		bcrypt.Mock.On("Compare", mock.Anything, mock.Anything).Return(true).Once()

		jwt.Mock.On("Sign", mock.Anything).Return(mock.Anything, nil).Once()

		a := service.Login(&request_body.Login{})

		helper.UnitTesting().CommonAssertion(t, &e, a, &helper.Options{DataChecking: true, TokenChecking: true})

	})
}
