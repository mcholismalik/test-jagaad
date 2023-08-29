package cmd

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/test-jagaad/internal/entity"
)

// MockUserUc is a mock implementation of the UserUc interface
type MockUserUc struct {
	mock.Mock
}

func (m *MockUserUc) Fetch() (entity.FetchUserResp, error) {
	args := m.Called()
	return args.Get(0).(entity.FetchUserResp), args.Error(1)
}

func (m *MockUserUc) Search(tags []string) ([]string, error) {
	args := m.Called(tags)
	return args.Get(0).([]string), args.Error(1)
}

func TestCommand_Fetch(t *testing.T) {
	userUcMock := new(MockUserUc)
	expectedResponse := entity.FetchUserResp{
		Details:  []entity.FetchUserRespDetail{{Name: "Endpoint1"}},
		Filename: "users.csv",
	}
	userUcMock.On("Fetch").Return(expectedResponse, nil)

	cmd := &command{userUc: userUcMock}
	cmd.fetch()

	userUcMock.AssertExpectations(t)
}

func TestCommand_Search(t *testing.T) {
	userUcMock := new(MockUserUc)
	expectedResponse := []string{"User1", "User2"}
	userUcMock.On("Search", mock.Anything).Return(expectedResponse, nil)

	cmd := &command{userUc: userUcMock}
	cmd.search([]string{"--tags=tag1,tag2"})

	userUcMock.AssertExpectations(t)
}
