package libraries

import (
	"github.com/golang-jwt/jwt"
	"time"
)

const privateKey = "42"

func CreateToken(role string, id int) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "TaS offical golang server",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims.(jwt.MapClaims)["role"] = role
	token.Claims.(jwt.MapClaims)["id"] = id

	return token.SignedString([]byte(privateKey))
}
