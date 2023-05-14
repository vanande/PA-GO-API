package libraries

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
)

func GetRole(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(privateKey), nil
	})

	fmt.Println(token)
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		if role, err := claims["Role"].(string); err {
			return role, nil
		} else {
			return "", errors.New("no role in token")
		}
	} else {
		return "", errors.New("invalid claims in token")
	}
}
