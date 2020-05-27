package hitbtc

import (
	"errors"
	"time"

	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"golang.org/x/time/rate"
)
type RateLimit struct {
	MarketData *rate.Limiter
	Trading    *rate.Limiter
	Other      *rate.Limiter
}
