package category

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/category"
)

type Category struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Ilustration string `json:"Ilustration" bson:"Ilustration"`
}

type CategoryService interface {
	Create(body request_body.CreateCategory) response.HTTPResponse

	Read() response.HTTPResponse

	Update(body request_body.UpdateOrDeleteCategory) response.HTTPResponse

	Delete(body request_body.UpdateOrDeleteCategory) response.HTTPResponse
}

type CategoryRepository interface {
	Create(body request_body.CreateCategory) (*Category, error)

	Read() ([]Category, error)

	Update(body request_body.UpdateOrDeleteCategory) error

	Delete(body request_body.UpdateOrDeleteCategory) error
}
