package recipe

import (
	response "kang-sayur-backend/model/web"
	request_body "kang-sayur-backend/model/web/request_body/recipe"
)

type Recipe struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Description  string `json:"description" bson:"description"`
	Price        int    `json:"price" bson:"price"`
	PointOfSales string `json:"point_of_sales" bson:"point_of_sales"`
	HowToCook    string `json:"how_to_cook" bson:"how_to_cook"`
	DeleteStatus string `json:"delete_status" bson:"delete_status"`
	CategoryID   string `json:"category_id" bson:"category_id"`
}

type RecipeService interface {
	Create(body *request_body.Create) *response.HTTPResponse

	ReadAll() *response.HTTPResponse

	ReadRecipeDetail(body *request_body.RecipeDetail) *response.HTTPResponse

	Searching(body *request_body.Searching) *response.HTTPResponse

	Update(body *request_body.Update) *response.HTTPResponse

	Delete(body *request_body.Delete) *response.HTTPResponse

	UploadImage(body *request_body.Image) *response.HTTPResponse

	DeleteImage(body *request_body.DeleteImage) *response.HTTPResponse

	UpdateImageOrder(body *request_body.UpdateOrderImage) *response.HTTPResponse

	AddDetail(body *request_body.AddDetail) *response.HTTPResponse

	UpdateDetail(body *request_body.UpdateOrDeleteDetail) *response.HTTPResponse

	DeleteDetail(body *request_body.UpdateOrDeleteDetail) *response.HTTPResponse
}

type RecipeRepository interface {
	Create(data *request_body.Create) (*Recipe, error)

	ReadAll() ([]Recipe, error)

	ReadDetail(data *request_body.RecipeDetail) (*Recipe, error)

	ByName(data *request_body.Searching) ([]Recipe, error)

	Delete(data *request_body.Delete) error

	Update(data *request_body.Update) error
}
