package stats

import (
	"sort"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
)
type Item struct {
	Exchange  string
	Pair      currency.Pair
	AssetType asset.Item
	Price     float64
	Volume    float64
}
