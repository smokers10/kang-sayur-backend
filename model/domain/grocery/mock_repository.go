package grocery

import (
	request_body "kang-sayur-backend/model/web/request_body/grocery"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *request_body.Create) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Read(data *request_body.Create) ([]Grocery, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).([]Grocery), args.Error(1)
}

func (mr *MockRepository) ReadByCategory(data *request_body.ByCategory) ([]Grocery, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).([]Grocery), args.Error(1)
}

func (mr *MockRepository) ReadByKeyword(data *request_body.ByKeyword) ([]Grocery, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).([]Grocery), args.Error(1)
}

func (mr *MockRepository) ReadCheapest() ([]Grocery, error) {
	args := mr.Mock.Called()
	return args.Get(0).([]Grocery), args.Error(1)
}

func (mr *MockRepository) ReadBestSeller() ([]Grocery, error) {
	args := mr.Mock.Called()
	return args.Get(0).([]Grocery), args.Error(1)
}

func (mr *MockRepository) Update(data *request_body.UpdateOrDelete) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Delete(data *request_body.UpdateOrDelete) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
