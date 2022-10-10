package cart

import (
	request_body "kang-sayur-backend/model/web/request_body/cart"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(body *request_body.BasicAction) (*Cart, error) {
	args := mr.Mock.Called()
	return args.Get(0).(*Cart), args.Error(1)
}

func (mr *MockRepository) Delete(body *request_body.BasicAction) error {
	args := mr.Mock.Called()
	return args.Error(0)
}

func (mr *MockRepository) UpdateQuantity(body *request_body.UpdateQuantity) error {
	args := mr.Mock.Called()
	return args.Error(0)
}

func (mr *MockRepository) ReadCart(data *request_body.ReadCart) ([]Cart, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).([]Cart), args.Error(1)
}
