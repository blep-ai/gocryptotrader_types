package yobit
import "github.com/blep-ai/gocryptotrader_types/currency"
type Response struct {
	Return  interface{} `json:"return"`
	Success int         `json:"success"`
	Error   string      `json:"error"`
}
type Info struct {
	ServerTime int64           `json:"server_time"`
	Pairs      map[string]Pair `json:"pairs"`
}
type Ticker struct {
	High          float64 // maximal price
	Low           float64 // minimal price
	Avg           float64 // average price
	Vol           float64 // traded volume
	VolumeCurrent float64 `json:"vol_cur"` // traded volume in currency
	Last          float64 // last transaction price
	Buy           float64 // buying price
	Sell          float64 // selling price
	Updated       int64   // last cache upgrade
}
type Orderbook struct {
	Asks [][]float64 `json:"asks"` // selling orders
	Bids [][]float64 `json:"bids"` // buying orders
}
type Trades struct {
	Type      string  `json:"type"`
	Price     float64 `json:"bid"`
	Amount    float64 `json:"amount"`
	TID       int64   `json:"tid"`
	Timestamp int64   `json:"timestamp"`
}
type ActiveOrders struct {
	Pair             string  `json:"pair"`
	Type             string  `json:"type"`
	Amount           float64 `json:"amount"`
	Rate             float64 `json:"rate"`
	TimestampCreated float64 `json:"timestamp_created"`
	Status           int     `json:"status"`
}
type Pair struct {
	DecimalPlaces int     `json:"decimal_places"` // Quantity of permitted numbers after decimal point
	MinPrice      float64 `json:"min_price"`      // Minimal permitted price
	MaxPrice      float64 `json:"max_price"`      // Maximal permitted price
	MinAmount     float64 `json:"min_amount"`     // Minimal permitted buy or sell amount
	Hidden        int     `json:"hidden"`         // Pair is hidden (0 or 1)
	Fee           float64 `json:"fee"`            // Pair commission
}
type AccountInfo struct {
	Funds           map[string]float64 `json:"funds"`
	FundsInclOrders map[string]float64 `json:"funds_incl_orders"`
	Rights          struct {
		Info     int `json:"info"`
		Trade    int `json:"trade"`
		Withdraw int `json:"withdraw"`
	} `json:"rights"`
	TransactionCount int     `json:"transaction_count"`
	OpenOrders       int     `json:"open_orders"`
	ServerTime       float64 `json:"server_time"`
	Error            string  `json:"error"`
}
type OrderInfo struct {
	Pair             string  `json:"pair"`
	Type             string  `json:"type"`
	StartAmount      float64 `json:"start_amount"`
	Amount           float64 `json:"amount"`
	Rate             float64 `json:"rate"`
	TimestampCreated float64 `json:"timestamp_created"`
	Status           int     `json:"status"`
}
type CancelOrder struct {
	OrderID float64            `json:"order_id"`
	Funds   map[string]float64 `json:"funds"`
	Error   string             `json:"error"`
}
type Trade struct {
	Received float64            `json:"received"`
	Remains  float64            `json:"remains"`
	OrderID  float64            `json:"order_id"`
	Funds    map[string]float64 `json:"funds"`
	Error    string             `json:"error"`
}
type TradeHistoryResponse struct {
	Success int64                   `json:"success"`
	Data    map[string]TradeHistory `json:"return,omitempty"`
	Error   string                  `json:"error,omitempty"`
}
type TradeHistory struct {
	Pair      string  `json:"pair"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	Rate      float64 `json:"rate"`
	OrderID   float64 `json:"order_id"`
	MyOrder   int     `json:"is_your_order"`
	Timestamp float64 `json:"timestamp"`
}
type DepositAddress struct {
	Success int `json:"success"`
	Return  struct {
		Address         string  `json:"address"`
		ProcessedAmount float64 `json:"processed_amount"`
		ServerTime      int64   `json:"server_time"`
	} `json:"return"`
	Error string `json:"error"`
}
type WithdrawCoinsToAddress struct {
	ServerTime int64  `json:"server_time"`
	Error      string `json:"error"`
}
type CreateCoupon struct {
	Coupon  string             `json:"coupon"`
	TransID int64              `json:"transID"`
	Funds   map[string]float64 `json:"funds"`
	Error   string             `json:"error"`
}
type RedeemCoupon struct {
	CouponAmount   float64            `json:"couponAmount,string"`
	CouponCurrency string             `json:"couponCurrency"`
	TransID        int64              `json:"transID"`
	Funds          map[string]float64 `json:"funds"`
	Error          string             `json:"error"`
}
