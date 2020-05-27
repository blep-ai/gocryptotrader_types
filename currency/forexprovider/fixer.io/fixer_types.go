package fixer

import (
	"github.com/blep-ai/gocryptotrader_types/currency/forexprovider/base"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
)

type Fixer struct {
	base.Base
	Requester *request.Requester
}
type Rates struct {
	Success    bool               `json:"success"`
	Error      RespError          `json:"error"`
	Historical bool               `json:"historical"`
	Timestamp  int64              `json:"timestamp"`
	Base       string             `json:"base"`
	Date       string             `json:"date"`
	Rates      map[string]float64 `json:"rates"`
}
type Conversion struct {
	Success bool      `json:"success"`
	Error   RespError `json:"error"`
	Query   struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	} `json:"query"`
	Info struct {
		Timestamp int64   `json:"timestamp"`
		Rate      float64 `json:"rate"`
	} `json:"info"`
	Historical bool    `json:"historical"`
	Date       string  `json:"date"`
	Result     float64 `json:"result"`
}
type TimeSeries struct {
	Success    bool                   `json:"success"`
	Error      RespError              `json:"error"`
	Timeseries bool                   `json:"timeseries"`
	StartDate  string                 `json:"start_date"`
	EndDate    string                 `json:"end_date"`
	Base       string                 `json:"base"`
	Rates      map[string]interface{} `json:"rates"`
}
type Fluctuation struct {
	Success     bool            `json:"success"`
	Error       RespError       `json:"error"`
	Fluctuation bool            `json:"fluctuation"`
	StartDate   string          `json:"start_date"`
	EndDate     string          `json:"end_date"`
	Base        string          `json:"base"`
	Rates       map[string]Flux `json:"rates"`
}
type Flux struct {
	StartRate float64 `json:"start_rate"`
	EndRate   float64 `json:"end_rate"`
	Change    float64 `json:"change"`
	ChangePCT float64 `json:"change_pct"`
}
type RespError struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Info string `json:"info"`
}
type Symbols struct {
	Success bool              `json:"success"`
	Error   RespError         `json:"error"`
	Map     map[string]string `json:"symbols"`
}
