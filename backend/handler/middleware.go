package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			c.JSON(
				http.StatusUnauthorized,
				map[string]string{"error": "Authorization Header missing"},
			)
		}
		claims, err := ValidateJWTToken(authToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}
		c.Set("user", claims)
		return next(c)
	}
}