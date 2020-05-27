package engine

import (
	"sync"
	"time"

	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/portfolio"
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
