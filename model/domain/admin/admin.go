package admin

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/admin"
)

type Admin struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type AdminService interface {
	RequestLogin(body request_body.LoginRequest) response.HTTPResponse

	Login(body request_body.Login) response.HTTPResponse

	Logout() response.HTTPResponse
}

type AdminRepository interface {
	CheckEmail(email string) *Admin
}
