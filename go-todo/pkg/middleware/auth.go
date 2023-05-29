package middleware

import (
	"go-todo/dto"
	"net/http"
	"strings"

	jwtToken "go-todo/pkg/jwt"

	"github.com/labstack/echo/v4"
)

type Result struct {
	Status  int
	Data    interface{}
	Message string
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Status: http.StatusBadRequest, Message: "unauthorized"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, Result{Status: http.StatusUnauthorized, Message: "unathorized"})
		}

		c.Set("auth", claims)
		return next(c)
	}
}
