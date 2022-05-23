package middlewares

import (
	"clean-architecture/constants"
	"clean-architecture/lib"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

var store = memory.NewStoreWithOptions(limiter.StoreOptions{CleanUpInterval: 10 * time.Second})

type RateLimitMiddleware struct {
	logger lib.Logger
}

func NewRateLimitMiddleware(logger lib.Logger) RateLimitMiddleware {
	return RateLimitMiddleware{
		logger: logger,
	}
}

func (lm RateLimitMiddleware) Handle(period time.Duration, limit int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()

		lm.logger.Info("Setting up rate limit middleware")
		rate := limiter.Rate{
			Period: period,
			Limit:  limit,
		}

		// using in-memory store with goroutine which clears expired keys.
		instance := limiter.New(store, rate)

		context, err := instance.Get(c, key)

		if err != nil {
			lm.logger.Panic(err.Error())
		}

		c.Set(constants.RateLimit, instance)
		c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

		if context.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate Limit Exceed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// DefaultKeyGetter is the default KeyGetter used by a new Middleware.
// It returns the Client IP address.
func DefaultKeyGetter(c *gin.Context) string {
	return c.ClientIP()
}
