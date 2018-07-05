package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"model"
	"service"
	"testing"
)

type MockAccountRepository struct {
}

func (m *MockAccountRepository) GetByName(ctx context.Context, name string) (*model.Account, error) {
	mockAccount := &model.Account{
		Id:    1,
		Name:  "Test Name 1",
		Email: "Test Email 1",
	}
	return mockAccount, nil
}

func TestLoginByName(t *testing.T) {
	expectedAccount := &model.Account{
		Id:    1,
		Name:  "Test Name 1",
		Email: "Test Email 1",
	}
	// Arrange
	l := service.NewLoginService(&MockAccountRepository{})

	// Act
	account, err := l.Login(context.TODO(), "Test Name 1")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, expectedAccount, account)
}
