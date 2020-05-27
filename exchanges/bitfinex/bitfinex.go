package bitfinex

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type Bitfinex struct {
	exchange.Base
	WebsocketConn              *wshandler.WebsocketConnection
	AuthenticatedWebsocketConn *wshandler.WebsocketConnection
	WebsocketSubdChannels      map[int]WebsocketChanInfo
}
