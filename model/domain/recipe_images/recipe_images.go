package recipeimages

import request_body "kang-sayur-backend/model/web/request_body/recipe"

type RecipeImages struct {
	ID       string `json:"id" bson:"_id"`
	Source   string `json:"source" bson:"source"`
	Order    int    `json:"order" bson:"order"`
	RecipeID string `json:"recipe_id" bson:"recipe_id"`
}

type RecipeImagesRepository interface {
	Read(recipe_id string) ([]RecipeImages, error)

	Create(data *request_body.Image) error

	DeleteImage(data *request_body.DeleteImage) error

	UpdateOrder(data *request_body.UpdateOrderImage) error
}
