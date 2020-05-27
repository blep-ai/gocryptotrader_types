package currency

import (
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency/coinmarketcap"
)
type MainConfiguration struct {
	ForexProviders         []FXSettings
	CryptocurrencyProvider coinmarketcap.Settings
	Cryptocurrencies       Currencies
	CurrencyPairFormat     interface{}
	FiatDisplayCurrency    Code
	CurrencyDelay          time.Duration
	FxRateDelay            time.Duration
}
type BotOverrides struct {
	Coinmarketcap       bool
	FxCurrencyConverter bool
	FxCurrencyLayer     bool
	FxFixer             bool
	FxOpenExchangeRates bool
}
type SystemsSettings struct {
	Coinmarketcap     coinmarketcap.Settings
	Currencyconverter FXSettings
	Currencylayer     FXSettings
	Fixer             FXSettings
	Openexchangerates FXSettings
}
type FXSettings struct {
	Name             string        `json:"name"`
	Enabled          bool          `json:"enabled"`
	Verbose          bool          `json:"verbose"`
	RESTPollingDelay time.Duration `json:"restPollingDelay"`
	APIKey           string        `json:"apiKey"`
	APIKeyLvl        int           `json:"apiKeyLvl"`
	PrimaryProvider  bool          `json:"primaryProvider"`
}
type File struct {
	LastMainUpdate time.Time `json:"lastMainUpdate"`
	Cryptocurrency []Item    `json:"cryptocurrencies"`
	FiatCurrency   []Item    `json:"fiatCurrencies"`
	UnsetCurrency  []Item    `json:"unsetCurrencies"`
	Contracts      []Item    `json:"contracts"`
	Token          []Item    `json:"tokens"`
}
