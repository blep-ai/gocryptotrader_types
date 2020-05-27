package currencyconverter

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency/forexprovider/base"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
)
type CurrencyConverter struct {
	base.Base
	Requester *request.Requester
}
type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
type CurrencyItem struct {
	CurrencyName   string `json:"currencyName"`
	CurrencySymbol string `json:"currencySymbol"`
	ID             string `json:"ID"`
}
type Currencies struct {
	Results map[string]CurrencyItem
}
type CountryItem struct {
	Alpha3         string `json:"alpha3"`
	CurrencyID     string `json:"currencyId"`
	CurrencyName   string `json:"currencyName"`
	CurrencySymbol string `json:"currencySymbol"`
	ID             string `json:"ID"`
	Name           string `json:"Name"`
}
type Countries struct {
	Results map[string]CountryItem
}
