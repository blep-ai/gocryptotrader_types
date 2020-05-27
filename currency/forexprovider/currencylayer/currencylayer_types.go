package currencylayer

import (
	"github.com/thrasher-corp/gocryptotrader/currency/forexprovider/base"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
)
type CurrencyLayer struct {
	base.Base
	Requester *request.Requester
}
type Error struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}
type LiveRates struct {
	Success   bool               `json:"success"`
	Error     Error              `json:"error"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Timestamp int64              `json:"timestamp"`
	Source    string             `json:"source"`
	Quotes    map[string]float64 `json:"quotes"`
}
type SupportedCurrencies struct {
	Success    bool              `json:"success"`
	Error      Error             `json:"error"`
	Terms      string            `json:"terms"`
	Privacy    string            `json:"privacy"`
	Currencies map[string]string `json:"currencies"`
}
type HistoricalRates struct {
	Success    bool               `json:"success"`
	Error      Error              `json:"error"`
	Terms      string             `json:"terms"`
	Privacy    string             `json:"privacy"`
	Historical bool               `json:"historical"`
	Date       string             `json:"date"`
	Timestamp  int64              `json:"timestamp"`
	Source     string             `json:"source"`
	Quotes     map[string]float64 `json:"quotes"`
}
type ConversionRate struct {
	Success bool   `json:"success"`
	Error   Error  `json:"error"`
	Privacy string `json:"privacy"`
	Terms   string `json:"terms"`
	Query   struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	} `json:"query"`
	Info struct {
		Timestamp int64   `json:"timestamp"`
		Quote     float64 `json:"quote"`
	} `json:"info"`
	Historical bool    `json:"historical"`
	Date       string  `json:"date"`
	Result     float64 `json:"result"`
}
type TimeFrame struct {
	Success   bool                   `json:"success"`
	Error     Error                  `json:"error"`
	Terms     string                 `json:"terms"`
	Privacy   string                 `json:"privacy"`
	Timeframe bool                   `json:"timeframe"`
	StartDate string                 `json:"start_date"`
	EndDate   string                 `json:"end_date"`
	Source    string                 `json:"source"`
	Quotes    map[string]interface{} `json:"quotes"`
}
type ChangeRate struct {
	Success   bool               `json:"success"`
	Error     Error              `json:"error"`
	Terms     string             `json:"terms"`
	Privacy   string             `json:"privacy"`
	Change    bool               `json:"change"`
	StartDate string             `json:"start_date"`
	EndDate   string             `json:"end_date"`
	Source    string             `json:"source"`
	Quotes    map[string]Changes `json:"quotes"`
}
type Changes struct {
	StartRate float64 `json:"start_rate"`
	EndRate   float64 `json:"end_rate"`
	Change    float64 `json:"change"`
	ChangePCT float64 `json:"change_pct"`
}
