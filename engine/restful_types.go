package engine

import (
	"net/http"

	"github.com/thrasher-corp/gocryptotrader/exchanges/account"
	"github.com/thrasher-corp/gocryptotrader/exchanges/orderbook"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
)
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type AllEnabledExchangeOrderbooks struct {
	Data []EnabledExchangeOrderbooks `json:"data"`
}
type EnabledExchangeOrderbooks struct {
	ExchangeName   string           `json:"exchangeName"`
	ExchangeValues []orderbook.Base `json:"exchangeValues"`
}
type AllEnabledExchangeCurrencies struct {
	Data []EnabledExchangeCurrencies `json:"data"`
}
type EnabledExchangeCurrencies struct {
	ExchangeName   string         `json:"exchangeName"`
	ExchangeValues []ticker.Price `json:"exchangeValues"`
}
type AllEnabledExchangeAccounts struct {
	Data []account.Holdings `json:"data"`
}
