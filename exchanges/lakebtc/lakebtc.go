package lakebtc

import (
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
)

type LakeBTC struct {
	exchange.Base
	WebsocketConn
}
