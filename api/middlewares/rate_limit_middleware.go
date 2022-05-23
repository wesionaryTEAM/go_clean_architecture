package middlewares

import (
	"clean-architecture/constants"
	"clean-architecture/lib"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

type limiterRate struct {
	period time.Duration // for given time period
	limit  int64         // number of requests
}

type RateLimitMiddleware struct {
	logger lib.Logger
}

func NewRateLimitMiddleware(logger lib.Logger) RateLimitMiddleware {
	return RateLimitMiddleware{
		logger: logger,
	}
}

func (lm RateLimitMiddleware) Handle(limit limiterRate, limitOptions limiter.Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		lm.logger.Info("Setting up rate limit middleware")
		rate := limiter.Rate{
			Period: limit.period,
			Limit:  limit.limit,
		}

		// using in-memory store with goroutine which clears expired keys.
		store := memory.NewStore()

		instance := limiter.New(store, rate, limitOptions)

		c.Set(constants.RateLimit, instance)
		c.Next()
	}
}
