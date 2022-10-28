package groceryimage

import (
	groceryimage "kang-sayur-backend/model/domain/grocery_image"
	request_body "kang-sayur-backend/model/web/request_body/grocery"

	"go.mongodb.org/mongo-driver/mongo"
)

type groceryImageRepository struct {
	collection mongo.Collection
}

// Create implements groceryimage.GroceryImageRepository
func (*groceryImageRepository) Create(data *request_body.CreateOrDeleteGroceryImage) error {
	panic("unimplemented")
}

// Delete implements groceryimage.GroceryImageRepository
func (*groceryImageRepository) Delete(data *request_body.CreateOrDeleteGroceryImage) error {
	panic("unimplemented")
}

// ReadByGroceryID implements groceryimage.GroceryImageRepository
func (*groceryImageRepository) ReadByGroceryID(grocery_id string) ([]groceryimage.GroceryImage, error) {
	panic("unimplemented")
}

func GroceryImageRepository(db *mongo.Database) groceryimage.GroceryImageRepository {
	return &groceryImageRepository{collection: *db.Collection("grocery_image")}
}
