package config

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/database"
	"github.com/blep-ai/gocryptotrader_types/exchanges/protocol"
	gctscript "github.com/blep-ai/gocryptotrader_types/gctscript/vm"
	"github.com/blep-ai/gocryptotrader_types/log"
	"github.com/blep-ai/gocryptotrader_types/portfolio"
	"github.com/blep-ai/gocryptotrader_types/portfolio/banking"
)

type Config struct {
	Name              string                  `json:"name"`
	EncryptConfig     int                     `json:"encryptConfig"`
	GlobalHTTPTimeout time.Duration           `json:"globalHTTPTimeout"`
	Database          database.Config         `json:"database"`
	Logging           log.Config              `json:"logging"`
	ConnectionMonitor ConnectionMonitorConfig `json:"connectionMonitor"`
	Profiler          Profiler                `json:"profiler"`
	NTPClient         NTPClientConfig         `json:"ntpclient"`
	GCTScript         gctscript.Config        `json:"gctscript"`
	Currency          CurrencyConfig          `json:"currencyConfig"`
	Communications    CommunicationsConfig    `json:"communications"`
	RemoteControl     RemoteControlConfig     `json:"remoteControl"`
	Portfolio         portfolio.Base          `json:"portfolioAddresses"`
	Exchanges         []ExchangeConfig        `json:"exchanges"`
	BankAccounts      []banking.Account       `json:"bankAccounts"`

	// Deprecated config settings, will be removed at a future date
	Webserver           *WebserverConfig          `json:"webserver,omitempty"`
	CurrencyPairFormat  *CurrencyPairFormatConfig `json:"currencyPairFormat,omitempty"`
	FiatDisplayCurrency *currency.Code            `json:"fiatDispayCurrency,omitempty"`
	Cryptocurrencies    *currency.Currencies      `json:"cryptocurrencies,omitempty"`
	SMS                 *SMSGlobalConfig          `json:"smsGlobal,omitempty"`
}
type ConnectionMonitorConfig struct {
	DNSList          []string      `json:"preferredDNSList"`
	PublicDomainList []string      `json:"preferredDomainList"`
	CheckInterval    time.Duration `json:"checkInterval"`
}
type ExchangeConfig struct {
	Name                          string                 `json:"name"`
	Enabled                       bool                   `json:"enabled"`
	Verbose                       bool                   `json:"verbose"`
	UseSandbox                    bool                   `json:"useSandbox,omitempty"`
	HTTPTimeout                   time.Duration          `json:"httpTimeout"`
	HTTPUserAgent                 string                 `json:"httpUserAgent,omitempty"`
	HTTPDebugging                 bool                   `json:"httpDebugging,omitempty"`
	WebsocketResponseCheckTimeout time.Duration          `json:"websocketResponseCheckTimeout"`
	WebsocketResponseMaxLimit     time.Duration          `json:"websocketResponseMaxLimit"`
	WebsocketTrafficTimeout       time.Duration          `json:"websocketTrafficTimeout"`
	WebsocketOrderbookBufferLimit int                    `json:"websocketOrderbookBufferLimit"`
	ProxyAddress                  string                 `json:"proxyAddress,omitempty"`
	BaseCurrencies                currency.Currencies    `json:"baseCurrencies"`
	CurrencyPairs                 *currency.PairsManager `json:"currencyPairs"`
	API                           APIConfig              `json:"api"`
	Features                      *FeaturesConfig        `json:"features"`
	BankAccounts                  []banking.Account      `json:"bankAccounts,omitempty"`

	// Deprecated settings which will be removed in a future update
	AvailablePairs                   *currency.Pairs      `json:"availablePairs,omitempty"`
	EnabledPairs                     *currency.Pairs      `json:"enabledPairs,omitempty"`
	AssetTypes                       *string              `json:"assetTypes,omitempty"`
	PairsLastUpdated                 *int64               `json:"pairsLastUpdated,omitempty"`
	ConfigCurrencyPairFormat         *currency.PairFormat `json:"configCurrencyPairFormat,omitempty"`
	RequestCurrencyPairFormat        *currency.PairFormat `json:"requestCurrencyPairFormat,omitempty"`
	AuthenticatedAPISupport          *bool                `json:"authenticatedApiSupport,omitempty"`
	AuthenticatedWebsocketAPISupport *bool                `json:"authenticatedWebsocketApiSupport,omitempty"`
	APIKey                           *string              `json:"apiKey,omitempty"`
	APISecret                        *string              `json:"apiSecret,omitempty"`
	APIAuthPEMKeySupport             *bool                `json:"apiAuthPemKeySupport,omitempty"`
	APIAuthPEMKey                    *string              `json:"apiAuthPemKey,omitempty"`
	APIURL                           *string              `json:"apiUrl,omitempty"`
	APIURLSecondary                  *string              `json:"apiUrlSecondary,omitempty"`
	ClientID                         *string              `json:"clientId,omitempty"`
	SupportsAutoPairUpdates          *bool                `json:"supportsAutoPairUpdates,omitempty"`
	Websocket                        *bool                `json:"websocket,omitempty"`
	WebsocketURL                     *string              `json:"websocketUrl,omitempty"`
}
type Profiler struct {
	Enabled              bool `json:"enabled"`
	MutexProfileFraction int  `json:"mutex_profile_fraction"`
}
type NTPClientConfig struct {
	Level                     int            `json:"enabled"`
	Pool                      []string       `json:"pool"`
	AllowedDifference         *time.Duration `json:"allowedDifference"`
	AllowedNegativeDifference *time.Duration `json:"allowedNegativeDifference"`
}
type GRPCConfig struct {
	Enabled                bool   `json:"enabled"`
	ListenAddress          string `json:"listenAddress"`
	GRPCProxyEnabled       bool   `json:"grpcProxyEnabled"`
	GRPCProxyListenAddress string `json:"grpcProxyListenAddress"`
}
type DepcrecatedRPCConfig struct {
	Enabled       bool   `json:"enabled"`
	ListenAddress string `json:"listenAddress"`
}
type WebsocketRPCConfig struct {
	Enabled             bool   `json:"enabled"`
	ListenAddress       string `json:"listenAddress"`
	ConnectionLimit     int    `json:"connectionLimit"`
	MaxAuthFailures     int    `json:"maxAuthFailures"`
	AllowInsecureOrigin bool   `json:"allowInsecureOrigin"`
}
type RemoteControlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`

	GRPC          GRPCConfig           `json:"gRPC"`
	DeprecatedRPC DepcrecatedRPCConfig `json:"deprecatedRPC"`
	WebsocketRPC  WebsocketRPCConfig   `json:"websocketRPC"`
}
type WebserverConfig struct {
	Enabled                      bool   `json:"enabled"`
	AdminUsername                string `json:"adminUsername"`
	AdminPassword                string `json:"adminPassword"`
	ListenAddress                string `json:"listenAddress"`
	WebsocketConnectionLimit     int    `json:"websocketConnectionLimit"`
	WebsocketMaxAuthFailures     int    `json:"websocketMaxAuthFailures"`
	WebsocketAllowInsecureOrigin bool   `json:"websocketAllowInsecureOrigin"`
}
type Post struct {
	Data Config `json:"data"`
}
type CurrencyPairFormatConfig struct {
	Uppercase bool   `json:"uppercase"`
	Delimiter string `json:"delimiter,omitempty"`
	Separator string `json:"separator,omitempty"`
	Index     string `json:"index,omitempty"`
}
type BankTransaction struct {
	ReferenceNumber     string `json:"referenceNumber"`
	TransactionNumber   string `json:"transactionNumber"`
	PaymentInstructions string `json:"paymentInstructions"`
}
type CurrencyConfig struct {
	ForexProviders                []currency.FXSettings     `json:"forexProviders"`
	CryptocurrencyProvider        CryptocurrencyProvider    `json:"cryptocurrencyProvider"`
	Cryptocurrencies              currency.Currencies       `json:"cryptocurrencies"`
	CurrencyPairFormat            *CurrencyPairFormatConfig `json:"currencyPairFormat"`
	FiatDisplayCurrency           currency.Code             `json:"fiatDisplayCurrency"`
	CurrencyFileUpdateDuration    time.Duration             `json:"currencyFileUpdateDuration"`
	ForeignExchangeUpdateDuration time.Duration             `json:"foreignExchangeUpdateDuration"`
}
type CryptocurrencyProvider struct {
	Name        string `json:"name"`
	Enabled     bool   `json:"enabled"`
	Verbose     bool   `json:"verbose"`
	APIkey      string `json:"apiKey"`
	AccountPlan string `json:"accountPlan"`
}
type CommunicationsConfig struct {
	SlackConfig     SlackConfig     `json:"slack"`
	SMSGlobalConfig SMSGlobalConfig `json:"smsGlobal"`
	SMTPConfig      SMTPConfig      `json:"smtp"`
	TelegramConfig  TelegramConfig  `json:"telegram"`
}
type SlackConfig struct {
	Name              string `json:"name"`
	Enabled           bool   `json:"enabled"`
	Verbose           bool   `json:"verbose"`
	TargetChannel     string `json:"targetChannel"`
	VerificationToken string `json:"verificationToken"`
}
type SMSContact struct {
	Name    string `json:"name"`
	Number  string `json:"number"`
	Enabled bool   `json:"enabled"`
}
type SMSGlobalConfig struct {
	Name     string       `json:"name"`
	From     string       `json:"from"`
	Enabled  bool         `json:"enabled"`
	Verbose  bool         `json:"verbose"`
	Username string       `json:"username"`
	Password string       `json:"password"`
	Contacts []SMSContact `json:"contacts"`
}
type SMTPConfig struct {
	Name            string `json:"name"`
	Enabled         bool   `json:"enabled"`
	Verbose         bool   `json:"verbose"`
	Host            string `json:"host"`
	Port            string `json:"port"`
	AccountName     string `json:"accountName"`
	AccountPassword string `json:"accountPassword"`
	From            string `json:"from"`
	RecipientList   string `json:"recipientList"`
}
type TelegramConfig struct {
	Name              string `json:"name"`
	Enabled           bool   `json:"enabled"`
	Verbose           bool   `json:"verbose"`
	VerificationToken string `json:"verificationToken"`
}
type FeaturesSupportedConfig struct {
	REST                  bool              `json:"restAPI"`
	RESTCapabilities      protocol.Features `json:"restCapabilities,omitempty"`
	Websocket             bool              `json:"websocketAPI"`
	WebsocketCapabilities protocol.Features `json:"websocketCapabilities,omitempty"`
}
type FeaturesEnabledConfig struct {
	AutoPairUpdates bool `json:"autoPairUpdates"`
	Websocket       bool `json:"websocketAPI"`
}
type FeaturesConfig struct {
	Supports FeaturesSupportedConfig `json:"supports"`
	Enabled  FeaturesEnabledConfig   `json:"enabled"`
}
type APIEndpointsConfig struct {
	URL          string `json:"url"`
	URLSecondary string `json:"urlSecondary"`
	WebsocketURL string `json:"websocketURL"`
}
type APICredentialsConfig struct {
	Key       string `json:"key,omitempty"`
	Secret    string `json:"secret,omitempty"`
	ClientID  string `json:"clientID,omitempty"`
	PEMKey    string `json:"pemKey,omitempty"`
	OTPSecret string `json:"otpSecret,omitempty"`
}
type APICredentialsValidatorConfig struct {
	// For Huobi (optional)
	RequiresPEM bool `json:"requiresPEM,omitempty"`

	RequiresKey                bool `json:"requiresKey,omitempty"`
	RequiresSecret             bool `json:"requiresSecret,omitempty"`
	RequiresClientID           bool `json:"requiresClientID,omitempty"`
	RequiresBase64DecodeSecret bool `json:"requiresBase64DecodeSecret,omitempty"`
}
type APIConfig struct {
	AuthenticatedSupport          bool `json:"authenticatedSupport"`
	AuthenticatedWebsocketSupport bool `json:"authenticatedWebsocketApiSupport"`
	PEMKeySupport                 bool `json:"pemKeySupport,omitempty"`

	Endpoints            APIEndpointsConfig             `json:"endpoints"`
	Credentials          APICredentialsConfig           `json:"credentials"`
	CredentialsValidator *APICredentialsValidatorConfig `json:"credentialsValidator,omitempty"`
}
