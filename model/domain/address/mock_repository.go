package address

import (
	request_body "kang-sayur-backend/model/web/request_body/address"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(body *request_body.CreateAddress) (*Address, error) {
	args := mr.Mock.Called(body)

	return args.Get(0).(*Address), args.Error(1)
}

func (mr *MockRepository) Read(body *request_body.DeleteOrReadAddress) ([]Address, error) {
	args := mr.Mock.Called(body)

	return args.Get(0).([]Address), args.Error(1)
}

func (mr *MockRepository) ReadOne(body *request_body.ReadOne) (*Address, error) {
	args := mr.Mock.Called(body)

	return args.Get(0).(*Address), args.Error(1)
}

func (mr *MockRepository) Update(body *request_body.UpdateAddress) (*Address, error) {
	args := mr.Mock.Called(body)

	return args.Get(0).(*Address), args.Error(1)
}

func (mr *MockRepository) Delelete(body *request_body.DeleteOrReadAddress) error {
	args := mr.Mock.Called(body)

	return args.Error(0)
}
