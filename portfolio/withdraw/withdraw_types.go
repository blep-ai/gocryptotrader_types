package withdraw
import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/blep-ai/gocryptotrader_types/common/cache"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/portfolio/banking"
)
type CryptoRequest struct {
	Address    string
	AddressTag string
	FeeAmount  float64
}
type FiatRequest struct {
	Bank *banking.Account

	IsExpressWire bool
	// Intermediary bank information
	RequiresIntermediaryBank      bool
	IntermediaryBankAccountNumber float64
	IntermediaryBankName          string
	IntermediaryBankAddress       string
	IntermediaryBankCity          string
	IntermediaryBankCountry       string
	IntermediaryBankPostalCode    string
	IntermediarySwiftCode         string
	IntermediaryBankCode          float64
	IntermediaryIBAN              string
	WireCurrency                  string
}
type Request struct {
	Exchange    string        `json:"exchange"`
	Currency    currency.Code `json:"currency"`
	Description string        `json:"description"`
	Amount      float64       `json:"amount"`
	Type        RequestType   `json:"type"`

	TradePassword   string
	OneTimePassword int64
	PIN             int64

	Crypto *CryptoRequest `json:",omitempty"`
	Fiat   *FiatRequest   `json:",omitempty"`
}
type Response struct {
	ID uuid.UUID `json:"id"`

	Exchange       *ExchangeResponse `json:"exchange"`
	RequestDetails *Request          `json:"request_details"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ExchangeResponse struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
}
