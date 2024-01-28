package http

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/ricardojonathanromero/order-products-vue-go/backend/auth/docs"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/app/handlers"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/constants"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/middlewares"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/transform"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/utils"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/validator"
	"github.com/swaggo/echo-swagger"
	"io"
	"time"
)

func useRateLimitMiddleware(e *echo.Echo) {
	// retrieve env vars
	rateLimitStr := utils.GetEnv(constants.ServerRateLimit, constants.Zero)
	burstStr := utils.GetEnv(constants.ServerBurst, constants.Zero)
	expiresInStr := utils.GetEnv(constants.ServerLimitExpireTime, constants.Zero)

	rateLimitInt := transform.StrToInt(rateLimitStr)
	burstInt := transform.StrToInt(burstStr)
	expiresInInt := transform.StrToInt(expiresInStr)

	rateLimitConf := middlewares.DefaultRateLimit
	if rateLimitInt > 0 {
		rateLimitConf.Rate = rateLimitInt
	}

	if burstInt > 0 {
		rateLimitConf.Burst = burstInt
	}

	if expiresInInt > 0 {
		rateLimitConf.ExpiresIn = time.Second * time.Duration(expiresInInt)
	}

	e.Use(middleware.RateLimiterWithConfig(middlewares.NewRateLimit(rateLimitConf)))
}

func NewServer(handles handlers.AuthHandler) (*echo.Echo, io.Closer) {
	var tracing io.Closer
	e := echo.New()

	// configure validator
	e.Validator = validator.NewValidator()

	// configure middlewares
	e.Use(middleware.LoggerWithConfig(middlewares.NewCustomLogger())) // configure logger
	e.Use(middlewares.NewRequestIdMiddleware())                       // set request id header
	if apiKey := utils.GetEnv(constants.ServerApiKey, constants.Empty); len(apiKey) > 0 {
		e.Use(middlewares.NewApiKey(apiKey))
	}
	useRateLimitMiddleware(e) // rate limit

	// jaeger
	if utils.GetEnv(constants.ServerEnableJaeger, constants.False) == constants.True {
		tracing = jaegertracing.New(e, nil)
	}

	// swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	auth := e.Group(constants.Version).Group(constants.Group)

	auth.POST("/login", handles.HandleLogin)
	auth.POST("/refresh-token", handles.HandleRefreshToken)

	return e, tracing
}
