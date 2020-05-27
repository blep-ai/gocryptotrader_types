package gct

import (
	"fmt"
	"strings"
	"time"

	objects "github.com/d5/tengo/v2"
	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
	"github.com/blep-ai/gocryptotrader_types/gctscript/modules/ta/indicators"
	"github.com/blep-ai/gocryptotrader_types/gctscript/wrappers"
	"github.com/blep-ai/gocryptotrader_types/portfolio/withdraw"
)
type OHLCV struct {
	objects.Map
}
