package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrmissingKey = errors.New("jwt key not found")

type JWTClaims struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func NewJWTClaims(id int, role string) *JWTClaims {
	return &JWTClaims{
		Id:        id,
		Role:      role,
		Issuer:    os.Getenv("JWT_ISSUER"),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	}

}

func (j *JWTClaims) GenToken() (string, error) {
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		return "", ErrmissingKey
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))

}

func (j *JWTClaims) VerifyToken(token string) error {
	jwtToken, err := jwt.ParseWithClaims(token, j, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if !jwtToken.Valid {
		return jwt.ErrTokenExpired
	}
	issuer, err := jwtToken.Claims.GetIssuer()
	if err != nil {
		return err
	}
	if issuer != os.Getenv("JWT_ISSUER") {
		return jwt.ErrTokenInvalidIssuer
	}
	return nil
}
