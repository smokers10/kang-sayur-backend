package domain

import "kang-sayur-backend/model/web/request_body/invoice"

type InvoiceItem struct {
	ID          string `json:"id" bson:"_id"`
	ProductID   string `json:"product_id" bson:"product_id"`
	ProductType string `json:"product_type" bson:"product_type"`
	Quantities  string `json:"quantities" bson:"quantities"`
	SubTotal    string `json:"sub_total" bson:"sub_total"`
	InvoiceID   string `json:"invoice_id" bson:"invoice_id"`
}

type InvoiceRepository interface {
	BulkCreate(data []invoice.Item)
}
