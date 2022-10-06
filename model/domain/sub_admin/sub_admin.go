package domain

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/sub_admin"
)

type SubAdmin struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Status   string `json:"status" bson:"status"`
	Position string `json:"position" bson:"position"`
}

type SubAdminService interface {
	Create(body *request_body.Create) response.HTTPResponse

	Read() response.HTTPResponse

	Update(body *request_body.Update) response.HTTPResponse

	UpdateStatus(body *request_body.UpdateStatus) response.HTTPResponse

	Delete(body *request_body.Delete) response.HTTPResponse

	SetPermission(body *request_body.SetPermission) response.HTTPResponse
}

type SubAdminRepository interface {
	Create(body *request_body.Create) error

	Read() ([]SubAdmin, error)

	Update(body *request_body.Update)

	UpdateStatus(body *request_body.SetPermission)

	Delete(body *request_body.Delete)
}
