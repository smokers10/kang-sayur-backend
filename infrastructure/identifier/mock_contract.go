package identifier

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) GenerateID() (ID string) {
	args := m.Mock.Called()

	return args.String(0)
}
