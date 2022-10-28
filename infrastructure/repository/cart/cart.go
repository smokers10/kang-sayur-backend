package cart

import (
	"kang-sayur-backend/model/domain/cart"
	request_body "kang-sayur-backend/model/web/request_body/cart"

	"go.mongodb.org/mongo-driver/mongo"
)

type cartRepository struct {
	collection mongo.Collection
}

// Create implements cart.CartRepository
func (*cartRepository) Create(data *request_body.BasicAction) (*cart.Cart, error) {
	panic("unimplemented")
}

// Delete implements cart.CartRepository
func (*cartRepository) Delete(data *request_body.BasicAction) error {
	panic("unimplemented")
}

// ReadCart implements cart.CartRepository
func (*cartRepository) ReadCart(data *request_body.ReadCart) ([]cart.Cart, error) {
	panic("unimplemented")
}

// UpdateQuantity implements cart.CartRepository
func (*cartRepository) UpdateQuantity(data *request_body.UpdateQuantity) error {
	panic("unimplemented")
}

func CartRepository(db *mongo.Database) cart.CartRepository {
	return &cartRepository{collection: *db.Collection("cart")}
}
