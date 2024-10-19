package auth

import (
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/golang-jwt/jwt/v4"
	"time"
)

var JWTkey = []byte("tgyD3YLKOSHILWEW9QA6H2MUsfAs36/28+FyFC5Fqj4=")

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func JwtGnarator(name string, email string) (string, error) {
	expiretime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiretime),
			Issuer:    "last-api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTkey)
}
