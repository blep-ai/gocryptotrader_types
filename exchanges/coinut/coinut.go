package coinut

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type COINUT struct {
	exchange.Base
	WebsocketConn *wshandler.WebsocketConnection
	instrumentMap instrumentMap
}
