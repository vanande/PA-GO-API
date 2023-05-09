package libraries

import (
	"errors"
	"net/http"
	"strings"
)

func GetTokenFromHeader(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header is missing")
	}

	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 {
		return "", errors.New("Invalid Authorization header format")
	}

	token := splitToken[1]
	return token, nil
}
