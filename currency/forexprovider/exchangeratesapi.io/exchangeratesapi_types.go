package exchangerates

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency/forexprovider/base"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
)
type ExchangeRates struct {
	base.Base
	Requester *request.Requester
}
type Rates struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}
type TimeSeriesRates struct {
	Base    string                 `json:"base"`
	StartAt string                 `json:"start_at"`
	EndAt   string                 `json:"end_at"`
	Rates   map[string]interface{} `json:"rates"`
}
