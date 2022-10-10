package domicile

import (
	request_body "kang-sayur-backend/model/web/request_body/domicile"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Create(data *request_body.Create) (*Domicile, error) {
	args := mr.Mock.Called(data)
	return args.Get(0).(*Domicile), args.Error(1)

}

func (mr *MockRepository) Read() ([]Domicile, error) {
	args := mr.Mock.Called()
	return args.Get(0).([]Domicile), args.Error(1)
}

func (mr *MockRepository) Update(data *request_body.UpdateOrDelete) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) Delete(data *request_body.UpdateOrDelete) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
