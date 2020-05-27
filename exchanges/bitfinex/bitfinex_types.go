package bitfinex

import "time"

// AcceptedOrderType defines the accepted market types, exchange strings denote
// non-contract order types.
var AcceptedOrderType = []string{"market", "limit", "stop", "trailing-stop",
	"fill-or-kill", "exchange market", "exchange limit", "exchange stop",
	"exchange trailing-stop", "exchange fill-or-kill"}

// AcceptedWalletNames defines different wallets supported by the exchange
var AcceptedWalletNames = []string{"trading", "exchange", "deposit", "margin",
	"funding"}

// AcceptableMethods defines a map of currency codes to methods
var AcceptableMethods = make(map[string]string)

type Ticker struct {
	FlashReturnRate    float64
	Bid                float64
	BidPeriod          int64
	BidSize            float64
	Ask                float64
	AskPeriod          int64
	AskSize            float64
	DailyChange        float64
	DailyChangePerc    float64
	Last               float64
	Volume             float64
	High               float64
	Low                float64
	FFRAmountAvailable float64
}
type Stat struct {
	Period int64   `json:"period"`
	Volume float64 `json:"volume,string"`
}
type FundingBook struct {
	Bids []FundingBookItem `json:"bids"`
	Asks []FundingBookItem `json:"asks"`
}
type Book struct {
	OrderID int64
	Price   float64
	Rate    float64
	Period  float64
	Count   int64
	Amount  float64
}
type Orderbook struct {
	Bids []Book
	Asks []Book
}
type Trade struct {
	Timestamp int64
	TID       int64
	Price     float64
	Amount    float64
	Exchange  string
	Rate      float64
	Period    int64
	Type      string
}
type Lendbook struct {
	Bids []Book `json:"bids"`
	Asks []Book `json:"asks"`
}
type FundingBookItem struct {
	Rate            float64 `json:"rate,string"`
	Amount          float64 `json:"amount,string"`
	Period          int     `json:"period"`
	Timestamp       string  `json:"timestamp"`
	FlashReturnRate string  `json:"frr"`
}
type Lends struct {
	Rate       float64 `json:"rate,string"`
	AmountLent float64 `json:"amount_lent,string"`
	AmountUsed float64 `json:"amount_used,string"`
	Timestamp  int64   `json:"timestamp"`
}
type SymbolDetails struct {
	Pair             string  `json:"pair"`
	PricePrecision   int     `json:"price_precision"`
	InitialMargin    float64 `json:"initial_margin,string"`
	MinimumMargin    float64 `json:"minimum_margin,string"`
	MaximumOrderSize float64 `json:"maximum_order_size,string"`
	MinimumOrderSize float64 `json:"minimum_order_size,string"`
	Expiration       string  `json:"expiration"`
}
type AccountInfoFull struct {
	Info    []AccountInfo
	Message string `json:"message"`
}
type AccountInfo struct {
	MakerFees float64           `json:"maker_fees,string"`
	TakerFees float64           `json:"taker_fees,string"`
	Fees      []AccountInfoFees `json:"fees"`
	Message   string            `json:"message"`
}
type AccountInfoFees struct {
	Pairs     string  `json:"pairs"`
	MakerFees float64 `json:"maker_fees,string"`
	TakerFees float64 `json:"taker_fees,string"`
}
type AccountFees struct {
	Withdraw map[string]interface{} `json:"withdraw"`
}
type AccountSummary struct {
	TradeVolumePer30D []Currency `json:"trade_vol_30d"`
	FundingProfit30D  []Currency `json:"funding_profit_30d"`
	MakerFee          float64    `json:"maker_fee"`
	TakerFee          float64    `json:"taker_fee"`
}
type Currency struct {
	Currency string  `json:"curr"`
	Volume   float64 `json:"vol,string"`
	Amount   float64 `json:"amount,string"`
}
type DepositResponse struct {
	Result   string `json:"string"`
	Method   string `json:"method"`
	Currency string `json:"currency"`
	Address  string `json:"address"`
}
type KeyPermissions struct {
	Account   Permission `json:"account"`
	History   Permission `json:"history"`
	Orders    Permission `json:"orders"`
	Positions Permission `json:"positions"`
	Funding   Permission `json:"funding"`
	Wallets   Permission `json:"wallets"`
	Withdraw  Permission `json:"withdraw"`
}
type Permission struct {
	Read  bool `json:"read"`
	Write bool `json:"write"`
}
type MarginInfo struct {
	Info    MarginData
	Message string `json:"message"`
}
type MarginData struct {
	MarginBalance     float64        `json:"margin_balance,string"`
	TradableBalance   float64        `json:"tradable_balance,string"`
	UnrealizedPL      int64          `json:"unrealized_pl"`
	UnrealizedSwap    int64          `json:"unrealized_swap"`
	NetValue          float64        `json:"net_value,string"`
	RequiredMargin    int64          `json:"required_margin"`
	Leverage          float64        `json:"leverage,string"`
	MarginRequirement float64        `json:"margin_requirement,string"`
	MarginLimits      []MarginLimits `json:"margin_limits"`
}
type MarginLimits struct {
	OnPair            string  `json:"on_pair"`
	InitialMargin     float64 `json:"initial_margin,string"`
	MarginRequirement float64 `json:"margin_requirement,string"`
	TradableBalance   float64 `json:"tradable_balance,string"`
}
type Balance struct {
	Type      string  `json:"type"`
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount,string"`
	Available float64 `json:"available,string"`
}
type WalletTransfer struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type Withdrawal struct {
	Status       string  `json:"status"`
	Message      string  `json:"message"`
	WithdrawalID int64   `json:"withdrawal_id"`
	Fees         string  `json:"fees"`
	WalletType   string  `json:"wallettype"`
	Method       string  `json:"method"`
	Address      string  `json:"address"`
	Invoice      string  `json:"invoice"`
	PaymentID    string  `json:"payment_id"`
	Amount       float64 `json:"amount,string"`
}
type Order struct {
	ID                    int64   `json:"id"`
	Symbol                string  `json:"symbol"`
	Exchange              string  `json:"exchange"`
	Price                 float64 `json:"price,string"`
	AverageExecutionPrice float64 `json:"avg_execution_price,string"`
	Side                  string  `json:"side"`
	Type                  string  `json:"type"`
	Timestamp             string  `json:"timestamp"`
	IsLive                bool    `json:"is_live"`
	IsCancelled           bool    `json:"is_cancelled"`
	IsHidden              bool    `json:"is_hidden"`
	WasForced             bool    `json:"was_forced"`
	OriginalAmount        float64 `json:"original_amount,string"`
	RemainingAmount       float64 `json:"remaining_amount,string"`
	ExecutedAmount        float64 `json:"executed_amount,string"`
	OrderID               int64   `json:"order_id,omitempty"`
}
type OrderMultiResponse struct {
	Orders []Order `json:"order_ids"`
	Status string  `json:"status"`
}
type PlaceOrder struct {
	Symbol   string  `json:"symbol"`
	Amount   float64 `json:"amount,string"`
	Price    float64 `json:"price,string"`
	Exchange string  `json:"exchange"`
	Side     string  `json:"side"`
	Type     string  `json:"type"`
}
type GenericResponse struct {
	Result string `json:"result"`
}
type Position struct {
	ID        int64   `json:"id"`
	Symbol    string  `json:"string"`
	Status    string  `json:"active"`
	Base      float64 `json:"base,string"`
	Amount    float64 `json:"amount,string"`
	Timestamp string  `json:"timestamp"`
	Swap      float64 `json:"swap,string"`
	PL        float64 `json:"pl,string"`
}
type BalanceHistory struct {
	Currency    string  `json:"currency"`
	Amount      float64 `json:"amount,string"`
	Balance     float64 `json:"balance,string"`
	Description string  `json:"description"`
	Timestamp   string  `json:"timestamp"`
}
type MovementHistory struct {
	ID               int64   `json:"id"`
	TxID             int64   `json:"txid"`
	Currency         string  `json:"currency"`
	Method           string  `json:"method"`
	Type             string  `json:"withdrawal"`
	Amount           float64 `json:"amount,string"`
	Description      string  `json:"description"`
	Address          string  `json:"address"`
	Status           string  `json:"status"`
	Timestamp        string  `json:"timestamp"`
	TimestampCreated string  `json:"timestamp_created"`
	Fee              float64 `json:"fee"`
}
type TradeHistory struct {
	Price       float64 `json:"price,string"`
	Amount      float64 `json:"amount,string"`
	Timestamp   string  `json:"timestamp"`
	Exchange    string  `json:"exchange"`
	Type        string  `json:"type"`
	FeeCurrency string  `json:"fee_currency"`
	FeeAmount   float64 `json:"fee_amount,string"`
	TID         int64   `json:"tid"`
	OrderID     int64   `json:"order_id"`
}
type Offer struct {
	ID              int64   `json:"id"`
	Currency        string  `json:"currency"`
	Rate            float64 `json:"rate,string"`
	Period          int64   `json:"period"`
	Direction       string  `json:"direction"`
	Timestamp       string  `json:"timestamp"`
	Type            string  `json:"type"`
	IsLive          bool    `json:"is_live"`
	IsCancelled     bool    `json:"is_cancelled"`
	OriginalAmount  float64 `json:"original_amount,string"`
	RemainingAmount float64 `json:"remaining_amount,string"`
	ExecutedAmount  float64 `json:"executed_amount,string"`
}
type MarginFunds struct {
	ID         int64   `json:"id"`
	PositionID int64   `json:"position_id"`
	Currency   string  `json:"currency"`
	Rate       float64 `json:"rate,string"`
	Period     int     `json:"period"`
	Amount     float64 `json:"amount,string"`
	Timestamp  string  `json:"timestamp"`
	AutoClose  bool    `json:"auto_close"`
}
type MarginTotalTakenFunds struct {
	PositionPair string  `json:"position_pair"`
	TotalSwaps   float64 `json:"total_swaps,string"`
}
type Fee struct {
	Currency  string
	TakerFees float64
	MakerFees float64
}
type WebsocketChanInfo struct {
	Channel string
	Pair    string
}
type WebsocketBook struct {
	Price  float64
	ID     int64
	Amount float64
}
type WebsocketTrade struct {
	ID        int64
	Timestamp int64
	Price     float64
	Amount    float64
	// Funding rate of the trade
	Rate float64
	// Funding offer period in days
	Period int64
}
type Candle struct {
	Timestamp int64
	Open      float64
	Close     float64
	High      float64
	Low       float64
	Volume    float64
}
type LeaderboardEntry struct {
	Timestamp     time.Time
	Username      string
	Ranking       int
	Value         float64
	TwitterHandle string
}
type WebsocketTicker struct {
	Bid             float64
	BidSize         float64
	Ask             float64
	AskSize         float64
	DailyChange     float64
	DialyChangePerc float64
	LastPrice       float64
	Volume          float64
}
type WebsocketPosition struct {
	Pair              string
	Status            string
	Amount            float64
	Price             float64
	MarginFunding     float64
	MarginFundingType int64
	ProfitLoss        float64
	ProfitLossPercent float64
	LiquidationPrice  float64
	Leverage          float64
}
type WebsocketWallet struct {
	Name              string
	Currency          string
	Balance           float64
	UnsettledInterest float64
}
type WebsocketOrder struct {
	OrderID    int64
	Pair       string
	Amount     float64
	OrigAmount float64
	OrderType  string
	Status     string
	Price      float64
	PriceAvg   float64
	Timestamp  int64
	Notify     int
}
type WebsocketTradeExecuted struct {
	TradeID        int64
	Pair           string
	Timestamp      int64
	OrderID        int64
	AmountExecuted float64
	PriceExecuted  float64
}
type WebsocketTradeData struct {
	TradeID        int64
	Pair           string
	Timestamp      int64
	OrderID        int64
	AmountExecuted float64
	PriceExecuted  float64
	OrderType      string
	OrderPrice     float64
	Maker          bool
	Fee            float64
	FeeCurrency    string
}
type ErrorCapture struct {
	Message string `json:"message"`
}
type WebsocketHandshake struct {
	Event   string  `json:"event"`
	Code    int64   `json:"code"`
	Version float64 `json:"version"`
}
type WsAuthRequest struct {
	Event         string `json:"event"`
	APIKey        string `json:"apiKey"`
	AuthPayload   string `json:"authPayload"`
	AuthSig       string `json:"authSig"`
	AuthNonce     string `json:"authNonce"`
	DeadManSwitch int64  `json:"dms,omitempty"`
}
type WsFundingOffer struct {
	ID             int64
	Symbol         string
	Created        int64
	Updated        int64
	Amount         float64
	OriginalAmount float64
	Type           string
	Flags          interface{}
	Status         string
	Rate           float64
	Period         int64
	Notify         bool
	Hidden         bool
	Insure         bool
	Renew          bool
	RateReal       float64
}
type WsCredit struct {
	ID           int64
	Symbol       string
	Side         string
	Created      int64
	Updated      int64
	Amount       float64
	Flags        interface{}
	Status       string
	Rate         float64
	Period       int64
	Opened       int64
	LastPayout   int64
	Notify       bool
	Hidden       bool
	Insure       bool
	Renew        bool
	RateReal     float64
	NoClose      bool
	PositionPair string
}
type WsWallet struct {
	Type              string
	Currency          string
	Balance           float64
	UnsettledInterest float64
	BalanceAvailable  float64
}
type WsBalanceInfo struct {
	TotalAssetsUnderManagement float64
	NetAssetsUnderManagement   float64
}
type WsFundingInfo struct {
	Symbol       string
	YieldLoan    float64
	YieldLend    float64
	DurationLoan float64
	DurationLend float64
}
type WsMarginInfoBase struct {
	UserProfitLoss float64
	UserSwaps      float64
	MarginBalance  float64
	MarginNet      float64
}
type WsFundingTrade struct {
	ID         int64
	Symbol     string
	MTSCreated int64
	OfferID    int64
	Amount     float64
	Rate       float64
	Period     int64
	Maker      bool
}
type WsNewOrderRequest struct {
	GroupID             int64   `json:"gid,omitempty"`
	CustomID            int64   `json:"cid,omitempty"`
	Type                string  `json:"type"`
	Symbol              string  `json:"symbol"`
	Amount              float64 `json:"amount,string"`
	Price               float64 `json:"price,string"`
	Leverage            int64   `json:"lev,omitempty"`
	TrailingPrice       float64 `json:"price_trailing,string,omitempty"`
	AuxiliaryLimitPrice float64 `json:"price_aux_limit,string,omitempty"`
	StopPrice           float64 `json:"price_oco_stop,string,omitempty"`
	Flags               int64   `json:"flags,omitempty"`
	TimeInForce         string  `json:"tif,omitempty"`
}
type WsUpdateOrderRequest struct {
	OrderID             int64   `json:"id,omitempty"`
	CustomID            int64   `json:"cid,omitempty"`
	CustomIDDate        string  `json:"cid_date,omitempty"`
	GroupID             int64   `json:"gid,omitempty"`
	Price               float64 `json:"price,string,omitempty"`
	Amount              float64 `json:"amount,string,omitempty"`
	Leverage            int64   `json:"lev,omitempty"`
	Delta               float64 `json:"delta,string,omitempty"`
	AuxiliaryLimitPrice float64 `json:"price_aux_limit,string,omitempty"`
	TrailingPrice       float64 `json:"price_trailing,string,omitempty"`
	Flags               int64   `json:"flags,omitempty"`
	TimeInForce         string  `json:"tif,omitempty"`
}
type WsCancelOrderRequest struct {
	OrderID      int64  `json:"id,omitempty"`
	CustomID     int64  `json:"cid,omitempty"`
	CustomIDDate string `json:"cid_date,omitempty"`
}
type WsCancelGroupOrdersRequest struct {
	OrderID      []int64   `json:"id,omitempty"`
	CustomID     [][]int64 `json:"cid,omitempty"`
	GroupOrderID []int64   `json:"gid,omitempty"`
}
type WsNewOfferRequest struct {
	Type   string  `json:"type,omitempty"`
	Symbol string  `json:"symbol,omitempty"`
	Amount float64 `json:"amount,string,omitempty"`
	Rate   float64 `json:"rate,string,omitempty"`
	Period float64 `json:"period,omitempty"`
	Flags  int64   `json:"flags,omitempty"`
}
type WsCancelOfferRequest struct {
	OrderID int64 `json:"id"`
}
type WsCancelAllOrdersRequest struct {
	All int64 `json:"all"`
}
