package orderbook

import (
	"sync"
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/dispatch"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/gofrs/uuid"
)

type Book struct {
	b     *Base
	Main  uuid.UUID
	Assoc []uuid.UUID
}
type Service struct {
	Books    map[string]map[*currency.Item]map[*currency.Item]map[asset.Item]*Book
	Exchange map[string]uuid.UUID
	mux      *dispatch.Mux
	sync.RWMutex
}
type Item struct {
	Amount float64
	Price  float64
	ID     int64

	// Contract variables
	LiquidationOrders int64
	OrderCount        int64
}
type Base struct {
	Pair         currency.Pair `json:"pair"`
	Bids         []Item        `json:"bids"`
	Asks         []Item        `json:"asks"`
	LastUpdated  time.Time     `json:"lastUpdated"`
	LastUpdateID int64         `json:"lastUpdateId"`
	AssetType    asset.Item    `json:"assetType"`
	ExchangeName string        `json:"exchangeName"`
}
type byOBPrice []Item
