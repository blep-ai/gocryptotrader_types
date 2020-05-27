package coinmarketcap

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
)

type Coinmarketcap struct {
	Verbose    bool
	Enabled    bool
	Name       string
	APIkey     string
	APIUrl     string
	APIVersion string
	Plan       uint8
	Requester  *request.Requester
}
type Settings struct {
	Name        string `json:"name"`
	Enabled     bool   `json:"enabled"`
	Verbose     bool   `json:"verbose"`
	APIkey      string `json:"apiKey"`
	AccountPlan string `json:"accountPlan"`
}
type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int64  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int64  `json:"elapsed"`
	CreditCount  int64  `json:"credit_count"`
}
type Currency struct {
	Price                  float64   `json:"price"`
	Volume24H              float64   `json:"volume_24h"`
	Volume24HAdjusted      float64   `json:"volume_24h_adjusted"`
	Volume7D               float64   `json:"volume_7d"`
	Volume30D              float64   `json:"volume_30d"`
	PercentChange1H        float64   `json:"percent_change_1h"`
	PercentChangeVolume24H float64   `json:"percent_change_volume_24h"`
	PercentChangeVolume7D  float64   `json:"percent_change_volume_7d"`
	PercentChangeVolume30D float64   `json:"percent_change_volume_30d"`
	MarketCap              float64   `json:"market_cap"`
	TotalMarketCap         float64   `json:"total_market_cap"`
	LastUpdated            time.Time `json:"last_updated"`
}
type OHLC struct {
	Open      float64   `json:"open"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Close     float64   `json:"close"`
	Volume    float64   `json:"volume"`
	Timestamp time.Time `json:"timestamp"`
}
type CryptoCurrencyMap struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Symbol              string    `json:"symbol"`
	Slug                string    `json:"slug"`
	IsActive            int       `json:"is_active"`
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData  time.Time `json:"last_historical_data"`
	Platform            struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		Symbol       string `json:"symbol"`
		Slug         string `json:"slug"`
		TokenAddress string `json:"token_address"`
	} `json:"platform"`
}
type CryptocurrencyHistoricalListings struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	Slug              string    `json:"slug"`
	CmcRank           int       `json:"cmc_rank"`
	NumMarketPairs    int       `json:"num_market_pairs"`
	CirculatingSupply float64   `json:"circulating_supply"`
	TotalSupply       float64   `json:"total_supply"`
	MaxSupply         float64   `json:"max_supply"`
	LastUpdated       time.Time `json:"last_updated"`
	Quote             struct {
		USD Currency `json:"USD"`
		BTC Currency `json:"BTC"`
	} `json:"quote"`
}
type CryptocurrencyLatestListings struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	CmcRank           int         `json:"cmc_rank"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	CirculatingSupply float64     `json:"circulating_supply"`
	TotalSupply       float64     `json:"total_supply"`
	MaxSupply         float64     `json:"max_supply"`
	LastUpdated       time.Time   `json:"last_updated"`
	DateAdded         time.Time   `json:"date_added"`
	Tags              []string    `json:"tags"`
	Platform          interface{} `json:"platform"`
	Quote             struct {
		USD Currency `json:"USD"`
		BTC Currency `json:"BTC"`
	} `json:"quote"`
}
type CryptocurrencyLatestMarketPairs struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	NumMarketPairs int    `json:"num_market_pairs"`
	MarketPairs    []struct {
		Exchange struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"exchange"`
		MarketPair     string `json:"market_pair"`
		MarketPairBase struct {
			CurrencyID     int    `json:"currency_id"`
			CurrencySymbol string `json:"currency_symbol"`
			CurrencyType   string `json:"currency_type"`
		} `json:"market_pair_base"`
		MarketPairQuote struct {
			CurrencyID     int    `json:"currency_id"`
			CurrencySymbol string `json:"currency_symbol"`
			CurrencyType   string `json:"currency_type"`
		} `json:"market_pair_quote"`
		Quote struct {
			ExchangeReported struct {
				Price          float64   `json:"price"`
				Volume24HBase  float64   `json:"volume_24h_base"`
				Volume24HQuote float64   `json:"volume_24h_quote"`
				LastUpdated    time.Time `json:"last_updated"`
			} `json:"exchange_reported"`
			USD Currency `json:"USD"`
		} `json:"quote"`
	} `json:"market_pairs"`
}
type CryptocurrencyOHLCHistorical struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quotes []struct {
		TimeOpen  time.Time `json:"time_open"`
		TimeClose time.Time `json:"time_close"`
		Quote     struct {
			USD OHLC `json:"USD"`
		} `json:"quote"`
	} `json:"quotes"`
}
type CryptocurrencyHistoricalQuotes struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quotes []struct {
		Timestamp time.Time `json:"timestamp"`
		Quote     struct {
			USD Currency `json:"USD"`
		} `json:"quote"`
	} `json:"quotes"`
}
type ExchangeMap struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Slug                string    `json:"slug"`
	IsActive            int       `json:"is_active"`
	FirstHistoricalData time.Time `json:"first_historical_data"`
	LastHistoricalData  time.Time `json:"last_historical_data"`
}
type ExchangeHistoricalListings struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	CmcRank        int       `json:"cmc_rank"`
	NumMarketPairs int       `json:"num_market_pairs"`
	Timestamp      time.Time `json:"timestamp"`
	Quote          struct {
		USD Currency `json:"USD"`
	} `json:"quote"`
}
type ExchangeLatestListings struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	NumMarketPairs int       `json:"num_market_pairs"`
	LastUpdated    time.Time `json:"last_updated"`
	Quote          struct {
		USD Currency `json:"USD"`
	} `json:"quote"`
}
type ExchangeLatestMarketPairs struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	NumMarketPairs int    `json:"num_market_pairs"`
	MarketPairs    []struct {
		MarketPair     string `json:"market_pair"`
		MarketPairBase struct {
			CurrencyID     int    `json:"currency_id"`
			CurrencySymbol string `json:"currency_symbol"`
			CurrencyType   string `json:"currency_type"`
		} `json:"market_pair_base"`
		MarketPairQuote struct {
			CurrencyID     int    `json:"currency_id"`
			CurrencySymbol string `json:"currency_symbol"`
			CurrencyType   string `json:"currency_type"`
		} `json:"market_pair_quote"`
		Quote struct {
			ExchangeReported struct {
				Price          float64   `json:"price"`
				Volume24HBase  float64   `json:"volume_24h_base"`
				Volume24HQuote float64   `json:"volume_24h_quote"`
				LastUpdated    time.Time `json:"last_updated"`
			} `json:"exchange_reported"`
			USD Currency `json:"USD"`
		} `json:"quote"`
	} `json:"market_pairs"`
}
type ExchangeLatestQuotes struct {
	Binance struct {
		ID             int       `json:"id"`
		Name           string    `json:"name"`
		Slug           string    `json:"slug"`
		NumMarketPairs int       `json:"num_market_pairs"`
		LastUpdated    time.Time `json:"last_updated"`
		Quote          struct {
			USD Currency `json:"USD"`
		} `json:"quote"`
	} `json:"binance"`
}
type ExchangeHistoricalQuotes struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Quotes []struct {
		Timestamp time.Time `json:"timestamp"`
		Quote     struct {
			USD Currency `json:"USD"`
		} `json:"quote"`
		NumMarketPairs int `json:"num_market_pairs"`
	} `json:"quotes"`
}
type GlobalMeticLatestQuotes struct {
	BtcDominance           float64   `json:"btc_dominance"`
	EthDominance           float64   `json:"eth_dominance"`
	ActiveCryptocurrencies int       `json:"active_cryptocurrencies"`
	ActiveMarketPairs      int       `json:"active_market_pairs"`
	ActiveExchanges        int       `json:"active_exchanges"`
	LastUpdated            time.Time `json:"last_updated"`
	Quote                  struct {
		USD Currency `json:"USD"`
	} `json:"quote"`
}
type GlobalMeticHistoricalQuotes struct {
	Quotes []struct {
		Timestamp    time.Time `json:"timestamp"`
		BtcDominance float64   `json:"btc_dominance"`
		Quote        struct {
			USD Currency `json:"USD"`
		} `json:"quote"`
	} `json:"quotes"`
}
type PriceConversion struct {
	Symbol      string    `json:"symbol"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	LastUpdated time.Time `json:"last_updated"`
	Quote       struct {
		GBP Currency `json:"GBP"`
		LTC Currency `json:"LTC"`
		USD Currency `json:"USD"`
	} `json:"quote"`
}
