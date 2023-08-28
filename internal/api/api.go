package api

import (
	"github.com/SaulBenitez/inventory/internal/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	serv service.Service
}

func New(serv service.Service) *API {
	return &API{serv: serv}
}

func (a *API) Start(e *echo.Echo, address string) error {
	return e.Start(address)
}