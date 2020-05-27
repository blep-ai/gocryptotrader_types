package engine

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/currency/coinmarketcap"
	"github.com/thrasher-corp/gocryptotrader/dispatch"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	gctscript "github.com/thrasher-corp/gocryptotrader/gctscript/vm"
	gctlog "github.com/thrasher-corp/gocryptotrader/log"
	"github.com/thrasher-corp/gocryptotrader/portfolio"
	"github.com/thrasher-corp/gocryptotrader/portfolio/withdraw"
	"github.com/thrasher-corp/gocryptotrader/utils"
)
type Engine struct {
	Config                      *config.Config
	Portfolio                   *portfolio.Base
	ExchangeCurrencyPairManager *ExchangeCurrencyPairSyncer
	NTPManager                  ntpManager
	ConnectionManager           connectionManager
	DatabaseManager             databaseManager
	GctScriptManager            gctScriptManager
	OrderManager                orderManager
	PortfolioManager            portfolioManager
	CommsManager                commsManager
	exchangeManager             exchangeManager
	DepositAddressManager       *DepositAddressManager
	Settings                    Settings
	Uptime                      time.Time
	ServicesWG                  sync.WaitGroup
}
