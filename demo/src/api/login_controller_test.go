package api_test

import (
	"context"
	"errors"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api"
	"model"
)

type mockAccountService struct{}

func (m *mockAccountService) Login(ctx context.Context, name string) (*model.Account, error) {
	mockAccount := &model.Account{
		Id:    1,
		Name:  "Test Name 1",
		Email: "Test Email 1",
	}
	return mockAccount, nil
}

func TestAuthenticationSuccess(t *testing.T) {
	mockService := new(mockAccountService)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/authen/somkiat", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("authen/:name")
	c.SetParamNames("name")
	c.SetParamValues("somkiat")
	controller := api.LoginController{
		LService: mockService,
	}
	controller.Authentication(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

type mockAccountServiceFail struct{}

func (m *mockAccountServiceFail) Login(ctx context.Context, name string) (*model.Account, error) {
	return nil, errors.New("Unexpexted Error")
}

func TestAuthenticationFail(t *testing.T) {
	mockService := new(mockAccountServiceFail)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/authen/somkiat", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("authen/:name")
	c.SetParamNames("name")
	c.SetParamValues("somkiat")
	controller := api.LoginController{
		LService: mockService,
	}
	controller.Authentication(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

type mockAccountServiceFail2 struct {
	status int
}

func (m *mockAccountServiceFail2) Login(ctx context.Context, name string) (*model.Account, error) {
	switch m.status {
	case 404:
		return nil, model.NOT_FOUND_ERROR
	default:
		return nil, model.INTERNAL_SERVER_ERROR
	}
}

func TestAuthenticationNotFound(t *testing.T) {
	mockService := &mockAccountServiceFail2{status: http.StatusNotFound}

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/authen/somkiat", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("authen/:name")
	c.SetParamNames("name")
	c.SetParamValues("somkiat")
	controller := api.LoginController{
		LService: mockService,
	}
	controller.Authentication(c)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestAuthenticationINTERNAL_SERVER_ERROR(t *testing.T) {
	mockService := &mockAccountServiceFail2{status: http.StatusInternalServerError}

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/authen/somkiat", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("authen/:name")
	c.SetParamNames("name")
	c.SetParamValues("somkiat")
	controller := api.LoginController{
		LService: mockService,
	}
	controller.Authentication(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
