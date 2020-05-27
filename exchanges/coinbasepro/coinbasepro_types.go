package coinbasepro

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
)

type Product struct {
	ID             string      `json:"id"`
	BaseCurrency   string      `json:"base_currency"`
	QuoteCurrency  string      `json:"quote_currency"`
	BaseMinSize    float64     `json:"base_min_size,string"`
	BaseMaxSize    interface{} `json:"base_max_size"`
	QuoteIncrement float64     `json:"quote_increment,string"`
	DisplayName    string      `json:"string"`
}
type Ticker struct {
	TradeID int64     `json:"trade_id"`
	Ask     float64   `json:"ask,string"`
	Bid     float64   `json:"bid,string"`
	Price   float64   `json:"price,string"`
	Size    float64   `json:"size,string"`
	Volume  float64   `json:"volume,string"`
	Time    time.Time `json:"time"`
}
type Trade struct {
	TradeID int64   `json:"trade_id"`
	Price   float64 `json:"price,string"`
	Size    float64 `json:"size,string"`
	Time    string  `json:"time"`
	Side    string  `json:"side"`
}
type History struct {
	Time   int64   `json:"time"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Volume float64 `json:"volume"`
}
type Stats struct {
	Open   float64 `json:"open,string"`
	High   float64 `json:"high,string"`
	Low    float64 `json:"low,string"`
	Volume float64 `json:"volume,string"`
}
type Currency struct {
	ID      string
	Name    string
	MinSize float64 `json:"min_size,string"`
}
type ServerTime struct {
	ISO   string  `json:"iso"`
	Epoch float64 `json:"epoch"`
}
type AccountResponse struct {
	ID            string  `json:"id"`
	Currency      string  `json:"currency"`
	Balance       float64 `json:"balance,string"`
	Available     float64 `json:"available,string"`
	Hold          float64 `json:"hold,string"`
	ProfileID     string  `json:"profile_id"`
	MarginEnabled bool    `json:"margin_enabled"`
	FundedAmount  float64 `json:"funded_amount,string"`
	DefaultAmount float64 `json:"default_amount,string"`
}
type AccountLedgerResponse struct {
	ID        string      `json:"id"`
	CreatedAt string      `json:"created_at"`
	Amount    float64     `json:"amount,string"`
	Balance   float64     `json:"balance,string"`
	Type      string      `json:"type"`
	Details   interface{} `json:"details"`
}
type AccountHolds struct {
	ID        string  `json:"id"`
	AccountID string  `json:"account_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Amount    float64 `json:"amount,string"`
	Type      string  `json:"type"`
	Reference string  `json:"ref"`
}
type GeneralizedOrderResponse struct {
	ID             string  `json:"id"`
	Price          float64 `json:"price,string"`
	Size           float64 `json:"size,string"`
	ProductID      string  `json:"product_id"`
	Side           string  `json:"side"`
	Stp            string  `json:"stp"`
	Type           string  `json:"type"`
	TimeInForce    string  `json:"time_in_force"`
	PostOnly       bool    `json:"post_only"`
	CreatedAt      string  `json:"created_at"`
	FillFees       float64 `json:"fill_fees,string"`
	FilledSize     float64 `json:"filled_size,string"`
	ExecutedValue  float64 `json:"executed_value,string"`
	Status         string  `json:"status"`
	Settled        bool    `json:"settled"`
	Funds          float64 `json:"funds,string"`
	SpecifiedFunds float64 `json:"specified_funds,string"`
	DoneReason     string  `json:"done_reason"`
	DoneAt         string  `json:"done_at"`
}
type Funding struct {
	ID            string  `json:"id"`
	OrderID       string  `json:"order_id"`
	ProfileID     string  `json:"profile_id"`
	Amount        float64 `json:"amount,string"`
	Status        string  `json:"status"`
	CreatedAt     string  `json:"created_at"`
	Currency      string  `json:"currency"`
	RepaidAmount  float64 `json:"repaid_amount"`
	DefaultAmount float64 `json:"default_amount,string"`
	RepaidDefault bool    `json:"repaid_default"`
}
type MarginTransfer struct {
	CreatedAt       string  `json:"created_at"`
	ID              string  `json:"id"`
	UserID          string  `json:"user_id"`
	ProfileID       string  `json:"profile_id"`
	MarginProfileID string  `json:"margin_profile_id"`
	Type            string  `json:"type"`
	Amount          float64 `json:"amount,string"`
	Currency        string  `json:"currency"`
	AccountID       string  `json:"account_id"`
	MarginAccountID string  `json:"margin_account_id"`
	MarginProductID string  `json:"margin_product_id"`
	Status          string  `json:"status"`
	Nonce           int     `json:"nonce"`
}
type AccountOverview struct {
	Status  string `json:"status"`
	Funding struct {
		MaxFundingValue   float64 `json:"max_funding_value,string"`
		FundingValue      float64 `json:"funding_value,string"`
		OldestOutstanding struct {
			ID        string  `json:"id"`
			OrderID   string  `json:"order_id"`
			CreatedAt string  `json:"created_at"`
			Currency  string  `json:"currency"`
			AccountID string  `json:"account_id"`
			Amount    float64 `json:"amount,string"`
		} `json:"oldest_outstanding"`
	} `json:"funding"`
	Accounts struct {
		LTC Account `json:"LTC"`
		ETH Account `json:"ETH"`
		USD Account `json:"USD"`
		BTC Account `json:"BTC"`
	} `json:"accounts"`
	MarginCall struct {
		Active bool    `json:"active"`
		Price  float64 `json:"price,string"`
		Side   string  `json:"side"`
		Size   float64 `json:"size,string"`
		Funds  float64 `json:"funds,string"`
	} `json:"margin_call"`
	UserID    string `json:"user_id"`
	ProfileID string `json:"profile_id"`
	Position  struct {
		Type       string  `json:"type"`
		Size       float64 `json:"size,string"`
		Complement float64 `json:"complement,string"`
		MaxSize    float64 `json:"max_size,string"`
	} `json:"position"`
	ProductID string `json:"product_id"`
}
type Account struct {
	ID            string  `json:"id"`
	Balance       float64 `json:"balance,string"`
	Hold          float64 `json:"hold,string"`
	FundedAmount  float64 `json:"funded_amount,string"`
	DefaultAmount float64 `json:"default_amount,string"`
}
type PaymentMethod struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Currency      string `json:"currency"`
	PrimaryBuy    bool   `json:"primary_buy"`
	PrimarySell   bool   `json:"primary_sell"`
	AllowBuy      bool   `json:"allow_buy"`
	AllowSell     bool   `json:"allow_sell"`
	AllowDeposits bool   `json:"allow_deposits"`
	AllowWithdraw bool   `json:"allow_withdraw"`
	Limits        struct {
		Buy        []LimitInfo `json:"buy"`
		InstantBuy []LimitInfo `json:"instant_buy"`
		Sell       []LimitInfo `json:"sell"`
		Deposit    []LimitInfo `json:"deposit"`
	} `json:"limits"`
}
type LimitInfo struct {
	PeriodInDays int `json:"period_in_days"`
	Total        struct {
		Amount   float64 `json:"amount,string"`
		Currency string  `json:"currency"`
	} `json:"total"`
}
type DepositWithdrawalInfo struct {
	ID       string  `json:"id"`
	Amount   float64 `json:"amount,string"`
	Currency string  `json:"currency"`
	PayoutAt string  `json:"payout_at"`
}
type CoinbaseAccounts struct {
	ID                     string  `json:"id"`
	Name                   string  `json:"name"`
	Balance                float64 `json:"balance,string"`
	Currency               string  `json:"currency"`
	Type                   string  `json:"type"`
	Primary                bool    `json:"primary"`
	Active                 bool    `json:"active"`
	WireDepositInformation struct {
		AccountNumber string `json:"account_number"`
		RoutingNumber string `json:"routing_number"`
		BankName      string `json:"bank_name"`
		BankAddress   string `json:"bank_address"`
		BankCountry   struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"bank_country"`
		AccountName    string `json:"account_name"`
		AccountAddress string `json:"account_address"`
		Reference      string `json:"reference"`
	} `json:"wire_deposit_information"`
	SepaDepositInformation struct {
		Iban            string `json:"iban"`
		Swift           string `json:"swift"`
		BankName        string `json:"bank_name"`
		BankAddress     string `json:"bank_address"`
		BankCountryName string `json:"bank_country_name"`
		AccountName     string `json:"account_name"`
		AccountAddress  string `json:"account_address"`
		Reference       string `json:"reference"`
	} `json:"sep_deposit_information"`
}
type Report struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at"`
	ExpiresAt   string `json:"expires_at"`
	FileURL     string `json:"file_url"`
	Params      struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
	} `json:"params"`
}
type Volume struct {
	ProductID      string  `json:"product_id"`
	ExchangeVolume float64 `json:"exchange_volume,string"`
	Volume         float64 `json:"volume,string"`
	RecordedAt     string  `json:"recorded_at"`
}
type OrderbookResponse struct {
	Sequence int64           `json:"sequence"`
	Bids     [][]interface{} `json:"bids"`
	Asks     [][]interface{} `json:"asks"`
}
type FillResponse struct {
	TradeID   int     `json:"trade_id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price,string"`
	Size      float64 `json:"size,string"`
	OrderID   string  `json:"order_id"`
	CreatedAt string  `json:"created_at"`
	Liquidity string  `json:"liquidity"`
	Fee       float64 `json:"fee,string"`
	Settled   bool    `json:"settled"`
	Side      string  `json:"side"`
}
type WebsocketSubscribe struct {
	Type       string       `json:"type"`
	ProductID  string       `json:"product_id,omitempty"`
	Channels   []WsChannels `json:"channels,omitempty"`
	Signature  string       `json:"signature,omitempty"`
	Key        string       `json:"key,omitempty"`
	Passphrase string       `json:"passphrase,omitempty"`
	Timestamp  string       `json:"timestamp,omitempty"`
}
type WsChannels struct {
	Name       string   `json:"name"`
	ProductIDs []string `json:"product_ids"`
}
type wsOrderReceived struct {
	Type          string    `json:"type"`
	OrderID       string    `json:"order_id"`
	OrderType     string    `json:"order_type"`
	Size          float64   `json:"size,string"`
	Price         float64   `json:"price,omitempty,string"`
	Funds         float64   `json:"funds,omitempty,string"`
	Side          string    `json:"side"`
	ClientOID     string    `json:"client_oid"`
	ProductID     string    `json:"product_id"`
	Sequence      int64     `json:"sequence"`
	Time          time.Time `json:"time"`
	RemainingSize float64   `json:"remaining_size,string"`
	NewSize       float64   `json:"new_size,string"`
	OldSize       float64   `json:"old_size,string"`
	Reason        string    `json:"reason"`
	Timestamp     float64   `json:"timestamp,string"`
	UserID        string    `json:"user_id"`
	ProfileID     string    `json:"profile_id"`
	StopType      string    `json:"stop_type"`
	StopPrice     float64   `json:"stop_price,string"`
	TakerFeeRate  float64   `json:"taker_fee_rate,string"`
	Private       bool      `json:"private"`
	TradeID       int64     `json:"trade_id"`
	MakerOrderID  string    `json:"maker_order_id"`
	TakerOrderID  string    `json:"taker_order_id"`
	TakerUserID   string    `json:"taker_user_id"`
}
type WebsocketHeartBeat struct {
	Type        string `json:"type"`
	Sequence    int64  `json:"sequence"`
	LastTradeID int64  `json:"last_trade_id"`
	ProductID   string `json:"product_id"`
	Time        string `json:"time"`
}
type WebsocketTicker struct {
	Type      string        `json:"type"`
	Sequence  int64         `json:"sequence"`
	ProductID currency.Pair `json:"product_id"`
	Price     float64       `json:"price,string"`
	Open24H   float64       `json:"open_24h,string"`
	Volume24H float64       `json:"volume_24h,string"`
	Low24H    float64       `json:"low_24h,string"`
	High24H   float64       `json:"high_24h,string"`
	Volume30D float64       `json:"volume_30d,string"`
	BestBid   float64       `json:"best_bid,string"`
	BestAsk   float64       `json:"best_ask,string"`
	Side      string        `json:"side"`
	Time      time.Time     `json:"time"`
	TradeID   int64         `json:"trade_id"`
	LastSize  float64       `json:"last_size,string"`
}
type WebsocketOrderbookSnapshot struct {
	ProductID string          `json:"product_id"`
	Type      string          `json:"type"`
	Bids      [][]interface{} `json:"bids"`
	Asks      [][]interface{} `json:"asks"`
}
type wsMsgType struct {
	Type      string `json:"type"`
	Sequence  int64  `json:"sequence"`
	ProductID string `json:"product_id"`
}
type wsStatus struct {
	Currencies []struct {
		ConvertibleTo []string    `json:"convertible_to"`
		Details       struct{}    `json:"details"`
		ID            string      `json:"id"`
		MaxPrecision  float64     `json:"max_precision,string"`
		MinSize       float64     `json:"min_size,string"`
		Name          string      `json:"name"`
		Status        string      `json:"status"`
		StatusMessage interface{} `json:"status_message"`
	} `json:"currencies"`
	Products []struct {
		BaseCurrency   string      `json:"base_currency"`
		BaseIncrement  float64     `json:"base_increment,string"`
		BaseMaxSize    float64     `json:"base_max_size,string"`
		BaseMinSize    float64     `json:"base_min_size,string"`
		CancelOnly     bool        `json:"cancel_only"`
		DisplayName    string      `json:"display_name"`
		ID             string      `json:"id"`
		LimitOnly      bool        `json:"limit_only"`
		MaxMarketFunds float64     `json:"max_market_funds,string"`
		MinMarketFunds float64     `json:"min_market_funds,string"`
		PostOnly       bool        `json:"post_only"`
		QuoteCurrency  string      `json:"quote_currency"`
		QuoteIncrement float64     `json:"quote_increment,string"`
		Status         string      `json:"status"`
		StatusMessage  interface{} `json:"status_message"`
	} `json:"products"`
	Type string `json:"type"`
}
