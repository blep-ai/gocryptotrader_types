package exchange

import (
	"sync"
	"time"

	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/account"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/blep-ai/gocryptotrader_types/exchanges/kline"
	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
	"github.com/blep-ai/gocryptotrader_types/exchanges/orderbook"
	"github.com/blep-ai/gocryptotrader_types/exchanges/ticker"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
	"github.com/blep-ai/gocryptotrader_types/portfolio/withdraw"
)

type IBotExchange interface {
	Setup(exch *config.ExchangeConfig) error
	Start(wg *sync.WaitGroup)
	SetDefaults()
	GetName() string
	IsEnabled() bool
	SetEnabled(bool)
	ValidateCredentials() error
	FetchTicker(p currency.Pair, a asset.Item) (*ticker.Price, error)
	UpdateTicker(p currency.Pair, a asset.Item) (*ticker.Price, error)
	FetchOrderbook(p currency.Pair, a asset.Item) (*orderbook.Base, error)
	UpdateOrderbook(p currency.Pair, a asset.Item) (*orderbook.Base, error)
	FetchTradablePairs(a asset.Item) ([]string, error)
	UpdateTradablePairs(forceUpdate bool) error
	GetEnabledPairs(a asset.Item) currency.Pairs
	GetAvailablePairs(a asset.Item) currency.Pairs
	FetchAccountInfo() (account.Holdings, error)
	UpdateAccountInfo() (account.Holdings, error)
	GetAuthenticatedAPISupport(endpoint uint8) bool
	SetPairs(pairs currency.Pairs, a asset.Item, enabled bool) error
	GetAssetTypes() asset.Items
	GetExchangeHistory(p currency.Pair, a asset.Item) ([]TradeHistory, error)
	SupportsAutoPairUpdates() bool
	SupportsRESTTickerBatchUpdates() bool
	GetFeeByType(f *FeeBuilder) (float64, error)
	GetLastPairsUpdateTime() int64
	GetWithdrawPermissions() uint32
	FormatWithdrawPermissions() string
	SupportsWithdrawPermissions(permissions uint32) bool
	GetFundingHistory() ([]FundHistory, error)
	SubmitOrder(s *order.Submit) (order.SubmitResponse, error)
	ModifyOrder(action *order.Modify) (string, error)
	CancelOrder(order *order.Cancel) error
	CancelAllOrders(orders *order.Cancel) (order.CancelAllResponse, error)
	GetOrderInfo(orderID string) (order.Detail, error)
	GetDepositAddress(cryptocurrency currency.Code, accountID string) (string, error)
	GetOrderHistory(getOrdersRequest *order.GetOrdersRequest) ([]order.Detail, error)
	GetActiveOrders(getOrdersRequest *order.GetOrdersRequest) ([]order.Detail, error)
	WithdrawCryptocurrencyFunds(withdrawRequest *withdraw.Request) (*withdraw.ExchangeResponse, error)
	WithdrawFiatFunds(withdrawRequest *withdraw.Request) (*withdraw.ExchangeResponse, error)
	WithdrawFiatFundsToInternationalBank(withdrawRequest *withdraw.Request) (*withdraw.ExchangeResponse, error)
	SetHTTPClientUserAgent(ua string)
	GetHTTPClientUserAgent() string
	SetClientProxyAddress(addr string) error
	SupportsWebsocket() bool
	SupportsREST() bool
	IsWebsocketEnabled() bool
	GetWebsocket() (*wshandler.Websocket, error)
	SubscribeToWebsocketChannels(channels []wshandler.WebsocketChannelSubscription) error
	UnsubscribeToWebsocketChannels(channels []wshandler.WebsocketChannelSubscription) error
	AuthenticateWebsocket() error
	GetSubscriptions() ([]wshandler.WebsocketChannelSubscription, error)
	GetDefaultConfig() (*config.ExchangeConfig, error)
	GetBase() *Base
	SupportsAsset(assetType asset.Item) bool
	GetHistoricCandles(p currency.Pair, a asset.Item, timeStart, timeEnd time.Time, interval time.Duration) (kline.Item, error)
	DisableRateLimiter() error
	EnableRateLimiter() error
}
