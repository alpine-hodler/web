package coinbasepro

import (
	"time"

	"golang.org/x/time/rate"
)

type ratelimiter uint8

const (
	candlesRateLimiter ratelimiter = iota
)

var ratelimiters = [uint8(1)]*rate.Limiter{}

func init() {
	ratelimiters[candlesRateLimiter] = nil
}

func getCandlesRateLimiter() *rate.Limiter {
	if ratelimiters[candlesRateLimiter] == nil {
		ratelimiters[candlesRateLimiter] = rate.NewLimiter(rate.Every(1*time.Second), 5)
	}
	return ratelimiters[candlesRateLimiter]
}
