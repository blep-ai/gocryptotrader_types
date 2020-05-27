package coinbasepro

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type CoinbasePro struct {
	exchange.Base
	WebsocketConn *wshandler.WebsocketConnection
}
