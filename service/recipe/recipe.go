package service

import (
	"kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/recipe"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/recipe"
)

type recipeService struct {
	repository recipe.RecipeRepository
}

// AddDetail implements recipe.RecipeService
func (*recipeService) AddDetail(body *request_body.AddDetail) *response.HTTPResponse {
	panic("unimplemented")
}

// Create implements recipe.RecipeService
func (*recipeService) Create(body *request_body.Create) *response.HTTPResponse {
	panic("unimplemented")
}

// Delete implements recipe.RecipeService
func (*recipeService) Delete(body *request_body.Delete) *response.HTTPResponse {
	panic("unimplemented")
}

// DeleteDetail implements recipe.RecipeService
func (*recipeService) DeleteDetail(body *request_body.UpdateOrDeleteDetail) *response.HTTPResponse {
	panic("unimplemented")
}

// DeleteImage implements recipe.RecipeService
func (*recipeService) DeleteImage(body *request_body.DeleteImage) *response.HTTPResponse {
	panic("unimplemented")
}

// ReadAll implements recipe.RecipeService
func (*recipeService) ReadAll() *response.HTTPResponse {
	panic("unimplemented")
}

// ReadRecipeDetail implements recipe.RecipeService
func (*recipeService) ReadRecipeDetail(body *request_body.RecipeDetail) *response.HTTPResponse {
	panic("unimplemented")
}

// Searching implements recipe.RecipeService
func (*recipeService) Searching(body *request_body.Searching) *response.HTTPResponse {
	panic("unimplemented")
}

// Update implements recipe.RecipeService
func (*recipeService) Update(body *request_body.Update) *response.HTTPResponse {
	panic("unimplemented")
}

// UpdateDetail implements recipe.RecipeService
func (*recipeService) UpdateDetail(body *request_body.UpdateOrDeleteDetail) *response.HTTPResponse {
	panic("unimplemented")
}

// UpdateImageOrder implements recipe.RecipeService
func (*recipeService) UpdateImageOrder(body *request_body.UpdateOrderImage) *response.HTTPResponse {
	panic("unimplemented")
}

// UploadImage implements recipe.RecipeService
func (*recipeService) UploadImage(body *request_body.Image) *response.HTTPResponse {
	panic("unimplemented")
}

func RecipeService(infra injector.Infrastructures) recipe.RecipeService {
	return &recipeService{repository: infra.Repositories().Recipe}
}
