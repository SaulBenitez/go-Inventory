package api

import (
	"errors"
	"net/http"

	"github.com/SaulBenitez/inventory/internal/api/dtos"
	"github.com/SaulBenitez/inventory/internal/service"
	"github.com/labstack/echo/v4"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {

	ctx := c.Request().Context()
	params := dtos.RegisterUserRequest{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, ResponseMessage{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error"))
	}

	return c.JSON(http.StatusCreated, nil)
}
