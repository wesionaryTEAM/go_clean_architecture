package constants

import "time"

const (
	RateLimitPeriod   = 15 * time.Minute
	RateLimitRequests = int64(200)
)
