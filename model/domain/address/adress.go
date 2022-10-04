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
	Create(request_body.CreateAddress) *response.HTTPResponse

	Read(request_body.DeleteOrReadAddress) *response.HTTPResponse

	Update(request_body.UpdateAddress) *response.HTTPResponse

	Delelete(request_body.DeleteOrReadAddress) *response.HTTPResponse
}

type AddressRepository interface {
	Create(request_body.CreateAddress) (*Address, error)

	Read(request_body.DeleteOrReadAddress) ([]Address, error)

	ReadOne(request_body.ReadOne) (*Address, error)

	Update(request_body.UpdateAddress) (*Address, error)

	Delelete(request_body.DeleteOrReadAddress) error
}
