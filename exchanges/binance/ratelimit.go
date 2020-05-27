package binance

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"golang.org/x/time/rate"
)
type RateLimit struct {
	GlobalRate *rate.Limiter
	Orders     *rate.Limiter
}
