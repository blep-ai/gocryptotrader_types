package huobi

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	"golang.org/x/time/rate"
)
type RateLimit struct {
	Spot          *rate.Limiter
	FuturesAuth   *rate.Limiter
	FuturesUnauth *rate.Limiter
	SwapAuth      *rate.Limiter
	SwapUnauth    *rate.Limiter
	FuturesXfer   *rate.Limiter
}
