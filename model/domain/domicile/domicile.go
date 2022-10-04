package domicile

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/domicile"
)

type Domicile struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type DomicileService interface {
	Create(body *request_body.Create) *response.HTTPResponse

	Read() *response.HTTPResponse

	Update(body *request_body.UpdateOrDelete) *response.HTTPResponse

	Delete(body *request_body.UpdateOrDelete) *response.HTTPResponse
}

type DomicileRepository interface {
	Create(body *request_body.Create) (*Domicile, error)

	Read() ([]Domicile, error)

	Update(body *request_body.UpdateOrDelete) error

	Delete(body *request_body.UpdateOrDelete) error
}
