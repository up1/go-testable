package api

import (
	"context"
	"github.com/labstack/echo"
	"model"
	"net/http"
	"service"
)

type LoginController struct {
	LService service.LoginService
}

type ResponseError struct {
	Message string `json:"message"`
}

func (lc *LoginController) Authentication(c echo.Context) error {
	name := c.Param("name")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	account, err := lc.LService.Login(ctx, name)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, account)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case model.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case model.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case model.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func NewLoginController(e *echo.Echo, ls service.LoginService) {
	controller := &LoginController{
		LService: ls,
	}
	e.GET("/authen/:name", controller.Authentication)
}
