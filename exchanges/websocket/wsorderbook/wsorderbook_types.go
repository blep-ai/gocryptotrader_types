package wsorderbook
import (
	"sync"
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/blep-ai/gocryptotrader_types/exchanges/orderbook"
)
type WebsocketOrderbookLocal struct {
	ob                    map[currency.Pair]map[asset.Item]*orderbook.Base
	buffer                map[currency.Pair]map[asset.Item][]*WebsocketOrderbookUpdate
	obBufferLimit         int
	bufferEnabled         bool
	sortBuffer            bool
	sortBufferByUpdateIDs bool // When timestamps aren't provided, an id can help sort
	updateEntriesByID     bool // Use the update IDs to match ob entries
	exchangeName          string
	m                     sync.Mutex
}
type WebsocketOrderbookUpdate struct {
	UpdateID   int64 // Used when no time is provided
	UpdateTime time.Time
	Asset      asset.Item
	Action     string // Used in conjunction with UpdateEntriesByID
	Bids       []orderbook.Item
	Asks       []orderbook.Item
	Pair       currency.Pair
}
