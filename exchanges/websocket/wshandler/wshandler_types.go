package wshandler

import (
	"sync"
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
	"github.com/blep-ai/gocryptotrader_types/exchanges/protocol"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wsorderbook"
	"github.com/gorilla/websocket"
)

type Websocket struct {
	canUseAuthenticatedEndpoints bool
	enabled                      bool
	init                         bool
	connected                    bool
	connecting                   bool
	trafficMonitorRunning        bool
	verbose                      bool
	connectionMonitorRunning     bool
	trafficTimeout               time.Duration
	proxyAddr                    string
	defaultURL                   string
	runningURL                   string
	exchangeName                 string
	m                            sync.Mutex
	subscriptionMutex            sync.Mutex
	connectionMutex              sync.RWMutex
	connector                    func() error
	subscribedChannels           []WebsocketChannelSubscription
	channelsToSubscribe          []WebsocketChannelSubscription
	channelSubscriber            func(channelToSubscribe WebsocketChannelSubscription) error
	channelUnsubscriber          func(channelToUnsubscribe WebsocketChannelSubscription) error
	DataHandler                  chan interface{}
	// ShutdownC is the main shutdown channel which controls all websocket go funcs
	ShutdownC chan struct{}
	// Orderbook is a local cache of orderbooks
	Orderbook wsorderbook.WebsocketOrderbookLocal
	// Wg defines a wait group for websocket routines for cleanly shutting down
	// routines
	Wg sync.WaitGroup
	// TrafficAlert monitors if there is a halt in traffic throughput
	TrafficAlert chan struct{}
	// ReadMessageErrors will received all errors from ws.ReadMessage() and verify if its a disconnection
	ReadMessageErrors chan error
	features          *protocol.Features
}
type WebsocketSetup struct {
	Enabled                          bool
	Verbose                          bool
	AuthenticatedWebsocketAPISupport bool
	WebsocketTimeout                 time.Duration
	DefaultURL                       string
	ExchangeName                     string
	RunningURL                       string
	Connector                        func() error
	Subscriber                       func(channelToSubscribe WebsocketChannelSubscription) error
	UnSubscriber                     func(channelToUnsubscribe WebsocketChannelSubscription) error
	Features                         *protocol.Features
}
type WebsocketChannelSubscription struct {
	Channel  string
	Currency currency.Pair
	Params   map[string]interface{}
}
type WebsocketResponse struct {
	Type int
	Raw  []byte
}
type WebsocketOrderbookUpdate struct {
	Pair     currency.Pair
	Asset    asset.Item
	Exchange string
}
type TradeData struct {
	Timestamp    time.Time
	CurrencyPair currency.Pair
	AssetType    asset.Item
	Exchange     string
	EventType    order.Type
	Price        float64
	Amount       float64
	Side         order.Side
}
type FundingData struct {
	Timestamp    time.Time
	CurrencyPair currency.Pair
	AssetType    asset.Item
	Exchange     string
	Amount       float64
	Rate         float64
	Period       int64
	Side         order.Side
}
type KlineData struct {
	Timestamp  time.Time
	Pair       currency.Pair
	AssetType  asset.Item
	Exchange   string
	StartTime  time.Time
	CloseTime  time.Time
	Interval   string
	OpenPrice  float64
	ClosePrice float64
	HighPrice  float64
	LowPrice   float64
	Volume     float64
}
type WebsocketPositionUpdated struct {
	Timestamp time.Time
	Pair      currency.Pair
	AssetType asset.Item
	Exchange  string
}
type WebsocketConnection struct {
	sync.Mutex
	Verbose         bool
	connected       bool
	connectionMutex sync.RWMutex
	RateLimit       float64
	ExchangeName    string
	URL             string
	ProxyURL        string
	Wg              sync.WaitGroup
	Connection      *websocket.Conn
	Shutdown        chan struct{}
	// These are the request IDs and the corresponding response JSON
	IDResponses          map[int64][]byte
	ResponseCheckTimeout time.Duration
	ResponseMaxLimit     time.Duration
	TrafficTimeout       time.Duration
}
type WebsocketPingHandler struct {
	UseGorillaHandler bool
	MessageType       int
	Message           []byte
	Delay             time.Duration
}
type UnhandledMessageWarning struct {
	Message string
}
