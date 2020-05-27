package bithumb

import (
	"encoding/json"
)

type Ticker struct {
	OpeningPrice              float64 `json:"opening_price,string"`
	ClosingPrice              float64 `json:"closing_price,string"`
	MinPrice                  float64 `json:"min_price,string"`
	MaxPrice                  float64 `json:"max_price,string"`
	UnitsTraded               float64 `json:"units_traded,string"`
	AccumulatedTradeValue     float64 `json:"acc_trade_value,string"`
	PreviousClosingPrice      float64 `json:"prev_closing_price,string"`
	UnitsTraded24Hr           float64 `json:"units_traded_24H,string"`
	AccumulatedTradeValue24hr float64 `json:"acc_trade_value_24H,string"`
	Fluctate24Hr              string  `json:"fluctate_24H"`
	FluctateRate24hr          float64 `json:"fluctate_rate_24H,string"`
	Date                      int64   `json:"date,string"`
}
type TickerResponse struct {
	Status  string `json:"status"`
	Data    Ticker `json:"data"`
	Message string `json:"message"`
}
type TickersResponse struct {
	Status  string                     `json:"status"`
	Data    map[string]json.RawMessage `json:"data"`
	Message string                     `json:"message"`
}
type Orderbook struct {
	Status string `json:"status"`
	Data   struct {
		Timestamp       int64  `json:"timestamp,string"`
		OrderCurrency   string `json:"order_currency"`
		PaymentCurrency string `json:"payment_currency"`
		Bids            []struct {
			Quantity float64 `json:"quantity,string"`
			Price    float64 `json:"price,string"`
		} `json:"bids"`
		Asks []struct {
			Quantity float64 `json:"quantity,string"`
			Price    float64 `json:"price,string"`
		} `json:"asks"`
	} `json:"data"`
	Message string `json:"message"`
}
type TransactionHistory struct {
	Status string `json:"status"`
	Data   []struct {
		ContNumber      int64   `json:"cont_no,string"`
		TransactionDate string  `json:"transaction_date"`
		Type            string  `json:"type"`
		UnitsTraded     float64 `json:"units_traded,string"`
		Price           float64 `json:"price,string"`
		Total           float64 `json:"total,string"`
	} `json:"data"`
	Message string `json:"message"`
}
type Account struct {
	Status string `json:"status"`
	Data   struct {
		Created   int64   `json:"created,string"`
		AccountID string  `json:"account_id"`
		TradeFee  float64 `json:"trade_fee,string"`
		Balance   float64 `json:"balance,string"`
	} `json:"data"`
	Message string `json:"message"`
}
type Balance struct {
	Status  string                 `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}
type WalletAddressRes struct {
	Status string `json:"status"`
	Data   struct {
		WalletAddress string `json:"wallet_address"`
		Currency      string `json:"currency"`
	} `json:"data"`
	Message string `json:"message"`
}
type LastTransactionTicker struct {
	Status string `json:"status"`
	Data   struct {
		OpeningPrice float64 `json:"opening_price,string"`
		ClosingPrice float64 `json:"closing_price,string"`
		MinPrice     float64 `json:"min_price,string"`
		MaxPrice     float64 `json:"max_price,string"`
		AveragePrice float64 `json:"average_price,string"`
		UnitsTraded  float64 `json:"units_traded,string"`
		Volume1Day   float64 `json:"volume_1day,string"`
		Volume7Day   float64 `json:"volume_7day,string"`
		BuyPrice     int64   `json:"buy_price,string"`
		SellPrice    int64   `json:"sell_price,string"`
		Date         int64   `json:"date,string"`
	} `json:"data"`
	Message string `json:"message"`
}
type Orders struct {
	Status  string      `json:"status"`
	Data    []OrderData `json:"data"`
	Message string      `json:"message"`
}
type OrderData struct {
	OrderID         string  `json:"order_id"`
	OrderCurrency   string  `json:"order_currency"`
	OrderDate       int64   `json:"order_date"`
	PaymentCurrency string  `json:"payment_currency"`
	Type            string  `json:"type"`
	Status          string  `json:"status"`
	Units           float64 `json:"units,string"`
	UnitsRemaining  float64 `json:"units_remaining,string"`
	Price           float64 `json:"price,string"`
	Fee             float64 `json:"fee,string"`
	Total           float64 `json:"total,string"`
	DateCompleted   int64   `json:"date_completed"`
}
type UserTransactions struct {
	Status string `json:"status"`
	Data   []struct {
		Search       string  `json:"search"`
		TransferDate int64   `json:"transfer_date"`
		Units        string  `json:"units"`
		Price        float64 `json:"price,string"`
		BTC1KRW      float64 `json:"btc1krw,string"`
		Fee          string  `json:"fee"`
		BTCRemain    float64 `json:"btc_remain,string"`
		KRWRemain    float64 `json:"krw_remain,string"`
	} `json:"data"`
	Message string `json:"message"`
}
type OrderPlace struct {
	Status string `json:"status"`
	Data   []struct {
		ContID string  `json:"cont_id"`
		Units  float64 `json:"units,string"`
		Price  float64 `json:"price,string"`
		Total  float64 `json:"total,string"`
		Fee    float64 `json:"fee,string"`
	} `json:"data"`
	Message string `json:"message"`
}
type OrderDetails struct {
	Status string `json:"status"`
	Data   []struct {
		TransactionDate int64   `json:"transaction_date,string"`
		Type            string  `json:"type"`
		OrderCurrency   string  `json:"order_currency"`
		PaymentCurrency string  `json:"payment_currency"`
		UnitsTraded     float64 `json:"units_traded,string"`
		Price           float64 `json:"price,string"`
		Total           float64 `json:"total,string"`
	} `json:"data"`
	Message string `json:"message"`
}
type ActionStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type KRWDeposit struct {
	Status   string `json:"status"`
	Account  string `json:"account"`
	Bank     string `json:"bank"`
	BankUser string `json:"BankUser"`
	Message  string `json:"message"`
}
type MarketBuy struct {
	Status  string `json:"status"`
	OrderID string `json:"order_id"`
	Data    []struct {
		ContID string  `json:"cont_id"`
		Units  float64 `json:"units,string"`
		Price  float64 `json:"price,string"`
		Total  float64 `json:"total,string"`
		Fee    float64 `json:"fee,string"`
	} `json:"data"`
	Message string `json:"message"`
}
type MarketSell struct {
	Status  string `json:"status"`
	OrderID string `json:"order_id"`
	Data    []struct {
		ContID string  `json:"cont_id"`
		Units  float64 `json:"units,string"`
		Price  float64 `json:"price,string"`
		Total  float64 `json:"total,string"`
		Fee    float64 `json:"fee,string"`
	} `json:"data"`
	Message string `json:"message"`
}
type FullBalance struct {
	InUse     map[string]float64
	Misu      map[string]float64
	Total     map[string]float64
	Xcoin     map[string]float64
	Available map[string]float64
}
