package customer

import (
	request_body "kang-sayur-backend/model/web/request_body/customer"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *request_body.Register) (*Customer, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Customer), args.Error(1)
}

func (mr *MockRepository) ReadByEmail(email string) *Customer {
	args := mr.Mock.Called(email)
	return args.Get(0).(*Customer)
}

func (mr *MockRepository) UpdateProfile(data *request_body.UpdateProfile) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Read() []Customer {
	args := mr.Mock.Called()
	return args.Get(0).([]Customer)
}

func (mr *MockRepository) ReadByID(id string) *Customer {
	args := mr.Mock.Called(id)
	return args.Get(0).(*Customer)
}

func (mr *MockRepository) UpdatePassword(password string, customer_id string) error {
	args := mr.Mock.Called(password, customer_id)
	return args.Error(0)
}

func (mr *MockRepository) VerifyVerification(customer_id string) error {
	args := mr.Mock.Called(customer_id)
	return args.Error(0)
}
