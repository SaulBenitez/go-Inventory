package main

import (
	"context"

	"github.com/SaulBenitez/inventory/database"
	"github.com/SaulBenitez/inventory/internal/repository"
	"github.com/SaulBenitez/inventory/internal/service"
	"github.com/SaulBenitez/inventory/settings"
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
		),
		fx.Invoke(
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "my@email.com", "myname", "mypassword")
				if err != nil {
					panic(err)
				}

				u, err := serv.LoginUser(ctx, "my@email.com", "mypassword")
				if err != nil {
					panic(err)
				}
				if u.Name != "myname" {
					panic("wrong user name")
				}
			},
		),
	)

	app.Run()
}
