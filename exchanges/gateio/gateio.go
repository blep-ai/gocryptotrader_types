package gateio

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type Gateio struct {
	WebsocketConn *wshandler.WebsocketConnection
	exchange.Base
}
