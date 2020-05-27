package bitmex

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type Bitmex struct {
	exchange.Base
	WebsocketConn *wshandler.WebsocketConnection
}
