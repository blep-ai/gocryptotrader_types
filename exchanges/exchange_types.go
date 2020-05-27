package exchange

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/protocol"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)

type FeeBuilder struct {
	FeeType FeeType
	// Used for calculating crypto trading fees, deposits & withdrawals
	Pair    currency.Pair
	IsMaker bool
	// Fiat currency used for bank deposits & withdrawals
	FiatCurrency        currency.Code
	BankTransactionType InternationalBankTransactionType
	// Used to multiply for fee calculations
	PurchasePrice float64
	Amount        float64
}
type TradeHistory struct {
	Timestamp   time.Time
	TID         string
	Price       float64
	Amount      float64
	Exchange    string
	Type        string
	Fee         float64
	Description string
}
type FundHistory struct {
	ExchangeName      string
	Status            string
	TransferID        string
	Description       string
	Timestamp         time.Time
	Currency          string
	Amount            float64
	Fee               float64
	TransferType      string
	CryptoToAddress   string
	CryptoFromAddress string
	CryptoTxID        string
	BankTo            string
	BankFrom          string
}
type Features struct {
	Supports FeaturesSupported
	Enabled  FeaturesEnabled
}
type FeaturesEnabled struct {
	AutoPairUpdates bool
}
type FeaturesSupported struct {
	REST                  bool
	RESTCapabilities      protocol.Features
	Websocket             bool
	WebsocketCapabilities protocol.Features
	WithdrawPermissions   uint32
}
type API struct {
	AuthenticatedSupport          bool
	AuthenticatedWebsocketSupport bool
	PEMKeySupport                 bool

	Endpoints struct {
		URL                 string
		URLDefault          string
		URLSecondary        string
		URLSecondaryDefault string
		WebsocketURL        string
	}

	Credentials struct {
		Key      string
		Secret   string
		ClientID string
		PEMKey   string
	}

	CredentialsValidator struct {
		// For Huobi (optional)
		RequiresPEM bool

		RequiresKey                bool
		RequiresSecret             bool
		RequiresClientID           bool
		RequiresBase64DecodeSecret bool
	}
}
type Base struct {
	Name                          string
	Enabled                       bool
	Verbose                       bool
	LoadedByConfig                bool
	SkipAuthCheck                 bool
	API                           API
	BaseCurrencies                currency.Currencies
	CurrencyPairs                 currency.PairsManager
	Features                      Features
	HTTPTimeout                   time.Duration
	HTTPUserAgent                 string
	HTTPRecording                 bool
	HTTPDebugging                 bool
	WebsocketResponseCheckTimeout time.Duration
	WebsocketResponseMaxLimit     time.Duration
	WebsocketOrderbookBufferLimit int64
	Websocket                     *wshandler.Websocket
	*request.Requester
	Config *config.ExchangeConfig
}
type FeeType uint8
type InternationalBankTransactionType uint8
