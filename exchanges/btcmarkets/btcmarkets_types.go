package btcmarkets

import (
	"errors"
	"time"
)
type Market struct {
	MarketID       string  `json:"marketId"`
	BaseAsset      string  `json:"baseAsset"`
	QuoteAsset     string  `json:"quoteAsset"`
	MinOrderAmount float64 `json:"minOrderAmount,string"`
	MaxOrderAmount float64 `json:"maxOrderAmount,string"`
	AmountDecimals int64   `json:"amountDecimals,string"`
	PriceDecimals  int64   `json:"priceDecimals,string"`
}
type Ticker struct {
	MarketID  string    `json:"marketId"`
	BestBID   float64   `json:"bestBid,string"`
	BestAsk   float64   `json:"bestAsk,string"`
	LastPrice float64   `json:"lastPrice,string"`
	Volume    float64   `json:"volume24h,string"`
	Change24h float64   `json:"price24h,string"`
	Low24h    float64   `json:"low24h,string"`
	High24h   float64   `json:"high24h,string"`
	Timestamp time.Time `json:"timestamp"`
}
type Trade struct {
	TradeID   string    `json:"id"`
	Amount    float64   `json:"amount,string"`
	Price     float64   `json:"price,string"`
	Timestamp time.Time `json:"timestamp"`
}
type tempOrderbook struct {
	MarketID   string      `json:"marketId"`
	SnapshotID int64       `json:"snapshotId"`
	Asks       [][2]string `json:"asks"`
	Bids       [][2]string `json:"bids"`
}
type OBData struct {
	Price  float64
	Volume float64
}
type Orderbook struct {
	MarketID   string
	SnapshotID int64
	Asks       []OBData
	Bids       []OBData
}
type MarketCandle struct {
	Time   time.Time
	Open   float64
	Close  float64
	Low    float64
	High   float64
	Volume float64
}
type TimeResp struct {
	Time time.Time `json:"timestamp"`
}
type TradingFee struct {
	Success        bool    `json:"success"`
	ErrorCode      int     `json:"errorCode"`
	ErrorMessage   string  `json:"errorMessage"`
	TradingFeeRate float64 `json:"tradingfeerate"`
	Volume30Day    float64 `json:"volume30day"`
}
type OrderToGo struct {
	Currency        string `json:"currency"`
	Instrument      string `json:"instrument"`
	Price           int64  `json:"price"`
	Volume          int64  `json:"volume"`
	OrderSide       string `json:"orderSide"`
	OrderType       string `json:"ordertype"`
	ClientRequestID string `json:"clientRequestId"`
}
type Order struct {
	ID              int64           `json:"id"`
	Currency        string          `json:"currency"`
	Instrument      string          `json:"instrument"`
	OrderSide       string          `json:"orderSide"`
	OrderType       string          `json:"ordertype"`
	CreationTime    time.Time       `json:"creationTime"`
	Status          string          `json:"status"`
	ErrorMessage    string          `json:"errorMessage"`
	Price           float64         `json:"price"`
	Volume          float64         `json:"volume"`
	OpenVolume      float64         `json:"openVolume"`
	ClientRequestID string          `json:"clientRequestId"`
	Trades          []TradeResponse `json:"trades"`
}
type TradeResponse struct {
	ID           int64     `json:"id"`
	CreationTime time.Time `json:"creationTime"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Volume       float64   `json:"volume"`
	Fee          float64   `json:"fee"`
}
type AccountData struct {
	AssetName string  `json:"assetName"`
	Balance   float64 `json:"balance,string"`
	Available float64 `json:"available,string"`
	Locked    float64 `json:"locked,string"`
}
type TradeHistoryData struct {
	ID            string    `json:"id"`
	MarketID      string    `json:"marketId"`
	Timestamp     time.Time `json:"timestamp"`
	Price         float64   `json:"price,string"`
	Amount        float64   `json:"amount,string"`
	Side          string    `json:"side"`
	Fee           float64   `json:"fee,string"`
	OrderID       string    `json:"orderId"`
	LiquidityType string    `json:"liquidityType"`
}
type OrderData struct {
	OrderID      string    `json:"orderId"`
	MarketID     string    `json:"marketId"`
	Side         string    `json:"side"`
	Type         string    `json:"type"`
	CreationTime time.Time `json:"creationTime"`
	Price        float64   `json:"price,string"`
	Amount       float64   `json:"amount,string"`
	OpenAmount   float64   `json:"openAmount,string"`
	Status       string    `json:"status"`
}
type CancelOrderResp struct {
	OrderID       string `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}
type PaymentDetails struct {
	Address string `json:"address"`
}
type TransferData struct {
	ID             string         `json:"id"`
	AssetName      string         `json:"assetName"`
	Amount         float64        `json:"amount,string"`
	RequestType    string         `json:"type"`
	CreationTime   time.Time      `json:"creationTime"`
	Status         string         `json:"status"`
	Description    string         `json:"description"`
	Fee            float64        `json:"fee,string"`
	LastUpdate     string         `json:"lastUpdate"`
	PaymentDetails PaymentDetails `json:"paymentDetail"`
}
type DepositAddress struct {
	Address   string `json:"address"`
	AssetName string `json:"assetName"`
}
type WithdrawalFeeData struct {
	AssetName string  `json:"assetName"`
	Fee       float64 `json:"fee,string"`
}
type AssetData struct {
	AssetName           string  `json:"assetName"`
	MinDepositAmount    float64 `json:"minDepositAmount,string"`
	MaxDepositAmount    float64 `json:"maxDepositAmount,string"`
	DepositDecimals     float64 `json:"depositDecimals,string"`
	MinWithdrawalAmount float64 `json:"minWithdrawalAmount,string"`
	MaxWithdrawalAmount float64 `json:"maxWithdrawalAmount,string"`
	WithdrawalDecimals  float64 `json:"withdrawalDecimals,string"`
	WithdrawalFee       float64 `json:"withdrawalFee,string"`
	DepositFee          float64 `json:"depositFee,string"`
}
type TransactionData struct {
	ID           string    `json:"id"`
	CreationTime time.Time `json:"creationTime"`
	Description  string    `json:"description"`
	AssetName    string    `json:"assetName"`
	Amount       float64   `json:"amount,string"`
	Balance      float64   `json:"balance,string"`
	FeeType      string    `json:"type"`
	RecordType   string    `json:"recordType"`
	ReferrenceID string    `json:"referrenceId"`
}
type CreateReportResp struct {
	ReportID string `json:"reportId"`
}
type ReportData struct {
	ID           string    `json:"id"`
	ContentURL   string    `json:"contentUrl"`
	CreationTime time.Time `json:"creationTime"`
	ReportType   string    `json:"reportType"`
	Status       string    `json:"status"`
	Format       string    `json:"format"`
}
type BatchPlaceData struct {
	OrderID       string    `json:"orderId"`
	MarketID      string    `json:"marketId"`
	Side          string    `json:"side"`
	Type          string    `json:"type"`
	CreationTime  time.Time `json:"creationTime"`
	Price         float64   `json:"price,string"`
	Amount        float64   `json:"amount,string"`
	OpenAmount    float64   `json:"openAmount,string"`
	Status        string    `json:"status"`
	ClientOrderID string    `json:"clientOrderId"`
}
type UnprocessedBatchResp struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}
type BatchPlaceCancelResponse struct {
	PlacedOrders      []BatchPlaceData       `json:"placeOrders"`
	CancelledOrders   []CancelOrderResp      `json:"cancelOrders"`
	UnprocessedOrders []UnprocessedBatchResp `json:"unprocessedRequests"`
}
type BatchTradeResponse struct {
	Orders              []BatchPlaceData       `json:"orders"`
	UnprocessedRequests []UnprocessedBatchResp `json:"unprocessedRequests"`
}
type BatchCancelResponse struct {
	CancelOrders        []CancelOrderResp      `json:"cancelOrders"`
	UnprocessedRequests []UnprocessedBatchResp `json:"unprocessedRequests"`
}
type WithdrawRequestCrypto struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Address  string `json:"address"`
}
type WithdrawRequestAUD struct {
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	BankName      string `json:"bankName"`
	BSBNumber     string `json:"bsbNumber"`
}
type CancelBatch struct {
	OrderID       string `json:"orderId,omitempty"`
	ClientOrderID string `json:"clientOrderId,omitempty"`
}
type PlaceBatch struct {
	MarketID      string  `json:"marketId"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`
	OrderType     string  `json:"type"`
	Side          string  `json:"side"`
	TriggerPrice  float64 `json:"triggerPrice,omitempty"`
	TriggerAmount float64 `json:"triggerAmount,omitempty"`
	TimeInForce   string  `json:"timeInForce,omitempty"`
	PostOnly      bool    `json:"postOnly,omitempty"`
	SelfTrade     string  `json:"selfTrade,omitempty"`
	ClientOrderID string  `json:"clientOrderId,omitempty"`
}
type PlaceOrderMethod struct {
	PlaceOrder PlaceBatch `json:"placeOrder,omitempty"`
}
type CancelOrderMethod struct {
	CancelOrder CancelBatch `json:"cancelOrder,omitempty"`
}
type TradingFeeData struct {
	MakerFeeRate float64 `json:"makerFeeRate,string"`
	TakerFeeRate float64 `json:"takerFeeRate,string"`
	MarketID     string  `json:"marketId"`
}
type TradingFeeResponse struct {
	MonthlyVolume float64          `json:"volume30Day,string"`
	FeeByMarkets  []TradingFeeData `json:"FeeByMarkets"`
}
type WsSubscribe struct {
	MarketIDs   []string `json:"marketIds,omitempty"`
	Channels    []string `json:"channels"`
	MessageType string   `json:"messageType"`
}
type WsAuthSubscribe struct {
	MarketIDs   []string `json:"marketIds,omitempty"`
	Channels    []string `json:"channels"`
	Key         string   `json:"key"`
	Signature   string   `json:"signature"`
	Timestamp   string   `json:"timestamp"`
	MessageType string   `json:"messageType"`
}
type WsMessageType struct {
	MessageType string `json:"messageType"`
}
type WsTick struct {
	Currency    string    `json:"marketId"`
	Timestamp   time.Time `json:"timestamp"`
	Bid         float64   `json:"bestBid,string"`
	Ask         float64   `json:"bestAsk,string"`
	Last        float64   `json:"lastPrice,string"`
	Volume      float64   `json:"volume24h,string"`
	Price24h    float64   `json:"price24h,string"`
	Low24h      float64   `json:"low24h,string"`
	High24      float64   `json:"high24h,string"`
	MessageType string    `json:"messageType"`
}
type WsTrade struct {
	Currency    string    `json:"marketId"`
	Timestamp   time.Time `json:"timestamp"`
	TradeID     int64     `json:"tradeId"`
	Price       float64   `json:"price,string"`
	Volume      float64   `json:"volume,string"`
	MessageType string    `json:"messageType"`
}
type WsOrderbook struct {
	Currency    string          `json:"marketId"`
	Timestamp   time.Time       `json:"timestamp"`
	Bids        [][]interface{} `json:"bids"`
	Asks        [][]interface{} `json:"asks"`
	MessageType string          `json:"messageType"`
	Snapshot    bool            `json:"snapshot"`
}
type WsFundTransfer struct {
	FundTransferID int64     `json:"fundtransferId"`
	TransferType   string    `json:"type"`
	Status         string    `json:"status"`
	Timestamp      time.Time `json:"timestamp"`
	Amount         float64   `json:"amount,string"`
	Currency       string    `json:"currency"`
	Fee            float64   `json:"fee,string"`
	MessageType    string    `json:"messageType"`
}
type WsTradeData struct {
	TradeID       int64   `json:"tradeId"`
	Price         float64 `json:"price,string"`
	Volume        float64 `json:"volume,string"`
	Fee           float64 `json:"fee,string"`
	LiquidityType string  `json:"liquidityType"`
}
type WsOrderChange struct {
	OrderID       int64         `json:"orderId"`
	MarketID      string        `json:"marketId"`
	Side          string        `json:"side"`
	OrderType     string        `json:"type"`
	OpenVolume    float64       `json:"openVolume,string"`
	Status        string        `json:"status"`
	TriggerStatus string        `json:"triggerStatus"`
	Trades        []WsTradeData `json:"trades"`
	Timestamp     time.Time     `json:"timestamp"`
	MessageType   string        `json:"messageType"`
}
type WsError struct {
	MessageType string `json:"messageType"`
	Code        int64  `json:"code"`
	Message     string `json:"message"`
}
