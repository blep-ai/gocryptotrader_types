package kline

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
)
type Item struct {
	Exchange string
	Pair     currency.Pair
	Asset    asset.Item
	Interval time.Duration
	Candles  []Candle
}
type Candle struct {
	Time   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}
