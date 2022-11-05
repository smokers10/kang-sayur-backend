package recipe

import (
	"kang-sayur-backend/model/domain/recipe"
	request_body "kang-sayur-backend/model/web/request_body/recipe"

	"go.mongodb.org/mongo-driver/mongo"
)

type recipeRepository struct {
	collection mongo.Collection
}

// Update implements recipe.RecipeRepository
func (*recipeRepository) Update(data *request_body.Update) error {
	panic("unimplemented")
}

// Create implements recipe.RecipeRepository
func (*recipeRepository) Create(data *request_body.Create) (*recipe.Recipe, error) {
	panic("unimplemented")
}

// Delete implements recipe.RecipeRepository
func (*recipeRepository) Delete(data *request_body.Delete) error {
	panic("unimplemented")
}

// ByName implements recipe.RecipeRepository
func (*recipeRepository) ByName(data *request_body.Searching) ([]recipe.Recipe, error) {
	panic("unimplemented")
}

// ReadAll implements recipe.RecipeRepository
func (*recipeRepository) ReadAll() ([]recipe.Recipe, error) {
	panic("unimplemented")
}

// ReadDetail implements recipe.RecipeRepository
func (*recipeRepository) ReadDetail(data *request_body.RecipeDetail) (*recipe.Recipe, error) {
	panic("unimplemented")
}

func RecipeRepository(db *mongo.Database) recipe.RecipeRepository {
	return &recipeRepository{*db.Collection("recipe")}
}
