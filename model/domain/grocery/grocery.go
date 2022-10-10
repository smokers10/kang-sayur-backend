package grocery

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/grocery"
)

type Grocery struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Description  string `json:"description" bson:"description"`
	UnitType     string `json:"unit_type" bson:"unit_type"`
	PointOfSales string `json:"point_of_sales" bson:"point_of_sales"`
	CategoryID   string `json:"category_id" bson:"category_id"`
	PriceID      string `json:"price_id" bson:"price_id"`
}

type GroceryService interface {
	Create(body *request_body.Create) *response.HTTPResponse

	Read(body *request_body.Create) *response.HTTPResponse

	ReadByCategory(body *request_body.ByCategory) *response.HTTPResponse

	ReadByKeyword(body *request_body.ByKeyword) *response.HTTPResponse

	ReadCheapest() *response.HTTPResponse

	ReadBestSeller() *response.HTTPResponse

	Update(body *request_body.UpdateOrDelete) *response.HTTPResponse

	Delete(body *request_body.UpdateOrDelete) *response.HTTPResponse

	SetPrice(body *request_body.SetOrUpdatePrice) *response.HTTPResponse

	UpdatePrice(body *request_body.SetOrUpdatePrice) *response.HTTPResponse
}

type GroceryRepository interface {
	Create(data *request_body.Create) error

	Read(data *request_body.Create) ([]Grocery, error)

	ReadByCategory(data *request_body.ByCategory) ([]Grocery, error)

	ReadByKeyword(data *request_body.ByKeyword) ([]Grocery, error)

	ReadCheapest() ([]Grocery, error)

	ReadBestSeller() ([]Grocery, error)

	Update(data *request_body.UpdateOrDelete) error

	Delete(data *request_body.UpdateOrDelete) error
}
