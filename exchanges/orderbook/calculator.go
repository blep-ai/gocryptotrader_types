package orderbook

import (
	"errors"
	"fmt"
	"sort"

	math "github.com/thrasher-corp/gocryptotrader/common/math"
	"github.com/thrasher-corp/gocryptotrader/log"
)
type WhaleBombResult struct {
	Amount               float64
	MinimumPrice         float64
	MaximumPrice         float64
	PercentageGainOrLoss float64
	Orders               orderSummary
	Status               string
}
