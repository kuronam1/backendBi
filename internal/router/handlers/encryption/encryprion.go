package encryption

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

var key = []byte("myKey")

var (
	ParsingErr  = errors.New("error while parsing")
	CreatingErr = errors.New("error while creating a token")
	NotValid    = errors.New("not valid token")
)

func MakeToken(id int) (string, error) {
	payload := jwt.MapClaims{
		"id": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", CreatingErr
	}

	return tokenString, nil
}

func ParsingToken(token string) (int, error) {
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token2 *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, ParsingErr
	}

	id, ok := claims["id"].(float64)
	if !ok {
		return 0, ParsingErr
	}

	if jwtToken.Valid {
		return int(id), nil
	} else {
		return 0, NotValid
	}
}
