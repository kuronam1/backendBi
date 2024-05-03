package encryption

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	creatingErr = errors.New("error while creating a token")
)

func MakeToken(id int) (string, error) {
	payload := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SigningString()
	if err != nil {
		return "", creatingErr
	}

	return tokenString, err
}
