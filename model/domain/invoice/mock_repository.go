package invoice

import (
	request_body "kang-sayur-backend/model/web/request_body/invoice"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *Invoice) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) UpdateStatus(data *request_body.UpdateStatus) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) ReadAll() ([]Invoice, error) {
	args := mr.Mock.Called()
	return args.Get(0).([]Invoice), args.Error(1)
}

func (mr *MockRepository) ReadOne(data *request_body.ReadOne) (*Invoice, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Invoice), args.Error(1)
}
