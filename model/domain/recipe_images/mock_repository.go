package recipeimages

import (
	request_body "kang-sayur-backend/model/web/request_body/recipe"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Read(recipe_id string) ([]RecipeImages, error) {
	args := mr.Mock.Called(recipe_id)
	return args.Get(0).([]RecipeImages), args.Error(0)
}

func (mr *MockRepository) Create(data *request_body.Image) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) DeleteImage(data *request_body.DeleteImage) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) UpdateOrder(data *request_body.UpdateOrderImage) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
