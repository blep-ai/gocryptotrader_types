package main
import (
	"encoding/json"

	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
)
type Config struct {
	OrderSubmission OrderSubmission                         `json:"orderSubmission"`
	WalletAddress   string                                  `json:"withdrawWalletAddress"`
	BankDetails     Bank                                    `json:"bankAccount"`
	Exchanges       map[string]*config.APICredentialsConfig `json:"exchanges"`
}
type Key struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret,omitempty"`
	ClientID  string `json:"clientId,omitempty"`
	OTPSecret string `json:"otpSecret,omitempty"`
}
type ExchangeResponses struct {
	ID                 string
	ExchangeName       string                       `json:"exchangeName"`
	AssetPairResponses []ExchangeAssetPairResponses `json:"responses"`
	ErrorCount         int64                        `json:"errorCount"`
	APIKeysSet         bool                         `json:"apiKeysSet"`
}
type ExchangeAssetPairResponses struct {
	ErrorCount        int64              `json:"errorCount"`
	AssetType         asset.Item         `json:"asset"`
	CurrencyPair      currency.Pair      `json:"currency"`
	EndpointResponses []EndpointResponse `json:"responses"`
}
type EndpointResponse struct {
	Function   string          `json:"function"`
	Error      string          `json:"error"`
	Response   interface{}     `json:"response"`
	SentParams json.RawMessage `json:"sentParams"`
}
type Bank struct {
	BankAccountName               string  `json:"bankAccountName"`
	BankAccountNumber             string  `json:"bankAccountNumber"`
	BankAddress                   string  `json:"bankAddress"`
	BankCity                      string  `json:"bankCity"`
	BankCountry                   string  `json:"bankCountry"`
	BankName                      string  `json:"bankName"`
	BankPostalCode                string  `json:"bankPostalCode"`
	Iban                          string  `json:"iban"`
	IntermediaryBankAccountName   string  `json:"intermediaryBankAccountName"`
	IntermediaryBankAccountNumber float64 `json:"intermediaryBankAccountNumber"`
	IntermediaryBankAddress       string  `json:"intermediaryBankAddress"`
	IntermediaryBankCity          string  `json:"intermediaryBankCity"`
	IntermediaryBankCountry       string  `json:"intermediaryBankCountry"`
	IntermediaryBankName          string  `json:"intermediaryBankName"`
	IntermediaryBankPostalCode    string  `json:"intermediaryBankPostalCode"`
	IntermediaryIban              string  `json:"intermediaryIban"`
	IntermediaryIsExpressWire     bool    `json:"intermediaryIsExpressWire"`
	IntermediarySwiftCode         string  `json:"intermediarySwiftCode"`
	IsExpressWire                 bool    `json:"isExpressWire"`
	RequiresIntermediaryBank      bool    `json:"requiresIntermediaryBank"`
	SwiftCode                     string  `json:"swiftCode"`
	BankCode                      float64 `json:"bankCode"`
	IntermediaryBankCode          float64 `json:"intermediaryBankCode"`
}
type OrderSubmission struct {
	OrderSide string  `json:"orderSide"`
	OrderType string  `json:"orderType"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	OrderID   string  `json:"orderID"`
}
