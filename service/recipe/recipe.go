package service

import (
	"fmt"
	"kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/model/domain/recipe"
	recipedetail "kang-sayur-backend/model/domain/recipe_detail"
	recipeimages "kang-sayur-backend/model/domain/recipe_images"
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/recipe"
)

type recipeService struct {
	repository             recipe.RecipeRepository
	recipeDetailRepository recipedetail.RecipeDetailRepository
	recipeImageRepository  recipeimages.RecipeImagesRepository
}

// AddDetail implements recipe.RecipeService
func (rs *recipeService) AddDetail(body *request_body.AddDetail) *response.HTTPResponse {
	inserted, err := rs.recipeDetailRepository.Create(body)

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat menyimpan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "detail resep berhasil disimpan",
		Status:    200,
		IsSuccess: true,
		Data:      inserted,
	}
}

// Create implements recipe.RecipeService
func (rs *recipeService) Create(body *request_body.Create) *response.HTTPResponse {
	inserted, err := rs.repository.Create(body)

	if err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat menyimpan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "resep berhasil disimpan",
		Status:    200,
		IsSuccess: true,
		Data:      inserted,
	}
}

// Delete implements recipe.RecipeService
func (rs *recipeService) Delete(body *request_body.Delete) *response.HTTPResponse {
	if err := rs.repository.Delete(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat pengapusan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "resep berhasil dihapus",
		Status:    200,
		IsSuccess: true,
	}
}

// DeleteDetail implements recipe.RecipeService
func (rs *recipeService) DeleteDetail(body *request_body.UpdateOrDeleteDetail) *response.HTTPResponse {
	if err := rs.recipeDetailRepository.DeleteDetail(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat pengapusan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "detail resep berhasil dihapus",
		Status:    200,
		IsSuccess: true,
	}
}

// DeleteImage implements recipe.RecipeService
func (rs *recipeService) DeleteImage(body *request_body.DeleteImage) *response.HTTPResponse {
	panic("unimplemented")
}

// ReadAll implements recipe.RecipeService
func (rs *recipeService) ReadAll() *response.HTTPResponse {
	recipes, err := rs.repository.ReadAll()

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "resep berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      recipes,
	}
}

// ReadRecipeDetail implements recipe.RecipeService
func (rs *recipeService) ReadRecipeDetail(body *request_body.RecipeDetail) *response.HTTPResponse {
	recipes, err := rs.recipeDetailRepository.Read(body.ID)

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "detail resep berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      recipes,
	}
}

// Searching implements recipe.RecipeService
func (rs *recipeService) Searching(body *request_body.Searching) *response.HTTPResponse {
	recipes, err := rs.repository.ByName(body)

	if err != nil {
		fmt.Println(err)
		return &response.HTTPResponse{
			Message: "kesalahan saat pengambilan data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "resep berhasil diambil",
		Status:    200,
		IsSuccess: true,
		Data:      recipes,
	}
}

// Update implements recipe.RecipeService
func (rs *recipeService) Update(body *request_body.Update) *response.HTTPResponse {
	if err := rs.repository.Update(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "resep berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

// UpdateDetail implements recipe.RecipeService
func (rs *recipeService) UpdateDetail(body *request_body.UpdateOrDeleteDetail) *response.HTTPResponse {
	if err := rs.recipeDetailRepository.Update(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "detail resep berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

// UpdateImageOrder implements recipe.RecipeService
func (rs *recipeService) UpdateImageOrder(body *request_body.UpdateOrderImage) *response.HTTPResponse {
	if err := rs.recipeImageRepository.UpdateOrder(body); err != nil {
		return &response.HTTPResponse{
			Message: "kesalahan saat update data",
			Status:  500,
		}
	}

	return &response.HTTPResponse{
		Message:   "urutan gambar resep berhasil diupdate",
		Status:    200,
		IsSuccess: true,
	}
}

// UploadImage implements recipe.RecipeService
func (rs *recipeService) UploadImage(body *request_body.Image) *response.HTTPResponse {
	panic("unimplemented")
}

func RecipeService(infra injector.Infrastructures) recipe.RecipeService {
	return &recipeService{repository: infra.Repositories().Recipe}
}
