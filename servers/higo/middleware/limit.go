package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/sereiner/parrot/servers/higo"
)

type LimitConfig struct {
	// Skipper defines a function to skip middleware.
	Skipper Skipper

	// lmt
	lmt *limiter.Limiter
}

var (
	DefaultLimitConfig = LimitConfig{
		Skipper: DefaultSkipper,
		lmt:     tollbooth.NewLimiter(2000, nil),
	}
)

func Limit() higo.MiddlewareFunc {
	return LimitWithConfig(DefaultLimitConfig)
}

func LimitWithConfig(config LimitConfig) higo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultSkipper
	}
	if config.lmt == nil {
		config.lmt = DefaultLimitConfig.lmt
	}

	return func(next higo.HandlerFunc) higo.HandlerFunc {
		return higo.HandlerFunc(func(c higo.Context) error {
			httpError := tollbooth.LimitByRequest(config.lmt, c.Response(), c.Request())
			if httpError != nil {
				return c.String(httpError.StatusCode, httpError.Message)
			}
			return next(c)
		})
	}
}
