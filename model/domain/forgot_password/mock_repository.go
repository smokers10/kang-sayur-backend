package forgotpassword

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Upsert(token string, user_id string, code string) error {
	args := mr.Mock.Called(token, user_id, code)
	return args.Error(0)
}

func (mr *MockRepository) ReadOne(token string) error {
	args := mr.Mock.Called(token)
	return args.Error(0)
}

func (mr *MockRepository) Delete(token string) error {
	args := mr.Mock.Called(token)
	return args.Error(0)
}
