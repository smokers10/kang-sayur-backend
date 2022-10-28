package recipeimage

import (
	recipeimages "kang-sayur-backend/model/domain/recipe_images"
	request_body "kang-sayur-backend/model/web/request_body/recipe"

	"go.mongodb.org/mongo-driver/mongo"
)

type recipeImageRepository struct {
	collection mongo.Collection
}

// Create implements recipeimages.RecipeImagesRepository
func (*recipeImageRepository) Create(data *request_body.Image) error {
	panic("unimplemented")
}

// DeleteImage implements recipeimages.RecipeImagesRepository
func (*recipeImageRepository) DeleteImage(data *request_body.DeleteImage) error {
	panic("unimplemented")
}

// Read implements recipeimages.RecipeImagesRepository
func (*recipeImageRepository) Read(recipe_id string) ([]recipeimages.RecipeImages, error) {
	panic("unimplemented")
}

// UpdateOrder implements recipeimages.RecipeImagesRepository
func (*recipeImageRepository) UpdateOrder(data *request_body.UpdateOrderImage) error {
	panic("unimplemented")
}

func RecipeImageRepository(db *mongo.Database) recipeimages.RecipeImagesRepository {
	return &recipeImageRepository{collection: *db.Collection("recipe_image")}
}
