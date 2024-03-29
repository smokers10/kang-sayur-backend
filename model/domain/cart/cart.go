package cart

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/cart"
)

type Cart struct {
	ID         string `json:"id" bson:"_id"`
	CustomerID string `json:"customer_id" bson:"customer_id"`
	ProductID  string `json:"product_id" bson:"product_id"`
	Quantity   int    `json:"quantity" bson:"quantity"`
}

type CartService interface {
	Add(body *request_body.BasicAction) *response.HTTPResponse

	Delete(body *request_body.BasicAction) *response.HTTPResponse

	UpdateQuantity(body *request_body.UpdateQuantity) *response.HTTPResponse

	GetCart(body *request_body.ReadCart) *response.HTTPResponse
}

type CartRepository interface {
	Create(data *request_body.BasicAction) (*Cart, error)

	Delete(data *request_body.BasicAction) error

	UpdateQuantity(data *request_body.UpdateQuantity) error

	ReadCart(data *request_body.ReadCart) ([]Cart, error)
}
