package groceryimage

import (
	request_body "kang-sayur-backend/model/web/request_body/grocery"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *request_body.CreateOrDeleteGroceryImage) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Delete(data *request_body.CreateOrDeleteGroceryImage) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) ReadByGroceryID(grocery_id string) ([]GroceryImage, error) {
	args := mr.Mock.Called(grocery_id)
	return args.Get(0).([]GroceryImage), args.Error(1)
}
