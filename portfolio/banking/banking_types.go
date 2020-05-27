package banking

type Account struct {
	Enabled             bool    `json:"enabled"`
	ID                  string  `json:"id,omitempty"`
	BankName            string  `json:"bankName"`
	BankAddress         string  `json:"bankAddress"`
	BankPostalCode      string  `json:"bankPostalCode"`
	BankPostalCity      string  `json:"bankPostalCity"`
	BankCountry         string  `json:"bankCountry"`
	AccountName         string  `json:"accountName"`
	AccountNumber       string  `json:"accountNumber"`
	SWIFTCode           string  `json:"swiftCode"`
	IBAN                string  `json:"iban"`
	BSBNumber           string  `json:"bsbNumber,omitempty"`
	BankCode            float64 `json:"bank_code,omitempty"`
	SupportedCurrencies string  `json:"supportedCurrencies"`
	SupportedExchanges  string  `json:"supportedExchanges,omitempty"`
}
