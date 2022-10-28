package subadmin

import (
	request_body "kang-sayur-backend/model/web/request_body/sub_admin"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *request_body.Create) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Read() ([]SubAdmin, error) {
	args := mr.Mock.Called()
	return args.Get(0).([]SubAdmin), args.Error(1)
}

func (mr *MockRepository) Update(data *request_body.Update) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) UpdateStatus(data *request_body.SetPermission) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Delete(data *request_body.Delete) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
