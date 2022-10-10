package groceryimage

import request_body "kang-sayur-backend/model/web/request_body/grocery"

type GroceryImage struct {
	ID        string `json:"id" bson:"_id"`
	Source    string `json:"source" bson:"source"`
	GroceryID string `json:"grocery_id" bson:"grocery_id"`
}

type GroceryImageRepository interface {
	Create(data *request_body.CreateOrDeleteGroceryImage) error

	Delete(data *request_body.CreateOrDeleteGroceryImage) error

	ReadByGroceryID(grocery_id string) ([]GroceryImage, error)
}
