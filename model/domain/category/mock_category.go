package category

import (
	request_body "kang-sayur-backend/model/web/request_body/category"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(body request_body.CreateCategory) (*Category, error) {
	args := mr.Mock.Called(body)
	return args.Get(0).(*Category), args.Error(1)
}

func (mr *MockRepository) Read() ([]Category, error) {
	args := mr.Mock.Called()
	return args.Get(0).([]Category), args.Error(1)
}

func (mr *MockRepository) Update(body request_body.UpdateOrDeleteCategory) error {
	args := mr.Mock.Called(body)
	return args.Error(0)
}

func (mr *MockRepository) Delete(body request_body.UpdateOrDeleteCategory) error {
	args := mr.Mock.Called(body)
	return args.Error(0)
}
