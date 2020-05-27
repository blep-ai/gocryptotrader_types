package alphapoint

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/gorilla/websocket"
)

type Alphapoint struct {
	exchange.Base
	WebsocketConn *websocket.Conn
}
