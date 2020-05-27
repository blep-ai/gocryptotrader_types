package poloniex

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
)
type Ticker struct {
	ID            float64 `json:"id"`
	Last          float64 `json:"last,string"`
	LowestAsk     float64 `json:"lowestAsk,string"`
	HighestBid    float64 `json:"highestBid,string"`
	PercentChange float64 `json:"percentChange,string"`
	BaseVolume    float64 `json:"baseVolume,string"`
	QuoteVolume   float64 `json:"quoteVolume,string"`
	IsFrozen      int64   `json:"isFrozen,string"`
	High24Hr      float64 `json:"high24hr,string"`
	Low24Hr       float64 `json:"low24hr,string"`
}
type OrderbookResponseAll struct {
	Data map[string]OrderbookResponse
}
type CompleteBalances struct {
	Currency map[string]CompleteBalance
}
type OrderbookResponse struct {
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
	IsFrozen string          `json:"isFrozen"`
	Error    string          `json:"error"`
}
type OrderbookItem struct {
	Price  float64
	Amount float64
}
type OrderbookAll struct {
	Data map[string]Orderbook
}
type Orderbook struct {
	Asks []OrderbookItem `json:"asks"`
	Bids []OrderbookItem `json:"bids"`
}
type TradeHistory struct {
	GlobalTradeID int64   `json:"globalTradeID"`
	TradeID       int64   `json:"tradeID"`
	Date          string  `json:"date"`
	Type          string  `json:"type"`
	Rate          float64 `json:"rate,string"`
	Amount        float64 `json:"amount,string"`
	Total         float64 `json:"total,string"`
}
type ChartData struct {
	Date            int64   `json:"date"`
	High            float64 `json:"high"`
	Low             float64 `json:"low"`
	Open            float64 `json:"open"`
	Close           float64 `json:"close"`
	Volume          float64 `json:"volume"`
	QuoteVolume     float64 `json:"quoteVolume"`
	WeightedAverage float64 `json:"weightedAverage"`
	Error           string  `json:"error"`
}
type Currencies struct {
	ID                 float64     `json:"id"`
	Name               string      `json:"name"`
	MaxDailyWithdrawal string      `json:"maxDailyWithdrawal"`
	TxFee              float64     `json:"txFee,string"`
	MinConfirmations   int64       `json:"minConf"`
	DepositAddresses   interface{} `json:"depositAddress"`
	Disabled           int64       `json:"disabled"`
	Delisted           int64       `json:"delisted"`
	Frozen             int64       `json:"frozen"`
}
type LoanOrder struct {
	Rate     float64 `json:"rate,string"`
	Amount   float64 `json:"amount,string"`
	RangeMin int64   `json:"rangeMin"`
	RangeMax int64   `json:"rangeMax"`
}
type LoanOrders struct {
	Offers  []LoanOrder `json:"offers"`
	Demands []LoanOrder `json:"demands"`
}
type Balance struct {
	Currency map[string]float64
}
type CompleteBalance struct {
	Available float64
	OnOrders  float64
	BTCValue  float64
}
type DepositAddresses struct {
	Addresses map[string]string
}
type DepositsWithdrawals struct {
	Deposits []struct {
		Currency      string  `json:"currency"`
		Address       string  `json:"address"`
		Amount        float64 `json:"amount,string"`
		Confirmations int64   `json:"confirmations"`
		TransactionID string  `json:"txid"`
		Timestamp     int64   `json:"timestamp"`
		Status        string  `json:"status"`
	} `json:"deposits"`
	Withdrawals []struct {
		WithdrawalNumber int64   `json:"withdrawalNumber"`
		Currency         string  `json:"currency"`
		Address          string  `json:"address"`
		Amount           float64 `json:"amount,string"`
		Confirmations    int64   `json:"confirmations"`
		TransactionID    string  `json:"txid"`
		Timestamp        int64   `json:"timestamp"`
		Status           string  `json:"status"`
		IPAddress        string  `json:"ipAddress"`
	} `json:"withdrawals"`
}
type Order struct {
	OrderNumber int64   `json:"orderNumber,string"`
	Type        string  `json:"type"`
	Rate        float64 `json:"rate,string"`
	Amount      float64 `json:"amount,string"`
	Total       float64 `json:"total,string"`
	Date        string  `json:"date"`
	Margin      float64 `json:"margin"`
}
type OpenOrdersResponseAll struct {
	Data map[string][]Order
}
type OpenOrdersResponse struct {
	Data []Order
}
type AuthenticatedTradeHistory struct {
	GlobalTradeID int64   `json:"globalTradeID"`
	TradeID       int64   `json:"tradeID,string"`
	Date          string  `json:"date"`
	Rate          float64 `json:"rate,string"`
	Amount        float64 `json:"amount,string"`
	Total         float64 `json:"total,string"`
	Fee           float64 `json:"fee,string"`
	OrderNumber   int64   `json:"orderNumber,string"`
	Type          string  `json:"type"`
	Category      string  `json:"category"`
}
type AuthenticatedTradeHistoryAll struct {
	Data map[string][]AuthenticatedTradeHistory
}
type AuthenticatedTradeHistoryResponse struct {
	Data []AuthenticatedTradeHistory
}
type ResultingTrades struct {
	Amount  float64 `json:"amount,string"`
	Date    string  `json:"date"`
	Rate    float64 `json:"rate,string"`
	Total   float64 `json:"total,string"`
	TradeID int64   `json:"tradeID,string"`
	Type    string  `json:"type"`
}
type OrderResponse struct {
	OrderNumber int64             `json:"orderNumber,string"`
	Trades      []ResultingTrades `json:"resultingTrades"`
}
type GenericResponse struct {
	Success int64  `json:"success"`
	Error   string `json:"error"`
}
type MoveOrderResponse struct {
	Success     int64                        `json:"success"`
	Error       string                       `json:"error"`
	OrderNumber int64                        `json:"orderNumber,string"`
	Trades      map[string][]ResultingTrades `json:"resultingTrades"`
}
type Withdraw struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}
type Fee struct {
	MakerFee        float64 `json:"makerFee,string"`
	TakerFee        float64 `json:"takerFee,string"`
	ThirtyDayVolume float64 `json:"thirtyDayVolume,string"`
}
type Margin struct {
	TotalValue    float64 `json:"totalValue,string"`
	ProfitLoss    float64 `json:"pl,string"`
	LendingFees   float64 `json:"lendingFees,string"`
	NetValue      float64 `json:"netValue,string"`
	BorrowedValue float64 `json:"totalBorrowedValue,string"`
	CurrentMargin float64 `json:"currentMargin,string"`
}
type MarginPosition struct {
	Amount           float64 `json:"amount,string"`
	Total            float64 `json:"total,string"`
	BasePrice        float64 `json:"basePrice,string"`
	LiquidationPrice float64 `json:"liquidationPrice"`
	ProfitLoss       float64 `json:"pl,string"`
	LendingFees      float64 `json:"lendingFees,string"`
	Type             string  `json:"type"`
}
type LoanOffer struct {
	ID        int64   `json:"id"`
	Rate      float64 `json:"rate,string"`
	Amount    float64 `json:"amount,string"`
	Duration  int64   `json:"duration"`
	AutoRenew bool    `json:"autoRenew"`
	Date      string  `json:"date"`
}
type ActiveLoans struct {
	Provided []LoanOffer `json:"provided"`
	Used     []LoanOffer `json:"used"`
}
type LendingHistory struct {
	ID       int64   `json:"id"`
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate,string"`
	Amount   float64 `json:"amount,string"`
	Duration float64 `json:"duration,string"`
	Interest float64 `json:"interest,string"`
	Fee      float64 `json:"fee,string"`
	Earned   float64 `json:"earned,string"`
	Open     string  `json:"open"`
	Close    string  `json:"close"`
}
type WebsocketTicker struct {
	CurrencyPair  string
	Last          float64
	LowestAsk     float64
	HighestBid    float64
	PercentChange float64
	BaseVolume    float64
	QuoteVolume   float64
	IsFrozen      bool
	High          float64
	Low           float64
}
type WebsocketTrollboxMessage struct {
	MessageNumber float64
	Username      string
	Message       string
	Reputation    float64
}
type WsCommand struct {
	Command string      `json:"command"`
	Channel interface{} `json:"channel"`
	APIKey  string      `json:"key,omitempty"`
	Payload string      `json:"payload,omitempty"`
	Sign    string      `json:"sign,omitempty"`
}
type WsTicker struct {
	LastPrice              float64
	LowestAsk              float64
	HighestBid             float64
	PercentageChange       float64
	BaseCurrencyVolume24H  float64
	QuoteCurrencyVolume24H float64
	IsFrozen               bool
	HighestTradeIn24H      float64
	LowestTradePrice24H    float64
}
type WsTrade struct {
	Symbol    string
	TradeID   int64
	Side      string
	Volume    float64
	Price     float64
	Timestamp int64
}
type WsAccountBalanceUpdateResponse struct {
	currencyID float64
	wallet     string
	amount     float64
}
type WsOrderUpdateResponse struct {
	OrderNumber float64
	NewAmount   string
}
type WsTradeNotificationResponse struct {
	TradeID       float64
	Rate          float64
	Amount        float64
	FeeMultiplier float64
	FundingType   float64
	OrderNumber   float64
	TotalFee      float64
	Date          time.Time
}
type WsAuthorisationRequest struct {
	Command string `json:"command"`
	Channel int64  `json:"channel"`
	Sign    string `json:"sign"`
	Key     string `json:"key"`
	Payload string `json:"payload"`
}
