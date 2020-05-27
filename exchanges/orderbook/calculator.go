package orderbook

import (
	"errors"
	"fmt"
	"sort"

	math "github.com/blep-ai/gocryptotrader_types/common/math"
	"github.com/blep-ai/gocryptotrader_types/log"
)
type WhaleBombResult struct {
	Amount               float64
	MinimumPrice         float64
	MaximumPrice         float64
	PercentageGainOrLoss float64
	Orders               orderSummary
	Status               string
}
