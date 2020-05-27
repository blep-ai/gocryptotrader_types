package coinut

import (
	"sync"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
)
type GenericResponse struct {
	Nonce         int64    `json:"nonce"`
	Reply         string   `json:"reply"`
	Status        []string `json:"status"`
	TransactionID int64    `json:"trans_id"`
}
type InstrumentBase struct {
	Base          string `json:"base"`
	DecimalPlaces int    `json:"decimal_places"`
	InstrumentID  int64  `json:"inst_id"`
	Quote         string `json:"quote"`
}
type Instruments struct {
	Instruments map[string][]InstrumentBase `json:"SPOT"`
}
type Ticker struct {
	High24                float64 `json:"high24,string"`
	HighestBuy            float64 `json:"highest_buy,string"`
	InstrumentID          int     `json:"inst_id"`
	Last                  float64 `json:"last,string"`
	Low24                 float64 `json:"low24,string"`
	LowestSell            float64 `json:"lowest_sell,string"`
	PreviousTransactionID int64   `json:"prev_trans_id"`
	PriceChange24         float64 `json:"price_change_24,string"`
	Reply                 string  `json:"reply"`
	OpenInterest          float64 `json:"open_interest,string"`
	Timestamp             int64   `json:"timestamp"`
	TransactionID         int64   `json:"trans_id"`
	Volume                float64 `json:"volume,string"`
	Volume24              float64 `json:"volume24,string"`
	Volume24Quote         float64 `json:"volume24_quote,string"`
	VolumeQuote           float64 `json:"volume_quote,string"`
}
type OrderbookBase struct {
	Count    int     `json:"count"`
	Price    float64 `json:"price,string"`
	Quantity float64 `json:"qty,string"`
}
type Orderbook struct {
	Buy           []OrderbookBase `json:"buy"`
	Sell          []OrderbookBase `json:"sell"`
	InstrumentID  int             `json:"inst_id"`
	TotalBuy      float64         `json:"total_buy,string"`
	TotalSell     float64         `json:"total_sell,string"`
	TransactionID int64           `json:"trans_id"`
}
type TradeBase struct {
	Price         float64 `json:"price,string"`
	Quantity      float64 `json:"quantity,string"`
	Side          string  `json:"side"`
	Timestamp     float64 `json:"timestamp"`
	TransactionID int64   `json:"trans_id"`
}
type Trades struct {
	Trades []TradeBase `json:"trades"`
}
type UserBalance struct {
	BCH     float64  `json:"BCH,string"`
	BTC     float64  `json:"BTC,string"`
	BTG     float64  `json:"BTG,string"`
	CAD     float64  `json:"CAD,string"`
	ETC     float64  `json:"ETC,string"`
	ETH     float64  `json:"ETH,string"`
	LCH     float64  `json:"LCH,string"`
	LTC     float64  `json:"LTC,string"`
	MYR     float64  `json:"MYR,string"`
	SGD     float64  `json:"SGD,string"`
	USD     float64  `json:"USD,string"`
	USDT    float64  `json:"USDT,string"`
	XMR     float64  `json:"XMR,string"`
	ZEC     float64  `json:"ZEC,string"`
	Nonce   int64    `json:"nonce"`
	Reply   string   `json:"reply"`
	Status  []string `json:"status"`
	TransID int64    `json:"trans_id"`
}
type Order struct {
	InstrumentID  int64   `json:"inst_id"`
	Price         float64 `json:"price,string"`
	Quantity      float64 `json:"qty,string"`
	ClientOrderID int     `json:"client_ord_id"`
	Side          string  `json:"side,string"`
}
type OrderResponse struct {
	OrderID       int64   `json:"order_id"`
	OpenQuantity  float64 `json:"open_qty,string"`
	Price         float64 `json:"price,string"`
	Quantity      float64 `json:"qty,string"`
	InstrumentID  int64   `json:"inst_id"`
	ClientOrderID int64   `json:"client_ord_id"`
	Timestamp     int64   `json:"timestamp"`
	OrderPrice    float64 `json:"order_price,string"`
	Side          string  `json:"side"`
}
type Commission struct {
	Currency currency.Pair `json:"currency"`
	Amount   float64       `json:"amount,string"`
}
type OrderFilledResponse struct {
	GenericResponse
	Commission   Commission    `json:"commission"`
	FillPrice    float64       `json:"fill_price,string"`
	FillQuantity float64       `json:"fill_qty,string"`
	Order        OrderResponse `json:"order"`
}
type OrderRejectResponse struct {
	OrderResponse
	Reasons []string `json:"reasons"`
}
type OrdersBase struct {
	GenericResponse
	OrderResponse
}
type GetOpenOrdersResponse struct {
	Nonce         int             `json:"nonce"`
	Orders        []OrderResponse `json:"orders"`
	Reply         string          `json:"reply"`
	Status        []string        `json:"status"`
	TransactionID int             `json:"trans_id"`
}
type OrdersResponse struct {
	Data []OrdersBase
}
type CancelOrders struct {
	InstrumentID int64 `json:"inst_id"`
	OrderID      int64 `json:"order_id"`
}
type CancelOrdersResponse struct {
	GenericResponse
	Results []struct {
		OrderID      int64  `json:"order_id"`
		Status       string `json:"status"`
		InstrumentID int64  `json:"inst_id"`
	} `json:"results"`
}
type TradeHistory struct {
	TotalNumber int64                 `json:"total_number"`
	Trades      []OrderFilledResponse `json:"trades"`
}
type IndexTicker struct {
	Asset string  `json:"asset"`
	Price float64 `json:"price,string"`
}
type Option struct {
	HighestBuy   float64 `json:"highest_buy,string"`
	InstrumentID int     `json:"inst_id"`
	Last         float64 `json:"last,string"`
	LowestSell   float64 `json:"lowest_sell,string"`
	OpenInterest float64 `json:"open_interest,string"`
}
type OptionChainResponse struct {
	ExpiryTime   int64  `json:"expiry_time"`
	SecurityType string `json:"sec_type"`
	Asset        string `json:"asset"`
	Entries      []struct {
		Call   Option  `json:"call"`
		Put    Option  `json:"put"`
		Strike float64 `json:"strike,string"`
	}
}
type OptionChainUpdate struct {
	Option
	GenericResponse
	Asset        string  `json:"asset"`
	ExpiryTime   int64   `json:"expiry_time"`
	SecurityType string  `json:"sec_type"`
	Volume       float64 `json:"volume,string"`
}
type PositionHistory struct {
	Positions []struct {
		PositionID int `json:"position_id"`
		Records    []struct {
			Commission    Commission `json:"commission"`
			FillPrice     float64    `json:"fill_price,string,omitempty"`
			TransactionID int        `json:"trans_id"`
			FillQuantity  float64    `json:"fill_qty,omitempty"`
			Position      struct {
				Commission Commission `json:"commission"`
				Timestamp  int64      `json:"timestamp"`
				OpenPrice  float64    `json:"open_price,string"`
				RealizedPL float64    `json:"realized_pl,string"`
				Quantity   float64    `json:"qty,string"`
			} `json:"position"`
			AssetAtExpiry float64 `json:"asset_at_expiry,string,omitempty"`
		} `json:"records"`
		Instrument struct {
			ExpiryTime     int64   `json:"expiry_time"`
			ContractSize   float64 `json:"contract_size,string"`
			ConversionRate float64 `json:"conversion_rate,string"`
			OptionType     string  `json:"option_type"`
			InstrumentID   int     `json:"inst_id"`
			SecType        string  `json:"sec_type"`
			Asset          string  `json:"asset"`
			Strike         float64 `json:"strike,string"`
		} `json:"inst"`
		OpenTimestamp int64 `json:"open_timestamp"`
	} `json:"positions"`
	TotalNumber int `json:"total_number"`
}
type OpenPosition struct {
	PositionID    int        `json:"position_id"`
	Commission    Commission `json:"commission"`
	OpenPrice     float64    `json:"open_price,string"`
	RealizedPL    float64    `json:"realized_pl,string"`
	Quantity      float64    `json:"qty,string"`
	OpenTimestamp int64      `json:"open_timestamp"`
	InstrumentID  int        `json:"inst_id"`
}
type wsRequest struct {
	Request      string `json:"request"`
	SecurityType string `json:"sec_type,omitempty"`
	InstrumentID int64  `json:"inst_id,omitempty"`
	TopN         int64  `json:"top_n,omitempty"`
	Subscribe    bool   `json:"subscribe,omitempty"`
	Nonce        int64  `json:"nonce,omitempty"`
}
type wsResponse struct {
	Nonce int64  `json:"nonce,omitempty"`
	Reply string `json:"reply"`
}
type wsHeartbeatResp struct {
	Nonce  int64         `json:"nonce"`
	Reply  string        `json:"reply"`
	Status []interface{} `json:"status"`
}
type WsTicker struct {
	High24        float64  `json:"high24,string"`
	HighestBuy    float64  `json:"highest_buy,string"`
	InstID        int64    `json:"inst_id"`
	Last          float64  `json:"last,string"`
	Low24         float64  `json:"low24,string"`
	LowestSell    float64  `json:"lowest_sell,string"`
	Nonce         int64    `json:"nonce"`
	PrevTransID   int64    `json:"prev_trans_id"`
	PriceChange24 float64  `json:"price_change_24,string"`
	Reply         string   `json:"reply"`
	Status        []string `json:"status"`
	Timestamp     int64    `json:"timestamp"`
	TransID       int64    `json:"trans_id"`
	Volume        float64  `json:"volume,string"`
	Volume24      float64  `json:"volume24,string"`
	Volume24Quote float64  `json:"volume24_quote,string"`
	VolumeQuote   float64  `json:"volume_quote,string"`
}
type WsOrderbookSnapshot struct {
	Buy       []WsOrderbookData `json:"buy"`
	Sell      []WsOrderbookData `json:"sell"`
	InstID    int64             `json:"inst_id"`
	Nonce     int64             `json:"nonce"`
	TotalBuy  float64           `json:"total_buy,string"`
	TotalSell float64           `json:"total_sell,string"`
	Reply     string            `json:"reply"`
	Status    []interface{}     `json:"status"`
}
type WsOrderbookData struct {
	Count  int64   `json:"count"`
	Price  float64 `json:"price,string"`
	Volume float64 `json:"qty,string"`
}
type WsOrderbookUpdate struct {
	Count    int64   `json:"count"`
	InstID   int64   `json:"inst_id"`
	Price    float64 `json:"price,string"`
	Volume   float64 `json:"qty,string"`
	TotalBuy float64 `json:"total_buy,string"`
	Reply    string  `json:"reply"`
	Side     string  `json:"side"`
	TransID  int64   `json:"trans_id"`
}
type WsTradeSnapshot struct {
	Nonce  int64         `json:"nonce"`
	Reply  string        `json:"reply"`
	Status []interface{} `json:"status"`
	Trades []WsTradeData `json:"trades"`
}
type WsTradeData struct {
	Price     float64 `json:"price,string"`
	Volume    float64 `json:"qty,string"`
	Side      string  `json:"side"`
	Timestamp int64   `json:"timestamp"`
	TransID   int64   `json:"trans_id"`
}
type WsTradeUpdate struct {
	InstID    int64   `json:"inst_id"`
	Price     float64 `json:"price,string"`
	Reply     string  `json:"reply"`
	Side      string  `json:"side"`
	Timestamp int64   `json:"timestamp"`
	TransID   int64   `json:"trans_id"`
}
type WsInstrumentList struct {
	Spot   map[string][]InstrumentBase `json:"SPOT"`
	Nonce  int64                       `json:"nonce,omitempty"`
	Reply  string                      `json:"inst_list,omitempty"`
	Status []interface{}               `json:"status,omitempty"`
}
type WsSupportedCurrency struct {
	Base          string `json:"base"`
	InstID        int64  `json:"inst_id"`
	DecimalPlaces int64  `json:"decimal_places"`
	Quote         string `json:"quote"`
}
type WsRequest struct {
	Request string `json:"request"`
	Nonce   int64  `json:"nonce"`
}
type WsTradeHistoryRequest struct {
	InstID int64 `json:"inst_id"`
	Start  int64 `json:"start,omitempty"`
	Limit  int64 `json:"limit,omitempty"`
	WsRequest
}
type WsCancelOrdersRequest struct {
	Entries []WsCancelOrdersRequestEntry `json:"entries"`
	WsRequest
}
type WsCancelOrdersRequestEntry struct {
	InstID  int64 `json:"inst_id"`
	OrderID int64 `json:"order_id"`
}
type WsCancelOrderParameters struct {
	Currency currency.Pair
	OrderID  int64
}
type WsCancelOrderRequest struct {
	InstrumentID int64 `json:"inst_id"`
	OrderID      int64 `json:"order_id"`
	WsRequest
}
type WsCancelOrderResponse struct {
	Nonce         int64    `json:"nonce"`
	Reply         string   `json:"reply"`
	OrderID       int64    `json:"order_id"`
	ClientOrderID int64    `json:"client_ord_id"`
	Status        []string `json:"status"`
}
type WsCancelOrdersResponse struct {
	Nonce         int64                        `json:"nonce"`
	Reply         string                       `json:"reply"`
	Results       []WsCancelOrdersResponseData `json:"results"`
	Status        []string                     `json:"status"`
	TransactionID int64                        `json:"trans_id"`
}
type WsCancelOrdersResponseData struct {
	InstrumentID int64  `json:"inst_id"`
	OrderID      int64  `json:"order_id"`
	Status       string `json:"status"`
}
type WsGetOpenOrdersRequest struct {
	InstrumentID int64 `json:"inst_id"`
	WsRequest
}
type WsSubmitOrdersRequest struct {
	Orders []WsSubmitOrdersRequestData `json:"orders"`
	WsRequest
}
type WsSubmitOrdersRequestData struct {
	InstrumentID  int64   `json:"inst_id"`
	Price         float64 `json:"price,string"`
	Quantity      float64 `json:"qty,string"`
	ClientOrderID int     `json:"client_ord_id"`
	Side          string  `json:"side"`
}
type WsSubmitOrderRequest struct {
	InstrumentID int64   `json:"inst_id"`
	Price        float64 `json:"price,string"`
	Quantity     float64 `json:"qty,string"`
	OrderID      int64   `json:"client_ord_id"`
	Side         string  `json:"side"`
	WsRequest
}
type WsSubmitOrderParameters struct {
	Currency      currency.Pair
	Side          order.Side
	Amount, Price float64
	OrderID       int64
}
type WsUserBalanceResponse struct {
	Nonce              int64    `json:"nonce"`
	Status             []string `json:"status"`
	Btc                float64  `json:"BTC,string"`
	Ltc                float64  `json:"LTC,string"`
	Etc                float64  `json:"ETC,string"`
	Eth                float64  `json:"ETH,string"`
	FloatingProfitLoss float64  `json:"floating_pl,string"`
	InitialMargin      float64  `json:"initial_margin,string"`
	RealisedProfitLoss float64  `json:"realized_pl,string"`
	MaintenanceMargin  float64  `json:"maintenance_margin,string"`
	Equity             float64  `json:"equity,string"`
	Reply              string   `json:"reply"`
	TransactionID      int64    `json:"trans_id"`
}
type WsOrderAcceptedResponse struct {
	Nonce         int64    `json:"nonce"`
	Status        []string `json:"status"`
	OrderID       int64    `json:"order_id"`
	OpenQuantity  float64  `json:"open_qty,string"`
	InstrumentID  int64    `json:"inst_id"`
	Quantity      float64  `json:"qty,string"`
	ClientOrderID int64    `json:"client_ord_id"`
	OrderPrice    float64  `json:"order_price,string"`
	Reply         string   `json:"reply"`
	Side          string   `json:"side"`
	TransactionID int64    `json:"trans_id"`
}
type WsOrderFilledResponse struct {
	Commission    WsOrderFilledCommissionData `json:"commission"`
	FillPrice     float64                     `json:"fill_price,string"`
	FillQuantity  float64                     `json:"fill_qty,string"`
	Nonce         int64                       `json:"nonce"`
	Order         WsOrderData                 `json:"order"`
	Reply         string                      `json:"reply"`
	Status        []string                    `json:"status"`
	Timestamp     int64                       `json:"timestamp"`
	TransactionID int64                       `json:"trans_id"`
}
type WsOrderData struct {
	ClientOrderID int64    `json:"client_ord_id"`
	InstrumentID  int64    `json:"inst_id"`
	OpenQuantity  float64  `json:"open_qty,string"`
	OrderID       int64    `json:"order_id"`
	Price         float64  `json:"price,string"`
	Quantity      float64  `json:"qty,string"`
	Side          string   `json:"side"`
	Timestamp     int64    `json:"timestamp"`
	Status        []string `json:"status"`
}
type WsOrderFilledCommissionData struct {
	Amount   float64       `json:"amount,string"`
	Currency currency.Pair `json:"currency"`
}
type WsOrderRejectedResponse struct {
	Nonce         int64    `json:"nonce"`
	Status        []string `json:"status"`
	OrderID       int64    `json:"order_id"`
	OpenQuantity  float64  `json:"open_qty,string"`
	Price         float64  `json:"price,string"`
	InstrumentID  int64    `json:"inst_id"`
	Reasons       []string `json:"reasons"`
	ClientOrderID int64    `json:"client_ord_id"`
	Timestamp     int64    `json:"timestamp"`
	Reply         string   `json:"reply"`
	Quantity      float64  `json:"qty,string"`
	Side          string   `json:"side"`
	TransactionID int64    `json:"trans_id"`
}
type wsInstList struct {
	Spot map[string][]struct {
		Base          string `json:"base"`
		DecimalPlaces int64  `json:"decimal_places"`
		InstrumentID  int64  `json:"inst_id"`
		Quote         string `json:"quote"`
	} `json:"spot"`
}
type WsUserOpenOrdersResponse struct {
	Nonce  int64         `json:"nonce"`
	Reply  string        `json:"reply"`
	Status []string      `json:"status"`
	Orders []WsOrderData `json:"orders"`
}
type WsTradeHistoryResponse struct {
	Nonce       int64         `json:"nonce"`
	Reply       string        `json:"reply"`
	Status      []string      `json:"status"`
	TotalNumber int64         `json:"total_number"`
	Trades      []WsOrderData `json:"trades"`
}
type WsTradeHistoryCommissionData struct {
	Amount   float64       `json:"amount,string"`
	Currency currency.Pair `json:"currency"`
}
type WsTradeHistoryTradeData struct {
	Commission    WsTradeHistoryCommissionData `json:"commission"`
	Order         WsOrderData                  `json:"order"`
	FillPrice     float64                      `json:"fill_price,string"`
	FillQuantity  float64                      `json:"fill_qty,string"`
	Timestamp     int64                        `json:"timestamp"`
	TransactionID int64                        `json:"trans_id"`
}
type WsLoginResponse struct {
	APIKey          string   `json:"api_key"`
	Country         string   `json:"country"`
	DepositEnabled  bool     `json:"deposit_enabled"`
	Deposited       bool     `json:"deposited"`
	Email           string   `json:"email"`
	FailedTimes     int64    `json:"failed_times"`
	KycPassed       bool     `json:"kyc_passed"`
	Language        string   `json:"lang"`
	Nonce           int64    `json:"nonce"`
	OTPEnabled      bool     `json:"otp_enabled"`
	PhoneNumber     string   `json:"phone_number"`
	ProductsEnabled []string `json:"products_enabled"`
	Referred        bool     `json:"referred"`
	Reply           string   `json:"reply"`
	SessionID       string   `json:"session_id"`
	Status          []string `json:"status"`
	Timezone        string   `json:"timezone"`
	Traded          bool     `json:"traded"`
	UnverifiedEmail string   `json:"unverified_email"`
	Username        string   `json:"username"`
	WithdrawEnabled bool     `json:"withdraw_enabled"`
}
type WsNewOrderResponse struct {
	Message string   `json:"msg"`
	Nonce   int64    `json:"nonce"`
	Reply   string   `json:"reply"`
	Status  []string `json:"status"`
}
type WsGetAccountBalanceResponse struct {
	BCH     float64  `json:"BCH,string"`
	BTC     float64  `json:"BTC,string"`
	BTG     float64  `json:"BTG,string"`
	CAD     float64  `json:"CAD,string"`
	ETC     float64  `json:"ETC,string"`
	ETH     float64  `json:"ETH,string"`
	LCH     float64  `json:"LCH,string"`
	LTC     float64  `json:"LTC,string"`
	MYR     float64  `json:"MYR,string"`
	SGD     float64  `json:"SGD,string"`
	USD     float64  `json:"USD,string"`
	USDT    float64  `json:"USDT,string"`
	XMR     float64  `json:"XMR,string"`
	ZEC     float64  `json:"ZEC,string"`
	Nonce   int64    `json:"nonce"`
	Reply   string   `json:"reply"`
	Status  []string `json:"status"`
	TransID int64    `json:"trans_id"`
}
type instrumentMap struct {
	Instruments map[string]int64
	Loaded      bool
	m           sync.Mutex
}
type wsOrderContainer struct {
	OrderID       int64    `json:"order_id"`
	ClientOrderID int64    `json:"client_ord_id"`
	InstrumentID  int64    `json:"inst_id"`
	Nonce         int64    `json:"nonce"`
	Timestamp     int64    `json:"timestamp"`
	TransactionID int64    `json:"trans_id"`
	OpenQuantity  float64  `json:"open_qty,string"`
	OrderPrice    float64  `json:"order_price,string"`
	Quantity      float64  `json:"qty,string"`
	FillPrice     float64  `json:"fill_price,string"`
	FillQuantity  float64  `json:"fill_qty,string"`
	Price         float64  `json:"price,string"`
	Reply         string   `json:"reply"`
	Side          string   `json:"side"`
	Status        []string `json:"status"`
	Reasons       []string `json:"reasons"`
	Order         struct {
		ClientOrderID int64   `json:"client_ord_id"`
		InstrumentID  int64   `json:"inst_id"`
		OrderID       int64   `json:"order_id"`
		Timestamp     int64   `json:"timestamp"`
		Price         float64 `json:"price,string"`
		Quantity      float64 `json:"qty,string"`
		OpenQuantity  float64 `json:"open_qty,string"`
		Side          string  `json:"side"`
	} `json:"order"`
	Commission struct {
		Amount   float64 `json:"amount,string"`
		Currency string  `json:"currency"`
	} `json:"commission"`
}
