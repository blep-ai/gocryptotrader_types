package binance

import (
	"time"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type ExchangeInfo struct {
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	Timezone   string `json:"timezone"`
	Servertime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	ExchangeFilters interface{} `json:"exchangeFilters"`
	Symbols         []struct {
		Symbol             string   `json:"symbol"`
		Status             string   `json:"status"`
		BaseAsset          string   `json:"baseAsset"`
		BaseAssetPrecision int      `json:"baseAssetPrecision"`
		QuoteAsset         string   `json:"quoteAsset"`
		QuotePrecision     int      `json:"quotePrecision"`
		OrderTypes         []string `json:"orderTypes"`
		IcebergAllowed     bool     `json:"icebergAllowed"`
		Filters            []struct {
			FilterType          string  `json:"filterType"`
			MinPrice            float64 `json:"minPrice,string"`
			MaxPrice            float64 `json:"maxPrice,string"`
			TickSize            float64 `json:"tickSize,string"`
			MultiplierUp        float64 `json:"multiplierUp,string"`
			MultiplierDown      float64 `json:"multiplierDown,string"`
			AvgPriceMins        int64   `json:"avgPriceMins"`
			MinQty              float64 `json:"minQty,string"`
			MaxQty              float64 `json:"maxQty,string"`
			StepSize            float64 `json:"stepSize,string"`
			MinNotional         float64 `json:"minNotional,string"`
			ApplyToMarket       bool    `json:"applyToMarket"`
			Limit               int64   `json:"limit"`
			MaxNumAlgoOrders    int64   `json:"maxNumAlgoOrders"`
			MaxNumIcebergOrders int64   `json:"maxNumIcebergOrders"`
		} `json:"filters"`
	} `json:"symbols"`
}
type OrderBookDataRequestParams struct {
	Symbol string `json:"symbol"` // Required field; example LTCBTC,BTCUSDT
	Limit  int    `json:"limit"`  // Default 100; max 1000. Valid limits:[5, 10, 20, 50, 100, 500, 1000]
}
type OrderbookItem struct {
	Price    float64
	Quantity float64
}
type OrderBookData struct {
	Code         int         `json:"code"`
	Msg          string      `json:"msg"`
	LastUpdateID int64       `json:"lastUpdateId"`
	Bids         [][2]string `json:"bids"`
	Asks         [][2]string `json:"asks"`
}
type OrderBook struct {
	LastUpdateID int64
	Code         int
	Msg          string
	Bids         []OrderbookItem
	Asks         []OrderbookItem
}
type WebsocketDepthStream struct {
	Event         string          `json:"e"`
	Timestamp     int64           `json:"E"`
	Pair          string          `json:"s"`
	FirstUpdateID int64           `json:"U"`
	LastUpdateID  int64           `json:"u"`
	UpdateBids    [][]interface{} `json:"b"`
	UpdateAsks    [][]interface{} `json:"a"`
}
type RecentTradeRequestParams struct {
	Symbol string `json:"symbol"` // Required field. example LTCBTC, BTCUSDT
	Limit  int    `json:"limit"`  // Default 500; max 500.
}
type RecentTrade struct {
	ID           int64   `json:"id"`
	Price        float64 `json:"price,string"`
	Quantity     float64 `json:"qty,string"`
	Time         float64 `json:"time"`
	IsBuyerMaker bool    `json:"isBuyerMaker"`
	IsBestMatch  bool    `json:"isBestMatch"`
}
type TradeStream struct {
	EventType      string `json:"e"`
	EventTime      int64  `json:"E"`
	Symbol         string `json:"s"`
	TradeID        int64  `json:"t"`
	Price          string `json:"p"`
	Quantity       string `json:"q"`
	BuyerOrderID   int64  `json:"b"`
	SellerOrderID  int64  `json:"a"`
	TimeStamp      int64  `json:"T"`
	Maker          bool   `json:"m"`
	BestMatchPrice bool   `json:"M"`
}
type KlineStream struct {
	EventType string `json:"e"`
	EventTime int64  `json:"E"`
	Symbol    string `json:"s"`
	Kline     struct {
		StartTime                int64   `json:"t"`
		CloseTime                int64   `json:"T"`
		Symbol                   string  `json:"s"`
		Interval                 string  `json:"i"`
		FirstTradeID             int64   `json:"f"`
		LastTradeID              int64   `json:"L"`
		OpenPrice                float64 `json:"o,string"`
		ClosePrice               float64 `json:"c,string"`
		HighPrice                float64 `json:"h,string"`
		LowPrice                 float64 `json:"l,string"`
		Volume                   float64 `json:"v,string"`
		NumberOfTrades           int64   `json:"n"`
		KlineClosed              bool    `json:"x"`
		Quote                    float64 `json:"q,string"`
		TakerBuyBaseAssetVolume  float64 `json:"V,string"`
		TakerBuyQuoteAssetVolume float64 `json:"Q,string"`
	} `json:"k"`
}
type TickerStream struct {
	EventType              string  `json:"e"`
	EventTime              int64   `json:"E"`
	Symbol                 string  `json:"s"`
	PriceChange            float64 `json:"p,string"`
	PriceChangePercent     float64 `json:"P,string"`
	WeightedAvgPrice       float64 `json:"w,string"`
	ClosePrice             float64 `json:"x,string"`
	LastPrice              float64 `json:"c,string"`
	LastPriceQuantity      float64 `json:"Q,string"`
	BestBidPrice           float64 `json:"b,string"`
	BestBidQuantity        float64 `json:"B,string"`
	BestAskPrice           float64 `json:"a,string"`
	BestAskQuantity        float64 `json:"A,string"`
	OpenPrice              float64 `json:"o,string"`
	HighPrice              float64 `json:"h,string"`
	LowPrice               float64 `json:"l,string"`
	TotalTradedVolume      float64 `json:"v,string"`
	TotalTradedQuoteVolume float64 `json:"q,string"`
	OpenTime               int64   `json:"O"`
	CloseTime              int64   `json:"C"`
	FirstTradeID           int64   `json:"F"`
	LastTradeID            int64   `json:"L"`
	NumberOfTrades         int64   `json:"n"`
}
type HistoricalTrade struct {
	Code         int     `json:"code"`
	Msg          string  `json:"msg"`
	ID           int64   `json:"id"`
	Price        float64 `json:"price,string"`
	Quantity     float64 `json:"qty,string"`
	Time         int64   `json:"time"`
	IsBuyerMaker bool    `json:"isBuyerMaker"`
	IsBestMatch  bool    `json:"isBestMatch"`
}
type AggregatedTrade struct {
	ATradeID       int64   `json:"a"`
	Price          float64 `json:"p,string"`
	Quantity       float64 `json:"q,string"`
	FirstTradeID   int64   `json:"f"`
	LastTradeID    int64   `json:"l"`
	TimeStamp      int64   `json:"T"`
	Maker          bool    `json:"m"`
	BestMatchPrice bool    `json:"M"`
}
type CandleStick struct {
	OpenTime                 time.Time
	Open                     float64
	High                     float64
	Low                      float64
	Close                    float64
	Volume                   float64
	CloseTime                time.Time
	QuoteAssetVolume         float64
	TradeCount               float64
	TakerBuyAssetVolume      float64
	TakerBuyQuoteAssetVolume float64
}
type AveragePrice struct {
	Mins  int64   `json:"mins"`
	Price float64 `json:"price,string"`
}
type PriceChangeStats struct {
	Symbol             string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange,string"`
	PriceChangePercent float64 `json:"priceChangePercent,string"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice,string"`
	PrevClosePrice     float64 `json:"prevClosePrice,string"`
	LastPrice          float64 `json:"lastPrice,string"`
	LastQty            float64 `json:"lastQty,string"`
	BidPrice           float64 `json:"bidPrice,string"`
	AskPrice           float64 `json:"askPrice,string"`
	OpenPrice          float64 `json:"openPrice,string"`
	HighPrice          float64 `json:"highPrice,string"`
	LowPrice           float64 `json:"lowPrice,string"`
	Volume             float64 `json:"volume,string"`
	QuoteVolume        float64 `json:"quoteVolume,string"`
	OpenTime           int64   `json:"openTime"`
	CloseTime          int64   `json:"closeTime"`
	FirstID            int64   `json:"firstId"`
	LastID             int64   `json:"lastId"`
	Count              int64   `json:"count"`
}
type SymbolPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}
type BestPrice struct {
	Symbol   string  `json:"symbol"`
	BidPrice float64 `json:"bidPrice,string"`
	BidQty   float64 `json:"bidQty,string"`
	AskPrice float64 `json:"askPrice,string"`
	AskQty   float64 `json:"askQty,string"`
}
type NewOrderRequest struct {
	// Symbol (currency pair to trade)
	Symbol string
	// Side Buy or Sell
	Side string
	// TradeType (market or limit order)
	TradeType RequestParamsOrderType
	// TimeInForce specifies how long the order remains in effect.
	// Examples are (Good Till Cancel (GTC), Immediate or Cancel (IOC) and Fill Or Kill (FOK))
	TimeInForce RequestParamsTimeForceType
	// Quantity is the total base qty spent or received in an order.
	Quantity float64
	// QuoteOrderQty is the total quote qty spent or received in a MARKET order.
	QuoteOrderQty    float64
	Price            float64
	NewClientOrderID string
	StopPrice        float64 // Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
	IcebergQty       float64 // Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
	NewOrderRespType string
}
type NewOrderResponse struct {
	Code            int     `json:"code"`
	Msg             string  `json:"msg"`
	Symbol          string  `json:"symbol"`
	OrderID         int64   `json:"orderId"`
	ClientOrderID   string  `json:"clientOrderId"`
	TransactionTime int64   `json:"transactTime"`
	Price           float64 `json:"price,string"`
	OrigQty         float64 `json:"origQty,string"`
	ExecutedQty     float64 `json:"executedQty,string"`
	// The cumulative amount of the quote that has been spent (with a BUY order) or received (with a SELL order).
	CumulativeQuoteQty float64 `json:"cummulativeQuoteQty,string"`
	Status             string  `json:"status"`
	TimeInForce        string  `json:"timeInForce"`
	Type               string  `json:"type"`
	Side               string  `json:"side"`
	Fills              []struct {
		Price           float64 `json:"price,string"`
		Qty             float64 `json:"qty,string"`
		Commission      float64 `json:"commission,string"`
		CommissionAsset string  `json:"commissionAsset"`
	} `json:"fills"`
}
type CancelOrderResponse struct {
	Symbol            string `json:"symbol"`
	OrigClientOrderID string `json:"origClientOrderId"`
	OrderID           int64  `json:"orderId"`
	ClientOrderID     string `json:"clientOrderId"`
}
type QueryOrderData struct {
	Code          int     `json:"code"`
	Msg           string  `json:"msg"`
	Symbol        string  `json:"symbol"`
	OrderID       int64   `json:"orderId"`
	ClientOrderID string  `json:"clientOrderId"`
	Price         float64 `json:"price,string"`
	OrigQty       float64 `json:"origQty,string"`
	ExecutedQty   float64 `json:"executedQty,string"`
	Status        string  `json:"status"`
	TimeInForce   string  `json:"timeInForce"`
	Type          string  `json:"type"`
	Side          string  `json:"side"`
	StopPrice     float64 `json:"stopPrice,string"`
	IcebergQty    float64 `json:"icebergQty,string"`
	Time          float64 `json:"time"`
	IsWorking     bool    `json:"isWorking"`
}
type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
type Account struct {
	MakerCommission  int       `json:"makerCommission"`
	TakerCommission  int       `json:"takerCommission"`
	BuyerCommission  int       `json:"buyerCommission"`
	SellerCommission int       `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	UpdateTime       int64     `json:"updateTime"`
	Balances         []Balance `json:"balances"`
}
type KlinesRequestParams struct {
	Symbol    string       // Required field; example LTCBTC, BTCUSDT
	Interval  TimeInterval // Time interval period
	Limit     int          // Default 500; max 500.
	StartTime int64
	EndTime   int64
}
type WithdrawResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	ID      string `json:"id"`
}
type UserAccountStream struct {
	ListenKey string `json:"listenKey"`
}
type wsAccountInfo struct {
	CanDeposit       bool    `json:"D"`
	CanTrade         bool    `json:"T"`
	CanWithdraw      bool    `json:"W"`
	EventTime        int64   `json:"E"`
	LastUpdated      int64   `json:"u"`
	BuyerCommission  float64 `json:"b"`
	MakerCommission  float64 `json:"m"`
	SellerCommission float64 `json:"s"`
	TakerCommission  float64 `json:"t"`
	EventType        string  `json:"e"`
	Currencies       []struct {
		Asset     string  `json:"a"`
		Available float64 `json:"f,string"`
		Locked    float64 `json:"l,string"`
	} `json:"B"`
}
type wsAccountPosition struct {
	Currencies []struct {
		Asset     string  `json:"a"`
		Available float64 `json:"f,string"`
		Locked    float64 `json:"l,string"`
	} `json:"B"`
	EventTime   int64  `json:"E"`
	LastUpdated int64  `json:"u"`
	EventType   string `json:"e"`
}
type wsBalanceUpdate struct {
	EventTime    int64   `json:"E"`
	ClearTime    int64   `json:"T"`
	BalanceDelta float64 `json:"d,string"`
	Asset        string  `json:"a"`
	EventType    string  `json:"e"`
}
type wsOrderUpdate struct {
	ClientOrderID                     string  `json:"C"`
	EventTime                         int64   `json:"E"`
	IcebergQuantity                   float64 `json:"F,string"`
	LastExecutedPrice                 float64 `json:"L,string"`
	CommissionAsset                   float64 `json:"N"`
	OrderCreationTime                 int64   `json:"O"`
	StopPrice                         float64 `json:"P,string"`
	QuoteOrderQuantity                float64 `json:"Q,string"`
	Side                              string  `json:"S"`
	TransactionTime                   int64   `json:"T"`
	OrderStatus                       string  `json:"X"`
	LastQuoteAssetTransactedQuantity  float64 `json:"Y,string"`
	CumulativeQuoteTransactedQuantity float64 `json:"Z,string"`
	CancelledClientOrderID            string  `json:"c"`
	EventType                         string  `json:"e"`
	TimeInForce                       string  `json:"f"`
	OrderListID                       int64   `json:"g"`
	OrderID                           int64   `json:"i"`
	LastExecutedQuantity              float64 `json:"l,string"`
	IsMaker                           bool    `json:"m"`
	Commission                        float64 `json:"n,string"`
	OrderType                         string  `json:"o"`
	Price                             float64 `json:"p,string"`
	Quantity                          float64 `json:"q,string"`
	RejectionReason                   string  `json:"r"`
	Symbol                            string  `json:"s"`
	TradeID                           int64   `json:"t"`
	IsOnOrderBook                     bool    `json:"w"`
	CurrentExecutionType              string  `json:"x"`
	CumulativeFilledQuantity          float64 `json:"z,string"`
}
type wsListStauts struct {
	ListClientOrderID string `json:"C"`
	EventTime         int64  `json:"E"`
	ListOrderStatus   string `json:"L"`
	Orders            []struct {
		ClientOrderID string `json:"c"`
		OrderID       int64  `json:"i"`
		Symbol        string `json:"s"`
	} `json:"O"`
	TransactionTime int64  `json:"T"`
	ContingencyType string `json:"c"`
	EventType       string `json:"e"`
	OrderListID     int64  `json:"g"`
	ListStatusType  string `json:"l"`
	RejectionReason string `json:"r"`
	Symbol          string `json:"s"`
}
