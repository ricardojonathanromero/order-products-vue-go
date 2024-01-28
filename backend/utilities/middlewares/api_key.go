package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

/*
NewApiKey configures the options to implements an api key.
*/
func NewApiKey(apiKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("X-Api-Key")

			if !strings.EqualFold(key, apiKey) {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":    "invalid_api_key",
					"message": "the current api key is not valid",
					"details": []string{"check your request to validate if 'x-api-key' is set and if it's valid"},
				})
			}

			return next(c)
		}
	}
}
