package kline

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
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
