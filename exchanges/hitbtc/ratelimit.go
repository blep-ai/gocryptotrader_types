package hitbtc

import (
	"golang.org/x/time/rate"
)

type RateLimit struct {
	MarketData *rate.Limiter
	Trading    *rate.Limiter
	Other      *rate.Limiter
}
