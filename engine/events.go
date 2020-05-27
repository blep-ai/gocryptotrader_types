package engine
import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/blep-ai/gocryptotrader_types/communications/base"
	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/blep-ai/gocryptotrader_types/exchanges/orderbook"
	"github.com/blep-ai/gocryptotrader_types/exchanges/ticker"
	"github.com/blep-ai/gocryptotrader_types/log"
)
type EventConditionParams struct {
	Condition string
	Price     float64

	CheckBids        bool
	CheckBidsAndAsks bool
	OrderbookAmount  float64
}
type Event struct {
	ID        int64
	Exchange  string
	Item      string
	Condition EventConditionParams
	Pair      currency.Pair
	Asset     asset.Item
	Action    string
	Executed  bool
}
