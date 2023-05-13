package libraries

import (
<<<<<<< HEAD
	"time"

	"github.com/golang-jwt/jwt"
=======
	"github.com/golang-jwt/jwt"
	"time"
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
)

const privateKey = "42"

<<<<<<< HEAD
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
			Id:        string(id),
		},
		"level1",
		UserInfo{user, "human"},
	}

	return t.SignedString([]byte(privateKey))
=======
func CreateToken(role string, id int) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    "TaS offical golang server",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Claims.(jwt.MapClaims)["role"] = role
	token.Claims.(jwt.MapClaims)["id"] = id

	return token.SignedString([]byte(privateKey))
>>>>>>> 24452d1a9b3dcd7ccc9c4f6bc6a865ae32926d2c
}
