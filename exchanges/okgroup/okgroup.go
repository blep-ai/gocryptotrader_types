package okgroup

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type OKGroup struct {
	exchange.Base
	ExchangeName  string
	WebsocketConn *wshandler.WebsocketConnection
	// Spot and contract market error codes as per https://www.okex.com/rest_request.html
	ErrorCodes map[string]error
	// Stores for corresponding variable checks
	ContractTypes         []string
	CurrencyPairsDefaults []string
	ContractPosition      []string
	Types                 []string
	// URLs to be overridden by implementations of OKGroup
	APIURL       string
	APIVersion   string
	WebsocketURL string
}
