package admin

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) CheckEmail(email string) *Admin {
	arg := mr.Mock.Called(email)

	return arg.Get(0).(*Admin)
}
