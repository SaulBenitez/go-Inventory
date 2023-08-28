package main

import (
	"context"
	"fmt"

	"github.com/SaulBenitez/inventory/database"
	"github.com/SaulBenitez/inventory/internal/api"
	"github.com/SaulBenitez/inventory/internal/repository"
	"github.com/SaulBenitez/inventory/internal/service"
	"github.com/SaulBenitez/inventory/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx" // For dependency injection
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
