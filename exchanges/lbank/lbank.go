package lbank

import (
	"crypto/rsa"

	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type Lbank struct {
	exchange.Base
	privateKey    *rsa.PrivateKey
	WebsocketConn *wshandler.WebsocketConnection
}
