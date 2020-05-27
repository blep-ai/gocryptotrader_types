package gct

import (
	"fmt"
	"strings"
	"time"

	objects "github.com/d5/tengo/v2"
	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
	"github.com/thrasher-corp/gocryptotrader/gctscript/modules/ta/indicators"
	"github.com/thrasher-corp/gocryptotrader/gctscript/wrappers"
	"github.com/thrasher-corp/gocryptotrader/portfolio/withdraw"
)
type OHLCV struct {
	objects.Map
}
