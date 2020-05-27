package hitbtc

import (
	"errors"
	"time"

	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	"golang.org/x/time/rate"
)
type RateLimit struct {
	MarketData *rate.Limiter
	Trading    *rate.Limiter
	Other      *rate.Limiter
}
