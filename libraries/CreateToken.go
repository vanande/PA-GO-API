package libraries

import (
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

const privateKey = "42"

type UserInfo struct {
	Id   string
	Role string
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	UserInfo
}

func CreateToken(role string, id int) (string, error) {

	t := jwt.New(jwt.SigningMethodHS256)

	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 90).Unix(),
			Issuer:    "TogetherAndStronger",
		},
		"level1",
		UserInfo{strconv.Itoa(id), role},
	}

	return t.SignedString([]byte(privateKey))
}
