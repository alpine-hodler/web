package polygon

import (
	"time"

	"golang.org/x/time/rate"
)

// * This is a generated file, do not edit

type ratelimiter uint8

const (
	_ ratelimiter = iota
	UpcomingRatelimiter
)

var ratelimiters = [uint8(2)]*rate.Limiter{}

func init() { ratelimiters[UpcomingRatelimiter] = nil }

// getRateLimiter will load the rate limiter for a specific request, lazy loaded.
func getRateLimiter(rl ratelimiter, b int) *rate.Limiter {
	if ratelimiters[rl] == nil {
		ratelimiters[rl] = rate.NewLimiter(rate.Every(1*time.Second), b)
	}
	return ratelimiters[rl]
}
