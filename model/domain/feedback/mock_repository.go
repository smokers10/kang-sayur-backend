package feedback

import (
	request_body "kang-sayur-backend/model/web/request_body/feedback"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *request_body.Create) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Read(data *request_body.ReadOnProduct) ([]Feedback, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).([]Feedback), args.Error(1)
}

func (mr *MockRepository) ReadOne(data *request_body.ReadOne) (*Feedback, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Feedback), args.Error(1)
}

func (mr *MockRepository) Update(data *request_body.UpdateOrDelete) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Delete(data *request_body.UpdateOrDelete) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
