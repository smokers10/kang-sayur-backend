package domain

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/invoice"
)

type Invoice struct {
	ID            string
	InvoiceCode   string
	PaymentStatus string
	PaymentToken  string
	PaymentMethod string
	TotalPayment  string
	CustomerID    string
}

type InvoiceService interface {
	Checkout(body *request_body.Checkout) *response.HTTPResponse

	Pay(body *request_body.Pay) *response.HTTPResponse

	Cancel(body *request_body.UpdateStatus) *response.HTTPResponse

	ReadAll() *response.HTTPResponse

	ReadOne(body *request_body.ReadOne) *response.HTTPResponse
}

type InvoiceRepository interface {
	Create(data *Invoice) error

	UpdateStatus(request_body.UpdateStatus) error

	ReadAll() ([]Invoice, error)

	ReadOne(body *request_body.ReadOne) (*Invoice, error)
}
