package binance

import (
	"golang.org/x/time/rate"
)

type RateLimit struct {
	GlobalRate *rate.Limiter
	Orders     *rate.Limiter
}
