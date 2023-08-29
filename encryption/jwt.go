package encryption

import (
	"github.com/SaulBenitez/inventory/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

func SignedLoginToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"name": u.Name,
	})

	return token.SignedString([]byte(key))
}
