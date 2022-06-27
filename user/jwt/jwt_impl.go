package jwt

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type JWTImpl struct {
	request *http.Request
}

func NewJWT() JWT {
	return JWTImpl{}
}

type Payload struct {
	Id string `json:"id"`
}

func (jwtAuth JWTImpl) GenerateToken(payload Payload) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": payload.Id,
	})

	tokenString, err := token.SignedString([]byte("keytest"))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
