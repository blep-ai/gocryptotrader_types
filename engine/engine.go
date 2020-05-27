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

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/currency/coinmarketcap"
	"github.com/blep-ai/gocryptotrader_types/dispatch"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	gctscript "github.com/blep-ai/gocryptotrader_types/gctscript/vm"
	gctlog "github.com/blep-ai/gocryptotrader_types/log"
	"github.com/blep-ai/gocryptotrader_types/portfolio"
	"github.com/blep-ai/gocryptotrader_types/portfolio/withdraw"
	"github.com/blep-ai/gocryptotrader_types/utils"
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
