package address

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/address"
)

type Address struct {
	ID         string `json:"id" bson:"_id"`
	Name       string `json:"name" bson:"name"`
	Address    string `json:"address" bson:"address"`
	Status     string `json:"status" bson:"status"`
	CustomerID string `json:"customer_id" bson:"customer_id"`
}

type AddressService interface {
	Create(body *request_body.CreateAddress) *response.HTTPResponse

	Read(body *request_body.DeleteOrReadAddress) *response.HTTPResponse

	Update(body *request_body.UpdateAddress) *response.HTTPResponse

	Delete(body *request_body.DeleteOrReadAddress) *response.HTTPResponse
}

type AddressRepository interface {
	Create(data *request_body.CreateAddress) (*Address, error)

	Read(data *request_body.DeleteOrReadAddress) ([]Address, error)

	ReadOne(data *request_body.ReadOne) (*Address, error)

	Update(data *request_body.UpdateAddress) (*Address, error)

	Delete(data *request_body.DeleteOrReadAddress) error
}
