package coinbene

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
)

type TickerData struct {
	Symbol      string  `json:"symbol"`
	LatestPrice float64 `json:"latestPrice,string"`
	BestBid     float64 `json:"bestBid,string"`
	BestAsk     float64 `json:"bestAsk,string"`
	DailyHigh   float64 `json:"high24h,string"`
	DailyLow    float64 `json:"low24h,string"`
	DailyVolume float64 `json:"volume24h,string"`
}
type OrderbookItem struct {
	Price  float64
	Amount float64
	Count  int64
}
type Orderbook struct {
	Bids   []OrderbookItem
	Asks   []OrderbookItem
	Symbol string
	Time   time.Time
}
type TradeItem struct {
	CurrencyPair string
	Price        float64
	Volume       float64
	Direction    string
	TradeTime    time.Time
}
type PairData struct {
	Symbol           string  `json:"symbol"`
	BaseAsset        string  `json:"baseAsset"`
	QuoteAsset       string  `json:"quoteAsset"`
	PricePrecision   int64   `json:"pricePrecision,string"`
	AmountPrecision  int64   `json:"amountPrecision,string"`
	TakerFeeRate     float64 `json:"takerFeeRate,string"`
	MakerFeeRate     float64 `json:"makerFeeRate,string"`
	MinAmount        float64 `json:"minAmount,string"`
	Site             string  `json:"site"`
	PriceFluctuation float64 `json:"priceFluctuation,string"`
}
type UserBalanceData struct {
	Asset     string  `json:"asset"`
	Available float64 `json:"available,string"`
	Reserved  float64 `json:"reserved,string"`
	Total     float64 `json:"total,string"`
}
type PlaceOrderRequest struct {
	Price     float64
	Quantity  float64
	Symbol    string
	Direction string
	OrderType string
	ClientID  string
	Notional  int
}
type CancelOrdersResponse struct {
	OrderID string `json:"orderId"`
	Message string `json:"message"`
}
type OrderInfo struct {
	OrderID      string    `json:"orderId"`
	BaseAsset    string    `json:"baseAsset"`
	QuoteAsset   string    `json:"quoteAsset"`
	OrderType    string    `json:"orderDirection"`
	Quantity     float64   `json:"quntity,string"`
	Amount       float64   `json:"amout,string"`
	FilledAmount float64   `json:"filledAmount"`
	TakerRate    float64   `json:"takerFeeRate,string"`
	MakerRate    float64   `json:"makerFeeRate,string"`
	AvgPrice     float64   `json:"avgPrice,string"`
	OrderPrice   float64   `json:"orderPrice,string"`
	OrderStatus  string    `json:"orderStatus"`
	OrderTime    time.Time `json:"orderTime"`
	TotalFee     float64   `json:"totalFee"`
}
type OrderFills struct {
	Price     float64   `json:"price,string"`
	Quantity  float64   `json:"quantity,string"`
	Amount    float64   `json:"amount,string"`
	Fee       float64   `json:"fee,string"`
	Direction string    `json:"direction"`
	TradeTime time.Time `json:"tradeTime"`
	FeeByConi float64   `json:"feeByConi,string"`
}
type WsSub struct {
	Operation string   `json:"op"`
	Arguments []string `json:"args"`
}
type WsTickerData struct {
	Symbol        string    `json:"symbol"`
	LastPrice     float64   `json:"lastPrice,string"`
	MarkPrice     float64   `json:"markPrice,string"`
	BestAskPrice  float64   `json:"bestAskPrice,string"`
	BestBidPrice  float64   `json:"bestBidPrice,string"`
	BestAskVolume float64   `json:"bestAskVolume,string"`
	BestBidVolume float64   `json:"bestBidVolume,string"`
	High24h       float64   `json:"high24h,string"`
	Low24h        float64   `json:"low24h,string"`
	Volume24h     float64   `json:"volume24h,string"`
	Timestamp     time.Time `json:"timestamp"`
}
type WsTicker struct {
	Topic string         `json:"topic"`
	Data  []WsTickerData `json:"data"`
}
type WsTradeList struct {
	Topic string     `json:"topic"`
	Data  [][]string `json:"data"`
}
type WsKline struct {
	Topic string          `json:"topic"`
	Data  [][]interface{} `json:"data"`
}
type WsUserData struct {
	Asset     string    `json:"string"`
	Available float64   `json:"availableBalance,string"`
	Locked    float64   `json:"frozenBalance,string"`
	Total     float64   `json:"balance,string"`
	Timestamp time.Time `json:"timestamp"`
}
type WsUserInfo struct {
	Topic string       `json:"topic"`
	Data  []WsUserData `json:"data"`
}
type WsPositionData struct {
	AvailableQuantity float64   `json:"availableQuantity,string"`
	AveragePrice      float64   `json:"avgPrice,string"`
	Leverage          int64     `json:"leverage,string"`
	LiquidationPrice  float64   `json:"liquidationPrice,string"`
	MarkPrice         float64   `json:"markPrice,string"`
	PositionMargin    float64   `json:"positionMargin,string"`
	Quantity          float64   `json:"quantity,string"`
	RealisedPNL       float64   `json:"realisedPnl,string"`
	Side              string    `json:"side"`
	Symbol            string    `json:"symbol"`
	MarginMode        int64     `json:"marginMode,string"`
	CreateTime        time.Time `json:"createTime"`
}
type WsPosition struct {
	Topic string           `json:"topic"`
	Data  []WsPositionData `json:"data"`
}
type WsOrderData struct {
	OrderID          string    `json:"orderId"`
	Direction        string    `json:"direction"`
	Leverage         int64     `json:"leverage,string"`
	Symbol           string    `json:"symbol"`
	OrderType        string    `json:"orderType"`
	Quantity         float64   `json:"quantity,string"`
	OrderPrice       float64   `json:"orderPrice,string"`
	OrderValue       float64   `json:"orderValue,string"`
	Fee              float64   `json:"fee,string"`
	FilledQuantity   float64   `json:"filledQuantity,string"`
	AveragePrice     float64   `json:"averagePrice,string"`
	OrderTime        time.Time `json:"orderTime"`
	Status           string    `json:"status"`
	LastFillQuantity float64   `json:"lastFillQuantity,string"`
	LastFillPrice    float64   `json:"lastFillPrice,string"`
	LastFillTime     string    `json:"lastFillTime"`
}
type WsUserOrders struct {
	Topic string        `json:"topic"`
	Data  []WsOrderData `json:"data"`
}
type SwapTicker struct {
	LastPrice     float64   `json:"lastPrice,string"`
	MarkPrice     float64   `json:"markPrice,string"`
	BestAskPrice  float64   `json:"bestAskPrice,string"`
	BestBidPrice  float64   `json:"bestBidPrice,string"`
	High24Hour    float64   `json:"high24h,string"`
	Low24Hour     float64   `json:"low24h,string"`
	Volume24Hour  float64   `json:"volume24h,string"`
	BestAskVolume float64   `json:"bestAskVolume,string"`
	BestBidVolume float64   `json:"bestBidVolume,string"`
	Turnover      float64   `json:"turnover,string"`
	Timestamp     time.Time `json:"timeStamp"`
}
type SwapKlineItem struct {
	Time        time.Time
	Open        float64
	Close       float64
	High        float64
	Low         float64
	Volume      float64
	Turnover    float64
	BuyVolume   float64
	BuyTurnover float64
}
type SwapTrade struct {
	Price  float64
	Side   order.Side
	Volume float64
	Time   time.Time
}
type SwapAccountInfo struct {
	AvailableBalance        float64 `json:"availableBalance,string"`
	FrozenBalance           float64 `json:"frozenBalance,string"`
	MarginBalance           float64 `json:"marginBalance,string"`
	MarginRate              float64 `json:"marginRate,string"`
	Balance                 float64 `json:"balance,string"`
	UnrealisedProfitAndLoss float64 `json:"unrealisedPnl,string"`
}
type SwapPosition struct {
	AvailableQuantity       float64   `json:"availableQuantity,string"`
	AveragePrice            float64   `json:"averagePrice,string"`
	CreateTime              time.Time `json:"createTime"`
	DeleveragePercentile    int64     `json:"deleveragePercentile,string"`
	Leverage                int64     `json:"leverage,string"`
	LiquidationPrice        float64   `json:"liquidationPrice,string"`
	MarkPrice               float64   `json:"markPrice,string"`
	PositionMargin          float64   `json:"positionMargin,string"`
	PositionValue           float64   `json:"positionValue,string"`
	Quantity                float64   `json:"quantity,string"`
	RateOfReturn            float64   `json:"roe,string"`
	Side                    string    `json:"side"`
	Symbol                  string    `json:"symbol"`
	UnrealisedProfitAndLoss float64   `json:"UnrealisedPnl,string"`
}
type SwapPlaceOrderResponse struct {
	OrderID  string `json:"orderId"`
	ClientID string `json:"clientId"`
}
type SwapOrder struct {
	OrderID        string    `json:"orderId"`
	Direction      string    `json:"direction"`
	Leverage       int64     `json:"leverage,string"`
	OrderType      string    `json:"orderType"`
	Quantity       float64   `json:"quantity,string"`
	OrderPrice     float64   `json:"orderPrice,string"`
	OrderValue     float64   `json:"orderValue,string"`
	Fee            float64   `json:"fee"`
	FilledQuantity float64   `json:"filledQuantity,string"`
	AveragePrice   float64   `json:"averagePrice,string"`
	OrderTime      time.Time `json:"orderTime"`
	Status         string    `json:"status"`
}
type OrderCancellationResponse struct {
	OrderID string `json:"orderId"`
	Code    int    `json:"code,string"`
	Message string `json:"message"`
}
type SwapOrderFill struct {
	Symbol    string    `json:"symbol"`
	TradeTime time.Time `json:"tradeTime"`
	TradeID   int64     `json:"tradeId,string"`
	OrderID   int64     `json:"orderId,string"`
	Price     float64   `json:"price,string"`
	Fee       float64   `json:"fee,string"`
	ExecType  string    `json:"execType"`
	Side      string    `json:"side"`
	Quantity  float64   `json:"quantity,string"`
}
type SwapFundingRate struct {
	Symbol        string  `json:"symbol"`
	Side          string  `json:"side"`
	MarkPrice     float64 `json:"markPrice,string"`
	PositionValue float64 `json:"positionValue,string"`
	Fee           float64 `json:"fee,string"`
	FeeRate       float64 `json:"feeRate,string"`
	Leverage      int64   `json:"leverage"`
}
type Trades []TradeItem
type OrdersInfo []OrderInfo
type SwapTickers map[string]SwapTicker
type SwapKlines []SwapKlineItem
type SwapTrades []SwapTrade
type SwapPositions []SwapPosition
type SwapOrders []SwapOrder
type OrderPlacementResponse OrderCancellationResponse
type SwapOrderFills []SwapOrderFill
