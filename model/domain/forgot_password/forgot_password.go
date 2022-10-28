package forgotpassword

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/forgot_password"
)

type ForgotPassword struct {
	ID     string `json:"id" bson:"_id"`
	Token  string `json:"token" bson:"token"`
	UserID string `json:"user_id" bson:"user_id"`
	Code   string `json:"code" bson:"code"`
}

type ForgotPasswordService interface {
	Request(body *request_body.Request) *response.HTTPResponse

	Reset(body *request_body.ResetPassword) *response.HTTPResponse
}

type ForgotPasswordRepository interface {
	Upsert(token string, user_id string, code string) error

	ReadOne(token string) *ForgotPassword

	Delete(token string) error
}
