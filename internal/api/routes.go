package api

import "github.com/labstack/echo/v4"

func (a *API) RegisterRoutes(e *echo.Echo) {
	users := e.Group("/users")
	users.POST("/", a.RegisterUser)
	users.POST("/register_role", a.RegisterUserRole)
	users.POST("/login", a.LoginUser)

	products := e.Group("/products")
	products.POST("/", a.RegisterProduct)
}
