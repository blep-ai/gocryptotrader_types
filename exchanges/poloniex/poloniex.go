package poloniex

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type Poloniex struct {
	exchange.Base
	WebsocketConn *wshandler.WebsocketConnection
}
