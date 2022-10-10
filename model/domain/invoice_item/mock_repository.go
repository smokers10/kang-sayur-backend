package domain

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	Mock mock.Mock
}

func (mr *MockRepository) BulkCreate(data []InvoiceItem) error {
	args := mr.Mock.Called(data)
	return args.Error(0)
}
