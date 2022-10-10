package domain

import (
	request_body "kang-sayur-backend/model/web/request_body/recipe"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *request_body.Create) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) ReadAll() ([]Recipe, error) {
	args := mr.Mock.Called()
	return args.Get(0).([]Recipe), args.Error(1)
}

func (mr *MockRepository) ReadDetail(data *request_body.RecipeDetail) (*Recipe, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Recipe), args.Error(1)
}

func (mr *MockRepository) ByName(data *request_body.Searching) ([]Recipe, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).([]Recipe), args.Error(1)
}
