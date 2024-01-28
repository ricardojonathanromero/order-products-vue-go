package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
)

func NewCustomLogger(log logger.Logger) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.AddFields(map[string]any{
				"start_time":     v.StartTime,
				"request_id":     c.Request().Header.Get(echo.HeaderXRequestID),
				"host":           c.Request().Host,
				"content_length": v.ContentLength,
				"protocol":       c.Request().Proto,
				"real_ip":        c.RealIP(),
				"method":         c.Request().Method,
				"uri":            c.Request().RequestURI,
				"status":         v.Status,
			})
			log.Info("request")

			return nil
		},
	}
}
