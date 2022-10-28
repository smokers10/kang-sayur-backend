package grocery

import (
	"kang-sayur-backend/model/domain/grocery"
	request_body "kang-sayur-backend/model/web/request_body/grocery"

	"go.mongodb.org/mongo-driver/mongo"
)

type groceryRepository struct {
	collection mongo.Collection
}

// Create implements grocery.GroceryRepository
func (groceryRepository) Create(data *request_body.Create) error {
	panic("unimplemented")
}

// Delete implements grocery.GroceryRepository
func (groceryRepository) Delete(data *request_body.UpdateOrDelete) error {
	panic("unimplemented")
}

// Read implements grocery.GroceryRepository
func (groceryRepository) Read(data *request_body.Create) ([]grocery.Grocery, error) {
	panic("unimplemented")
}

// ReadBestSeller implements grocery.GroceryRepository
func (groceryRepository) ReadBestSeller() ([]grocery.Grocery, error) {
	panic("unimplemented")
}

// ReadByCategory implements grocery.GroceryRepository
func (groceryRepository) ReadByCategory(data *request_body.ByCategory) ([]grocery.Grocery, error) {
	panic("unimplemented")
}

// ReadByKeyword implements grocery.GroceryRepository
func (groceryRepository) ReadByKeyword(data *request_body.ByKeyword) ([]grocery.Grocery, error) {
	panic("unimplemented")
}

// ReadCheapest implements grocery.GroceryRepository
func (groceryRepository) ReadCheapest() ([]grocery.Grocery, error) {
	panic("unimplemented")
}

// Update implements grocery.GroceryRepository
func (groceryRepository) Update(data *request_body.UpdateOrDelete) error {
	panic("unimplemented")
}

func GroceryRepository(db *mongo.Database) grocery.GroceryRepository {
	return groceryRepository{collection: *db.Collection("grocery")}
}
