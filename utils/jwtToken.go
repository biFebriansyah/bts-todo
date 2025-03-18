package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtClaims struct {
	Id string `json:"_id"`
	jwt.RegisteredClaims
}

func GenerateJwt(userId string) (string, error) {
	claims := jwtClaims{
		Id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gotodo",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}

	secrets := os.Getenv("JWT_SECRETS")
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return tokens.SignedString([]byte(secrets))
}

func ParseJwt(token string) (*jwtClaims, error) {
	secrets := os.Getenv("JWT_SECRETS")
	jwtData, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secrets), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail parse jwt: %w", err)
	}

	claimData := jwtData.Claims.(*jwtClaims)
	return claimData, nil
}
