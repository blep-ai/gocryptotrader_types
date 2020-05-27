package engine
import (
	"errors"
	"strings"
	"sync"

	"github.com/blep-ai/gocryptotrader_types/common"
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/binance"
	"github.com/blep-ai/gocryptotrader_types/exchanges/bitfinex"
	"github.com/blep-ai/gocryptotrader_types/exchanges/bitflyer"
	"github.com/blep-ai/gocryptotrader_types/exchanges/bithumb"
	"github.com/blep-ai/gocryptotrader_types/exchanges/bitmex"
	"github.com/blep-ai/gocryptotrader_types/exchanges/bitstamp"
	"github.com/blep-ai/gocryptotrader_types/exchanges/bittrex"
	"github.com/blep-ai/gocryptotrader_types/exchanges/btcmarkets"
	"github.com/blep-ai/gocryptotrader_types/exchanges/btse"
	"github.com/blep-ai/gocryptotrader_types/exchanges/coinbasepro"
	"github.com/blep-ai/gocryptotrader_types/exchanges/coinbene"
	"github.com/blep-ai/gocryptotrader_types/exchanges/coinut"
	"github.com/blep-ai/gocryptotrader_types/exchanges/exmo"
	"github.com/blep-ai/gocryptotrader_types/exchanges/gateio"
	"github.com/blep-ai/gocryptotrader_types/exchanges/gemini"
	"github.com/blep-ai/gocryptotrader_types/exchanges/hitbtc"
	"github.com/blep-ai/gocryptotrader_types/exchanges/huobi"
	"github.com/blep-ai/gocryptotrader_types/exchanges/itbit"
	"github.com/blep-ai/gocryptotrader_types/exchanges/kraken"
	"github.com/blep-ai/gocryptotrader_types/exchanges/lakebtc"
	"github.com/blep-ai/gocryptotrader_types/exchanges/lbank"
	"github.com/blep-ai/gocryptotrader_types/exchanges/localbitcoins"
	"github.com/blep-ai/gocryptotrader_types/exchanges/okcoin"
	"github.com/blep-ai/gocryptotrader_types/exchanges/okex"
	"github.com/blep-ai/gocryptotrader_types/exchanges/poloniex"
	"github.com/blep-ai/gocryptotrader_types/exchanges/yobit"
	"github.com/blep-ai/gocryptotrader_types/exchanges/zb"
	"github.com/blep-ai/gocryptotrader_types/log"
)
type exchangeManager struct {
	m         sync.Mutex
	exchanges map[string]exchange.IBotExchange
}
