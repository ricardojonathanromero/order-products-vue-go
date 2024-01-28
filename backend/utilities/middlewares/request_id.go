package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func NewRequestIdMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(echo.HeaderXRequestID, uuid.NewString())
			return next(c)
		}
	}
}
