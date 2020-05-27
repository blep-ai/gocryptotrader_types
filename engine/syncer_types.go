package engine

import (
	"sync"
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
)
type CurrencyPairSyncerConfig struct {
	SyncTicker       bool
	SyncOrderbook    bool
	SyncTrades       bool
	SyncContinuously bool
	SyncTimeout      time.Duration
	NumWorkers       int
	Verbose          bool
}
type ExchangeSyncerConfig struct {
	SyncDepositAddresses bool
	SyncOrders           bool
}
type ExchangeCurrencyPairSyncer struct {
	Cfg                      CurrencyPairSyncerConfig
	CurrencyPairs            []CurrencyPairSyncAgent
	tickerBatchLastRequested map[string]time.Time
	mux                      sync.Mutex
	initSyncWG               sync.WaitGroup

	initSyncCompleted int32
	initSyncStarted   int32
	initSyncStartTime time.Time
	shutdown          int32
}
type SyncBase struct {
	IsUsingWebsocket bool
	IsUsingREST      bool
	IsProcessing     bool
	LastUpdated      time.Time
	HaveData         bool
	NumErrors        int
}
type CurrencyPairSyncAgent struct {
	Created   time.Time
	Exchange  string
	AssetType asset.Item
	Pair      currency.Pair
	Ticker    SyncBase
	Orderbook SyncBase
	Trade     SyncBase
}
