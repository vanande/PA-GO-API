package libraries

import (
	"github.com/golang-jwt/jwt"
	"time"
)

const privateKey = "42"

type UserInfo struct {
	Name string
	Kind string
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	TokenType string
	UserInfo
}

func CreateToken(user string, id int) (string, error) {

	t := jwt.New(jwt.SigningMethodHS256)

	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Id:        string(rune(id)),
		},
		"level1",
		UserInfo{user, "human"},
	}

	return t.SignedString([]byte(privateKey))
}
