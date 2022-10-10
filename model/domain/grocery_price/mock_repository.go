package domain

import (
	request_body "kang-sayur-backend/model/web/request_body/grocery"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Upsert(body *request_body.SetOrUpdatePrice) error {
	args := mr.Mock.Called(body)
	return args.Error(0)
}

func (mr *MockRepository) ByGroceryID(product_id string) (*Price, error) {
	args := mr.Mock.Called(product_id)
	return args.Get(0).(*Price), args.Error(1)
}
