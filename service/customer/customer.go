package customer

import (
	codegenerator "kang-sayur-backend/infrastructure/code_generator"
	"kang-sayur-backend/infrastructure/encryption"
	"kang-sayur-backend/infrastructure/identifier"
	infrastructures "kang-sayur-backend/infrastructure/injector"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	"kang-sayur-backend/infrastructure/mailer"
	"kang-sayur-backend/model/domain/customer"
	forgotpassword "kang-sayur-backend/model/domain/forgot_password"
	"kang-sayur-backend/model/domain/verification"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/customer"
	verification_request_body "kang-sayur-backend/model/web/request_body/verification"
)

type customerService struct {
	customerRepository       customer.CustomerRepository
	forgotPasswordRepository forgotpassword.ForgotPasswordRepository
	verificationRepository   verification.VerificationRepository
	jwt                      jsonwebtoken.JWTContract
	bcrypt                   encryption.EncryptionContract
	mailer                   mailer.Contract
	uuid                     identifier.IdentifierContract
	codeGenerator            codegenerator.CodeGeneratorContract
}

// ForgotPassword implements customer.CustomerService
func (cs *customerService) ForgotPassword(body *request_body.ForgotPassword) *response.HTTPResponse {
	// retrieve customer by its email
	customer := cs.customerRepository.ReadByEmail(body.Email)

	// check if customer exists or not
	if customer.ID == "" || customer.Email == "" {
		return &response.HTTPResponse{
			Message: "pengguna tidak terdaftar",
			Status:  404,
		}
	}

	// generate Code
	otp := cs.codeGenerator.Generate()
	secureOTP := cs.bcrypt.Hash(otp)

	// generate token
	token := cs.uuid.GenerateID()

	// store forgot password
	if err := cs.forgotPasswordRepository.Upsert(token, customer.ID, secureOTP); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan keamanan",
			Status:  500,
		}
	}

	// send to email
	if err := cs.mailer.Send([]string{customer.Email}, "Reset Password", forgotPasswordEmail(otp, customer.Name)); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan pengiriman akses",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "request forgot password berhasil",
		Status:    200,
		IsSuccess: true,
		Token:     token,
	}
}

// ReadAll implements customer.CustomerService
func (cs *customerService) ReadAll() *response.HTTPResponse {
	customer := cs.customerRepository.Read()

	if len(customer) == 0 {
		return &response.HTTPResponse{
			Message:   "data pelanggan kosong",
			Status:    200,
			IsSuccess: true,
		}
	}

	return &response.HTTPResponse{
		Message:   "data pelanggan berhasil di ambil",
		Status:    200,
		IsSuccess: true,
		Data:      customer,
	}
}

// Register implements customer.CustomerService
func (cs *customerService) Register(body *request_body.Register) *response.HTTPResponse {
	// retrieve customer by it's email
	customer := cs.customerRepository.ReadByEmail(body.Email)

	// check if customer already registered or not
	if customer.ID != "" || customer.Email != "" {
		return &response.HTTPResponse{
			Message: "data pelanggan sudah terdaftar",
			Status:  409,
		}
	}

	// secure password
	body.Password = cs.bcrypt.Hash(body.Password)

	// create customer
	new_customer, err := cs.customerRepository.Create(body)
	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}
	}

	// generate OTP
	otp := cs.codeGenerator.Generate()
	secureOTP := cs.bcrypt.Hash(otp)

	// generate token
	token := cs.uuid.GenerateID()

	// store verification
	if err := cs.verificationRepository.Upsert(token, new_customer.ID, secureOTP); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}
	}

	// send to email
	if err := cs.mailer.Send([]string{customer.Email}, "Verifikasi Email", verificationEmail(otp, body.Name)); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan pengiriman",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "registrasi berhasil",
		Status:    200,
		IsSuccess: true,
	}
}

// RequestVerification implements customer.CustomerService
func (cs *customerService) RequestVerification(body *verification_request_body.RequestVerification) *response.HTTPResponse {
	// retrieve customer by it's id
	customer := cs.customerRepository.ReadByID(body.CustomerID)

	// if customer not foun
	if customer.ID == "" || customer.Email == "" {
		return &response.HTTPResponse{
			Message: "pengguna tidak terdaftar",
			Status:  404,
		}
	}

	// generate OTP
	otp := cs.codeGenerator.Generate()
	secureOTP := cs.bcrypt.Hash(otp)

	// generate token
	token := cs.uuid.GenerateID()

	// save verification
	if err := cs.verificationRepository.Upsert(token, customer.ID, secureOTP); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan penyimpanan data",
			Status:  500,
		}
	}

	// send to email
	if err := cs.mailer.Send([]string{customer.Email}, "Verifikasi Email", verificationEmail(otp, customer.Name)); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan pengiriman",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "request verifikasi berhasil",
		Status:    200,
		IsSuccess: true,
	}
}

// ResetPassword implements customer.CustomerService
func (cs *customerService) ResetPassword(body *request_body.ResetPassword) *response.HTTPResponse {
	// retrieve forgot password token
	fp_data := cs.forgotPasswordRepository.ReadOne(body.Token)

	// if fp data not found
	if fp_data.ID == "" {
		return &response.HTTPResponse{
			Message: "sesi reset password tidak valid",
			Status:  440,
		}
	}

	// compare encryoted code with inputed code by user
	if !cs.bcrypt.Compare(body.Password, fp_data.Code) {
		return &response.HTTPResponse{
			Message: "kode reset password salah",
			Status:  401,
		}
	}

	// update password
	safePassword := cs.bcrypt.Hash(body.Password)
	if err := cs.customerRepository.UpdatePassword(safePassword, fp_data.UserID); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan update password",
			Status:  500,
		}
	}

	// delete reset password data
	if err := cs.forgotPasswordRepository.Delete(body.Token); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan penghapusan sesi reset password",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "reset password berhasil",
		Status:    200,
		IsSuccess: true,
	}
}

// UpdateProfile implements customer.CustomerService
func (cs *customerService) UpdateProfile(body *request_body.UpdateProfile) *response.HTTPResponse {
	if err := cs.customerRepository.UpdateProfile(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan update data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "profile berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

// VerifyVerification implements customer.CustomerService
func (cs *customerService) VerifyVerification(body *verification_request_body.Verify) *response.HTTPResponse {
	// retrieve verification data
	verif_data := cs.verificationRepository.ReadOne(body.Token, body.CustomerID)

	// check if verification data exists or not
	if verif_data.ID == "" || verif_data.Token == "" {
		return &response.HTTPResponse{
			Message: "sesi verifikasi tidak valid",
			Status:  440,
		}
	}

	// compare verification code
	if !cs.bcrypt.Compare(body.VerificationCode, verif_data.VerificationCode) {
		return &response.HTTPResponse{
			Message: "kode verifikasi salah",
			Status:  401,
		}
	}

	// update verification status
	if err := cs.customerRepository.VerifyVerification(verif_data.CustomerID); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat verifikasi",
			Status:  500,
		}
	}

	// delete verification session
	if err := cs.verificationRepository.Delete(verif_data.CustomerID); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat verifikasi tahap 2",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "akun berhasil diverifikasi",
		Status:    200,
		IsSuccess: true,
	}
}

// ViewProfile implements customer.CustomerService
func (cs *customerService) ViewProfile(body *request_body.ViewProfile) *response.HTTPResponse {
	customer := cs.customerRepository.ReadByID(body.Id)

	if customer.ID == "" {
		return &response.HTTPResponse{
			Message: "pengguna tidak ditemukan",
			Status:  404,
		}
	}

	customer.Password = ""

	return &response.HTTPResponse{
		Message:   "profile berhasil di ambil",
		Status:    200,
		IsSuccess: true,
		Data:      customer,
	}
}

// login implements customer.CustomerService
func (cs *customerService) Login(body *request_body.Login) *response.HTTPResponse {
	// retrieve customer by it's email
	customer := cs.customerRepository.ReadByEmail(body.Email)

	// check if customer already registered or not
	if customer.ID == "" || customer.Email == "" {
		return &response.HTTPResponse{
			Message: "pelanggan tidak terdaftar",
			Status:  409,
		}
	}

	// compare password
	if !cs.bcrypt.Compare(body.Password, customer.Password) {
		return &response.HTTPResponse{
			Message: "password salah",
			Status:  401,
		}
	}

	// payload
	payload := map[string]interface{}{
		"id":    customer.ID,
		"email": customer.Email,
	}

	token, err := cs.jwt.Sign(payload)
	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan pembuatan akses",
			Status:  500,
		}
	}

	customer.Password = ""

	return &response.HTTPResponse{
		Message:   "login berhasil",
		Status:    200,
		IsSuccess: true,
		Data:      customer,
		Token:     token,
	}
}

func CustomerService(infra *infrastructures.Infrastructures) customer.CustomerService {
	return &customerService{
		customerRepository:       infra.Repositories().Customer,
		forgotPasswordRepository: infra.Repositories().ForgotPassword,
		verificationRepository:   infra.Repositories().Verification,
		jwt:                      *infra.JsonWebToken(),
		bcrypt:                   *infra.Encryption(),
		mailer:                   *infra.Mailer(),
		uuid:                     *infra.Identifier(),
		codeGenerator:            *infra.CodeGenerator(),
	}
}
