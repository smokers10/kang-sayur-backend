package address

import (
	request_body "kang-sayur-backend/model/web/request_body/address"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *request_body.CreateAddress) (*Address, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Address), args.Error(1)
}

func (mr *MockRepository) Read(data *request_body.DeleteOrReadAddress) ([]Address, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).([]Address), args.Error(1)
}

func (mr *MockRepository) ReadOne(data *request_body.ReadOne) (*Address, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Address), args.Error(1)
}

func (mr *MockRepository) Update(data *request_body.UpdateAddress) (*Address, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Address), args.Error(1)
}

func (mr *MockRepository) Delete(data *request_body.DeleteOrReadAddress) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
