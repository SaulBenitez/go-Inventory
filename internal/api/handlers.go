package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/SaulBenitez/inventory/encryption"
	"github.com/SaulBenitez/inventory/internal/api/dtos"
	"github.com/SaulBenitez/inventory/internal/models"
	"github.com/SaulBenitez/inventory/internal/service"
	"github.com/labstack/echo/v4"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

// TODO: As far as possible, refactor these handlers

func (a *API) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.LogingUserRequest{}

	err := c.Bind(&params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: err.Error()})
	}

	u, err := a.serv.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusUnauthorized, ResponseMessage{Message: "Internal server error"})
	}

	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, ResponseMessage{Message: "Internal server error"})
	}

	cookie := &http.Cookie{
		Name: "Authorization",
		Value: token,
		Secure: true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"success": "true"})
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

	return c.JSON(http.StatusCreated, ResponseMessage{Message: "User was created sucessfully"})
}

func (a *API) RegisterUserRole(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUserRoleRequest{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: err.Error()})
	}

	err = a.serv.AddUserRole(ctx, params.UserID, params.RoleID)
	log.Printf("%v", err.Error())
	log.Printf("%v", params)
	if err != nil {
		if err == service.ErrRoleAlreadyAdded {
			return c.JSON(http.StatusConflict, ResponseMessage{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, ResponseMessage{Message: "Unexpected error"})
	}

	return c.JSON(http.StatusCreated, ResponseMessage{Message: "Product was created sucessfully"})
}

func (a *API) RegisterProduct(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterProductRequest{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: "invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseMessage{Message: err.Error()})
	}

	//TODO: Get user email from JWT or other auth method
	p := models.Product{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
	}
	err = a.serv.AddProduct(ctx, p, params.UserEmail)
	if err != nil {
		if err == service.ErrInvalidPermissions {
			return c.JSON(http.StatusForbidden, ResponseMessage{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, ResponseMessage{Message: "Unexpected error"})
	}

	return c.JSON(http.StatusCreated, ResponseMessage{Message: "Product was created sucessfully"})
}
