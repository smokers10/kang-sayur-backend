package domain

import (
	request_body "kang-sayur-backend/model/web/request_body/recipe"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Read(recipe_id string) ([]RecipeDetail, error) {
	args := mr.Mock.Called(recipe_id)
	return args.Get(0).([]RecipeDetail), args.Error(1)
}

func (mr *MockRepository) Create(data *request_body.AddDetail) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Update(data *request_body.UpdateOrDeleteDetail) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) DeleteDetail(data *request_body.UpdateOrDeleteDetail) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
