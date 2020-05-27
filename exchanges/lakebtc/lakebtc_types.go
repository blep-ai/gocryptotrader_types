import "github.com/toorop/go-pusher"

type Ticker struct {
	Last   float64
	Bid    float64
	Ask    float64
	High   float64
	Low    float64
	Volume float64
}
type OrderbookStructure struct {
	Price  float64
	Amount float64
}
type Orderbook struct {
	Bids []OrderbookStructure `json:"bids"`
	Asks []OrderbookStructure `json:"asks"`
}
type TickerResponse struct {
	Last   interface{}
	Bid    interface{}
	Ask    interface{}
	High   interface{}
	Low    interface{}
	Volume interface{}
}
type TradeHistory struct {
	Date   int64   `json:"data"`
	Price  float64 `json:"price,string"`
	Amount float64 `json:"amount,string"`
	TID    int64   `json:"tid"`
}
type AccountInfo struct {
	Balance map[string]string `json:"balance"`
	Locked  map[string]string `json:"locked"`
	Profile struct {
		Email             string `json:"email"`
		UID               string `json:"uid"`
		BTCDepositAddress string `json:"btc_deposit_addres"` // nolint // API misspelling
	} `json:"profile"`
}
type Trade struct {
	ID     int64  `json:"id"`
	Result string `json:"result"`
}
type OpenOrders struct {
	ID     int64   `json:"id"`
	Amount float64 `json:"amount,string"`
	Price  float64 `json:"price,string"`
	Symbol string  `json:"symbol"`
	Type   string  `json:"type"`
	At     int64   `json:"at"`
}
type Orders struct {
	ID             int64   `json:"id"`
	OriginalAmount float64 `json:"original_amount,string"`
	Amount         float64 `json:"amount,string"`
	Price          float64 `json:"price,string"`
	Symbol         string  `json:"symbol"`
	Type           string  `json:"type"`
	State          string  `json:"state"`
	At             int64   `json:"at"`
}
type AuthenticatedTradeHistory struct {
	Type   string  `json:"type"`
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount,string"`
	Total  float64 `json:"total,string"`
	At     int64   `json:"at"`
}
type ExternalAccounts struct {
	ID         int64       `json:"id,string"`
	Type       string      `json:"type"`
	Address    string      `json:"address"`
	Alias      interface{} `json:"alias"`
	Currencies string      `json:"currencies"`
	State      string      `json:"state"`
	UpdatedAt  int64       `json:"updated_at,string"`
}
type Withdraw struct {
	ID                int64   `json:"id,string"`
	Amount            float64 `json:"amount,string"`
	Currency          string  `json:"currency"`
	Fee               float64 `json:"fee,string"`
	State             string  `json:"state"`
	Source            string  `json:"source"`
	ExternalAccountID int64   `json:"external_account_id,string"`
	At                int64   `json:"at"`
	Error             string  `json:"error"`
}
type WebsocketConn struct {
	Client    *pusher.Client
	Ticker    chan *pusher.Event
	Orderbook chan *pusher.Event
	Trade     chan *pusher.Event
}
type WsOrderbookUpdate struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}
type WsTrades struct {
	Trades []WsTrade `json:"trades"`
}
type WsTrade struct {
	Type   string  `json:"type"`
	Date   int64   `json:"date"`
	Price  float64 `json:"price,string"`
	Amount float64 `json:"amount,string"`
}
