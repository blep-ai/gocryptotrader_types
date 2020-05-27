package engine

import (
	"errors"
	"strings"
	"sync"

	"github.com/thrasher-corp/gocryptotrader/common"
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/binance"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitfinex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitflyer"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bithumb"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitmex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bitstamp"
	"github.com/thrasher-corp/gocryptotrader/exchanges/bittrex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/btcmarkets"
	"github.com/thrasher-corp/gocryptotrader/exchanges/btse"
	"github.com/thrasher-corp/gocryptotrader/exchanges/coinbasepro"
	"github.com/thrasher-corp/gocryptotrader/exchanges/coinbene"
	"github.com/thrasher-corp/gocryptotrader/exchanges/coinut"
	"github.com/thrasher-corp/gocryptotrader/exchanges/exmo"
	"github.com/thrasher-corp/gocryptotrader/exchanges/gateio"
	"github.com/thrasher-corp/gocryptotrader/exchanges/gemini"
	"github.com/thrasher-corp/gocryptotrader/exchanges/hitbtc"
	"github.com/thrasher-corp/gocryptotrader/exchanges/huobi"
	"github.com/thrasher-corp/gocryptotrader/exchanges/itbit"
	"github.com/thrasher-corp/gocryptotrader/exchanges/kraken"
	"github.com/thrasher-corp/gocryptotrader/exchanges/lakebtc"
	"github.com/thrasher-corp/gocryptotrader/exchanges/lbank"
	"github.com/thrasher-corp/gocryptotrader/exchanges/localbitcoins"
	"github.com/thrasher-corp/gocryptotrader/exchanges/okcoin"
	"github.com/thrasher-corp/gocryptotrader/exchanges/okex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/poloniex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/yobit"
	"github.com/thrasher-corp/gocryptotrader/exchanges/zb"
	"github.com/thrasher-corp/gocryptotrader/log"
)
type exchangeManager struct {
	m         sync.Mutex
	exchanges map[string]exchange.IBotExchange
}
