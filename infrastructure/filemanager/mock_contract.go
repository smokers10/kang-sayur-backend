package filemanager

import "github.com/stretchr/testify/mock"

type MockContract struct {
	Mock mock.Mock
}

func (m *MockContract) Upload(base64String string, format string, filename string, path string) (stored_file_path string, failure error) {
	args := m.Mock.Called(base64String, format, filename, path)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockContract) Remove(path string) (failure error) {
	args := m.Mock.Called(path)
	return args.Error(0)
}
