package domain

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Upsert(data *Permission) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}

func (mr *MockRepository) ReadOne(sub_admin_id string) (*Permission, error) {
	args := mr.Mock.Called(sub_admin_id)
	return args.Get(0).(*Permission), args.Error(1)
}
