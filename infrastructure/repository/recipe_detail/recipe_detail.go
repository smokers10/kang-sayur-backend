package recipedetail

import (
	recipedetail "kang-sayur-backend/model/domain/recipe_detail"
	request_body "kang-sayur-backend/model/web/request_body/recipe"

	"go.mongodb.org/mongo-driver/mongo"
)

type recipeDetailRepository struct {
	collection mongo.Collection
}

// Create implements recipedetail.RecipeDetailRepository
func (*recipeDetailRepository) Create(data *request_body.AddDetail) (*recipedetail.RecipeDetail, error) {
	panic("unimplemented")
}

// DeleteDetail implements recipedetail.RecipeDetailRepository
func (*recipeDetailRepository) DeleteDetail(data *request_body.UpdateOrDeleteDetail) error {
	panic("unimplemented")
}

// Read implements recipedetail.RecipeDetailRepository
func (*recipeDetailRepository) Read(recipe_id string) ([]recipedetail.RecipeDetail, error) {
	panic("unimplemented")
}

// Update implements recipedetail.RecipeDetailRepository
func (*recipeDetailRepository) Update(data *request_body.UpdateOrDeleteDetail) error {
	panic("unimplemented")
}

func RecipeDetailRepository(db *mongo.Database) recipedetail.RecipeDetailRepository {
	return &recipeDetailRepository{collection: *db.Collection("recipe_detail")}
}
