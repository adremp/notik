package utils

import (
	"net/http"
	"notik/pkg/httpErrors"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.MapClaims
	Id int32
	Email string
}

func GenerateToken(id int32, email string, expires time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Id: id,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(expires).Unix(),
		},
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, httpErrors.Error{http.StatusUnauthorized, "Invalid token"}
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
