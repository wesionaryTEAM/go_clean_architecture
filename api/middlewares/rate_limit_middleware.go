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

// Global store
// using in-memory store with goroutine which clears expired keys.
var store = memory.NewStore()

// Default values
const RATE_LIMIT_PERIOD = 15 * time.Minute
const RATE_LIMIT_RATE = int64(200)

type RateLimitOption struct {
	period time.Duration
	limit  int64
}

type RateLimitMiddleware struct {
	logger lib.Logger
	option RateLimitOption
}

func NewRateLimitMiddleware(logger lib.Logger) RateLimitMiddleware {
	return RateLimitMiddleware{
		logger: logger,
		option: RateLimitOption{
			period: RATE_LIMIT_PERIOD,
			limit:  RATE_LIMIT_RATE,
		},
	}
}

func (lm RateLimitMiddleware) Handle(option *RateLimitOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP() // Gets cient IP Address

		lm.logger.Info("Setting up rate limit middleware")

		// Setting up rate limit
		// Limit -> # of API Calls
		// Period -> in a given time frame
		// setting default value
		rate := limiter.Rate{
			Limit:  lm.option.limit,
			Period: lm.option.period,
		}

		if option != nil {
			rate.Limit = option.limit
			rate.Period = option.period
		}

		// Limiter instance
		instance := limiter.New(store, rate)

		// Returns the rate limit details for given identifier.
		// FullPath is appended with IP address. `/api/users&&127.0.0.1` as key
		context, err := instance.Get(c, c.FullPath()+"&&"+key)

		if err != nil {
			lm.logger.Panic(err.Error())
		}

		c.Set(constants.RateLimit, instance)

		// Setting custom headers
		c.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(context.Reset, 10))

		// Limit exceeded
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
