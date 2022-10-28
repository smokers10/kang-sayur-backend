package invoice

import (
	"kang-sayur-backend/model/domain/invoice"
	request_body "kang-sayur-backend/model/web/request_body/invoice"

	"go.mongodb.org/mongo-driver/mongo"
)

type invoiceRepository struct {
	collection mongo.Collection
}

// Create implements invoice.InvoiceRepository
func (*invoiceRepository) Create(data *invoice.Invoice) error {
	panic("unimplemented")
}

// ReadAll implements invoice.InvoiceRepository
func (*invoiceRepository) ReadAll() ([]invoice.Invoice, error) {
	panic("unimplemented")
}

// ReadOne implements invoice.InvoiceRepository
func (*invoiceRepository) ReadOne(data *request_body.ReadOne) (*invoice.Invoice, error) {
	panic("unimplemented")
}

// UpdateStatus implements invoice.InvoiceRepository
func (*invoiceRepository) UpdateStatus(data *request_body.UpdateStatus) error {
	panic("unimplemented")
}

func InvoiceRepository(db *mongo.Database) invoice.InvoiceRepository {
	return &invoiceRepository{collection: *db.Collection("invoice")}
}
