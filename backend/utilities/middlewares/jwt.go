package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/ricardojonathanromero/orders-products-vue-go/backend/utilities/jwt"
	"net/http"
	"strings"
)

func NewAuth0ValidatorMiddleware(validator jwt.Auth0Validator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			token := request.Header.Get(echo.HeaderAuthorization)

			if token == "" {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":    "unauthorized",
					"message": "invalid token",
					"details": []string{"header 'Authorization' is required"},
				})
			}

			if strings.Contains(token, "Bearer ") {
				token = strings.ReplaceAll(token, "Bearer ", "")
			}

			err := validator.SetContext(request.Context())
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":    "unauthorized",
					"message": "invalid token",
					"details": []string{err.Error()},
				})
			}

			claims, err := validator.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"code":    "unauthorized",
					"message": "invalid token",
					"details": []string{err.Error()},
				})
			}

			c.Set("user_id", claims.Sub)
			return next(c)
		}
	}
}
