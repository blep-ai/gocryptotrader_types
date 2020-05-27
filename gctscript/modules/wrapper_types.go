package modules

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/account"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/blep-ai/gocryptotrader_types/exchanges/kline"
	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
	"github.com/blep-ai/gocryptotrader_types/exchanges/orderbook"
	"github.com/blep-ai/gocryptotrader_types/exchanges/ticker"
	"github.com/blep-ai/gocryptotrader_types/portfolio/withdraw"
)

type GCT interface {
	Exchange
}
type Exchange interface {
	Exchanges(enabledOnly bool) []string
	IsEnabled(exch string) bool
	Orderbook(exch string, pair currency.Pair, item asset.Item) (*orderbook.Base, error)
	Ticker(exch string, pair currency.Pair, item asset.Item) (*ticker.Price, error)
	Pairs(exch string, enabledOnly bool, item asset.Item) (*currency.Pairs, error)
	QueryOrder(exch, orderid string) (*order.Detail, error)
	SubmitOrder(submit *order.Submit) (*order.SubmitResponse, error)
	CancelOrder(exch, orderid string) (bool, error)
	AccountInformation(exch string) (account.Holdings, error)
	DepositAddress(exch string, currencyCode currency.Code) (string, error)
	WithdrawalFiatFunds(exch, bankAccountID string, request *withdraw.Request) (out string, err error)
	WithdrawalCryptoFunds(exch string, request *withdraw.Request) (out string, err error)
	OHLCV(exch string, pair currency.Pair, item asset.Item, start, end time.Time, interval time.Duration) (kline.Item, error)
}
