package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type RateLimitOpts struct {
	Rate      int
	Burst     int
	ExpiresIn time.Duration
}

var DefaultRateLimit = RateLimitOpts{
	Rate:      100,
	Burst:     10,
	ExpiresIn: 1 * time.Minute,
}

/*
NewRateLimit configures the options to limit the request based on env
variables or the default values will be configured.
*/
func NewRateLimit(opts RateLimitOpts) middleware.RateLimiterConfig {
	return middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(opts.Rate), Burst: opts.Burst, ExpiresIn: opts.ExpiresIn},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusForbidden, map[string]any{
				"code":    "forbidden",
				"message": "error while extracting identifier",
				"details": []string{err.Error()},
			})
		},
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return c.JSON(http.StatusTooManyRequests, map[string]any{
				"code":    "too_many_request",
				"message": "rate limit exceeded",
				"details": []string{err.Error()},
			})
		},
	}
}
