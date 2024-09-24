package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

var jwtSecret = []byte("mysecret")

type RegisteredClaims struct{}

func AuthUser() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		userId := "6969"
		if username == "justin" && password == "sining" {
			claims := &jwt.RegisteredClaims{
				Subject:   userId,
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenString, err := token.SignedString(jwtSecret)
			if err != nil {
				return false, echo.ErrInternalServerError
			}

			c.JSON(http.StatusOK, map[string]string{
				"message": "Successfully authorized",
				"token":   "Bearer " + tokenString,
			})
			return true, nil
		}
		return false, echo.ErrUnauthorized
	})
}
func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: jwtSecret,
	})
}
