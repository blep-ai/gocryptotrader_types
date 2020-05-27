package binance

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type Binance struct {
	exchange.Base
	WebsocketConn *wshandler.WebsocketConnection

	// Valid string list that is required by the exchange
	validLimits    []int
	validIntervals []TimeInterval
}
