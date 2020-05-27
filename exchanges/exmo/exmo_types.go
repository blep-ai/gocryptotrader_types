package exmo

type Trades struct {
	TradeID  int64   `json:"trade_id"`
	Type     string  `json:"string"`
	Quantity float64 `json:"quantity,string"`
	Price    float64 `json:"price,string"`
	Amount   float64 `json:"amount,string"`
	Date     int64   `json:"date"`
	Pair     string  `json:"pair"`
}
type Orderbook struct {
	AskQuantity float64    `json:"ask_quantity,string"`
	AskAmount   float64    `json:"ask_amount,string"`
	AskTop      float64    `json:"ask_top,string"`
	BidQuantity float64    `json:"bid_quantity,string"`
	BidTop      float64    `json:"bid_top,string"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}
type Ticker struct {
	Buy           float64 `json:"buy_price,string"`
	Sell          float64 `json:"sell_price,string"`
	Last          float64 `json:"last_trade,string"`
	High          float64 `json:"high,string"`
	Low           float64 `json:"low,string"`
	Average       float64 `json:"average,string"`
	Volume        float64 `json:"vol,string"`
	VolumeCurrent float64 `json:"vol_curr,string"`
	Updated       int64   `json:"updated"`
}
type PairSettings struct {
	MinQuantity float64 `json:"min_quantity,string"`
	MaxQuantity float64 `json:"max_quantity,string"`
	MinPrice    float64 `json:"min_price,string"`
	MaxPrice    float64 `json:"max_price,string"`
	MaxAmount   float64 `json:"max_amount,string"`
	MinAmount   float64 `json:"min_amount,string"`
}
type AuthResponse struct {
	Result bool   `json:"bool"`
	Error  string `json:"error"`
}
type UserInfo struct {
	AuthResponse
	UID        int               `json:"uid"`
	ServerDate int               `json:"server_date"`
	Balances   map[string]string `json:"balances"`
	Reserved   map[string]string `json:"reserved"`
}
type OpenOrders struct {
	OrderID  int64   `json:"order_id,string"`
	Created  int64   `json:"created,string"`
	Type     string  `json:"type"`
	Pair     string  `json:"pair"`
	Price    float64 `json:"price,string"`
	Quantity float64 `json:"quantity,string"`
	Amount   float64 `json:"amount,string"`
}
type UserTrades struct {
	TradeID  int64   `json:"trade_id"`
	Date     int64   `json:"date"`
	Type     string  `json:"type"`
	Pair     string  `json:"pair"`
	OrderID  int64   `json:"order_id"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
	Amount   float64 `json:"amount"`
}
type CancelledOrder struct {
	Date     int64   `json:"date"`
	OrderID  int64   `json:"order_id,string"`
	Type     string  `json:"type"`
	Pair     string  `json:"pair"`
	Price    float64 `json:"price,string"`
	Quantity float64 `json:"quantity,string"`
	Amount   float64 `json:"amount,string"`
}
type OrderTrades struct {
	Type        string       `json:"type"`
	InCurrency  string       `json:"in_currency"`
	InAmount    float64      `json:"in_amount,string"`
	OutCurrency string       `json:"out_currency"`
	OutAmount   float64      `json:"out_amount,string"`
	Trades      []UserTrades `json:"trades"`
}
type RequiredAmount struct {
	Quantity float64 `json:"quantity,string"`
	Amount   float64 `json:"amount,string"`
	AvgPrice float64 `json:"avg_price,string"`
}
type ExcodeCreate struct {
	TaskID   int64             `json:"task_id"`
	Code     string            `json:"code"`
	Amount   float64           `json:"amount,string"`
	Currency string            `json:"currency"`
	Balances map[string]string `json:"balances"`
}
type ExcodeLoad struct {
	TaskID   int64             `json:"task_id"`
	Amount   float64           `json:"amount,string"`
	Currency string            `json:"currency"`
	Balances map[string]string `json:"balances"`
}
type WalletHistory struct {
	Begin   int64 `json:"begin,string"`
	End     int64 `json:"end,string"`
	History []struct {
		Timestamp int64   `json:"dt"`
		Type      string  `json:"string"`
		Currency  string  `json:"curr"`
		Status    string  `json:"status"`
		Provider  string  `json:"provider"`
		Amount    float64 `json:"amount,string"`
		Account   string  `json:"account,string"`
	}
}
