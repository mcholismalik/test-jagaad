package usecase

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/test-jagaad/internal/entity"
)

// MockMockyDom is a mock implementation of the MockyDom interface
type MockMockyDom struct {
	mock.Mock
}

func (m *MockMockyDom) GetMocky1() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func (m *MockMockyDom) GetMocky2() ([]entity.User, error) {
	args := m.Called()
	return args.Get(0).([]entity.User), args.Error(1)
}

func TestUserUc_Fetch(t *testing.T) {
	os.Setenv("TMP_DIR", "../../")
	mockyDomMock := new(MockMockyDom)
	mockyDomMock.On("GetMocky1").Return([]entity.User{{Balance: "100", Tags: []string{"tag1"}}}, nil)
	mockyDomMock.On("GetMocky2").Return([]entity.User{}, errors.New("some error"))

	userUc := NewUserUc(mockyDomMock)
	assert.NotNil(t, userUc)

	resp, err := userUc.Fetch()

	assert.NoError(t, err)
	assert.Equal(t, entity.UserFilename, resp.Filename)
	assert.Len(t, resp.Details, 2)
	assert.Nil(t, resp.Details[0].Err)
	assert.NotNil(t, resp.Details[1].Err)

	mockyDomMock.AssertExpectations(t)
}

func TestUserUc_Search(t *testing.T) {
	os.Setenv("TMP_DIR", "../../")
	mockyDomMock := new(MockMockyDom)
	userUc := NewUserUc(mockyDomMock)
	assert.NotNil(t, userUc)

	// Test Search function
	resp, err := userUc.Search([]string{"tag1"})
	assert.NoError(t, err)
	assert.Len(t, resp, 1)

	mockyDomMock.AssertExpectations(t)
}
