package service

import (
	"kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/invoice"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/invoice"
)

type invoiceService struct {
	repository invoice.InvoiceRepository
}

// Cancel implements invoice.InvoiceService
func (*invoiceService) Cancel(body *request_body.UpdateStatus) *response.HTTPResponse {
	panic("unimplemented")
}

// Checkout implements invoice.InvoiceService
func (*invoiceService) Checkout(body *request_body.Checkout) *response.HTTPResponse {
	panic("unimplemented")
}

// Pay implements invoice.InvoiceService
func (*invoiceService) Pay(body *request_body.Pay) *response.HTTPResponse {
	panic("unimplemented")
}

// ReadAll implements invoice.InvoiceService
func (*invoiceService) ReadAll() *response.HTTPResponse {
	panic("unimplemented")
}

// ReadOne implements invoice.InvoiceService
func (*invoiceService) ReadOne(body *request_body.ReadOne) *response.HTTPResponse {
	panic("unimplemented")
}

func InvoiceService(infra injector.Infrastructures) invoice.InvoiceService {
	return &invoiceService{repository: infra.Repositories().Invoice}
}
