package codegenerator

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Generate() string {
	args := m.Mock.Called(0)

	return args.String(0)
}
