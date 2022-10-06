package domain

import (
	request_body "kang-sayur-backend/model/web/request_body/grocery"
)

type Price struct {
	ID         string `json:"id" bson:"_id"`
	Price      int    `json:"price" bson:"price"`
	DomicileID string `json:"domicile_id" bson:"domicile_id"`
}

type PriceRepository interface {
	Upsert(body *request_body.SetOrUpdatePrice) error

	ByGroceryID(product_id string) (*Price, error)
}
