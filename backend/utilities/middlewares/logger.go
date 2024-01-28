package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ricardojonathanromero/orders-products-vue-go/backend/utilities/logger"
	"time"
)

func NewCustomLogger(log logger.Logger) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.AddFields(map[string]any{
				"StartTime":     v.StartTime,
				"EndTime":       time.Now(),
				"RequestId":     v.RequestID,
				"Host":          v.Host,
				"ContentLength": v.ContentLength,
				"Protocol":      v.Protocol,
				"RealIp":        v.RemoteIP,
				"Method":        v.Method,
				"URI":           v.URI,
				"Status":        v.Status,
			})
			log.Info("request")

			return nil
		},
	}
}
