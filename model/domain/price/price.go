package domain

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/price"
)

type Price struct {
	ID         string `json:"id" bson:"_id"`
	Price      int    `json:"price" bson:"price"`
	DomicileID string `json:"domicile_id" bson:"domicile_id"`
}

type PriceService interface {
	Create(body *request_body.Data) response.HTTPResponse

	Read() response.HTTPResponse

	ReadByDomicile(body *request_body.ByDomicile) response.HTTPResponse
}

type PriceRepository interface {
	Upsert(body *request_body.Data) error

	Read() ([]Price, error)

	ReadByDomicile(body *request_body.ByDomicile) ([]Price, error)
}
