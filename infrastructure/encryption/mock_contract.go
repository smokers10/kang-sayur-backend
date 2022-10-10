package encryption

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Hash(plaintext string) (hashed_string string) {
	args := m.Mock.Called(plaintext)
	return args.String(0)
}

func (m *MockContract) Compare(plaintext string, hashed_string string) (is_correct bool) {
	args := m.Mock.Called(plaintext, hashed_string)

	return args.Bool(0)
}
