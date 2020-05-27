package binance

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	"golang.org/x/time/rate"
)
type RateLimit struct {
	GlobalRate *rate.Limiter
	Orders     *rate.Limiter
}
