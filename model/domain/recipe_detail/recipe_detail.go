package recipedetail

import request_body "kang-sayur-backend/model/web/request_body/recipe"

type RecipeDetail struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	UnitType string `json:"unit_type" bson:"unit_type"`
	RecipeID string `json:"recipe_id" bson:"recipe_id"`
}

type RecipeDetailRepository interface {
	Read(recipe_id string) ([]RecipeDetail, error)

	Create(data *request_body.AddDetail) error

	Update(data *request_body.UpdateOrDeleteDetail) error

	DeleteDetail(data *request_body.UpdateOrDeleteDetail) error
}
