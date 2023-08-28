package database

import (
	"context"
	"fmt"

	"github.com/SaulBenitez/inventory/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		s.DB.User,
		s.DB.Password,
		s.DB.Host,
		s.DB.Port,
		s.DB.Name,
	)

	fmt.Println("Database conexion URI", connectionString)

	return sqlx.ConnectContext(ctx, "mysql", connectionString)
}
