package utils

import (
	"net/http"
	"notik/pkg/httpErrors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserClaims struct {
	jwt.MapClaims
	Id    int32  `json:"id"`
	Email string `json:"email"`
}

func GenerateToken(id int32, email string, expires time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Email: email,
		Id:    id,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(expires).Unix(),
		},
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, httpErrors.Error{http.StatusUnauthorized, "Invalid token"}
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, httpErrors.Error{http.StatusUnauthorized, "Invalid token"}
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, httpErrors.Error{http.StatusUnauthorized, "Invalid token"}
	}

	return claims, nil
}


func GetUserClaimsFromCtx(c echo.Context) (*UserClaims, error) {
	user, ok := c.Get("user").(*UserClaims)
	if !ok {
		return nil, httpErrors.Error{http.StatusUnauthorized, "Invalid token"}
	}
	return user, nil
}
