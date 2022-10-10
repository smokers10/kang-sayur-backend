package verification

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) Upsert(token string, customer_id string, verification_code string) error {
	args := mr.Mock.Called(token, customer_id, verification_code)
	return args.Error(0)
}

func (mr *MockRepository) ReadOne(token string, customer_id string) (*Verification, error) {
	args := mr.Mock.Called(token, customer_id)
	return args.Get(0).(*Verification), args.Error(0)

}

func (mr *MockRepository) Delete(customer_id string) error {
	args := mr.Mock.Called(customer_id)
	return args.Error(0)
}
