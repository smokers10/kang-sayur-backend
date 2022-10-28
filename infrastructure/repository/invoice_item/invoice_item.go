package invoiceitem

import (
	invoiceitem "kang-sayur-backend/model/domain/invoice_item"
	"kang-sayur-backend/model/web/request_body/invoice"

	"go.mongodb.org/mongo-driver/mongo"
)

type invoiceItemRepository struct {
	collection mongo.Collection
}

// BulkCreate implements invoiceitem.InvoiceItemRepository
func (*invoiceItemRepository) BulkCreate(data []invoice.Item) {
	panic("unimplemented")
}

func InvoiceItemRepository(db *mongo.Database) invoiceitem.InvoiceItemRepository {
	return &invoiceItemRepository{collection: *db.Collection("invoice_item")}
}
