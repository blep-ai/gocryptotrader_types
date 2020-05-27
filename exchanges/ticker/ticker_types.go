package ticker
import (
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/dispatch"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
)
type Service struct {
	Tickers  map[string]map[*currency.Item]map[*currency.Item]map[asset.Item]*Ticker
	Exchange map[string]uuid.UUID
	mux      *dispatch.Mux
	sync.RWMutex
}
type Price struct {
	Last         float64       `json:"Last"`
	High         float64       `json:"High"`
	Low          float64       `json:"Low"`
	Bid          float64       `json:"Bid"`
	Ask          float64       `json:"Ask"`
	Volume       float64       `json:"Volume"`
	QuoteVolume  float64       `json:"QuoteVolume"`
	PriceATH     float64       `json:"PriceATH"`
	Open         float64       `json:"Open"`
	Close        float64       `json:"Close"`
	Pair         currency.Pair `json:"Pair"`
	ExchangeName string        `json:"exchangeName"`
	AssetType    asset.Item    `json:"assetType"`
	LastUpdated  time.Time
}
type Ticker struct {
	Price
	Main  uuid.UUID
	Assoc []uuid.UUID
}
