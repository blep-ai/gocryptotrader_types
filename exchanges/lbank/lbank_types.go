package lbank

import (
	"encoding/json"

	"github.com/blep-ai/gocryptotrader_types/currency"
)
type Ticker struct {
	Change   float64 `json:"change"`
	High     float64 `json:"high"`
	Latest   float64 `json:"latest"`
	Low      float64 `json:"low"`
	Turnover float64 `json:"turnover"`
	Volume   float64 `json:"vol"`
}
type TickerResponse struct {
	Symbol    currency.Pair `json:"symbol"`
	Timestamp int64         `json:"timestamp"`
	Ticker    Ticker        `json:"ticker"`
}
type MarketDepthResponse struct {
	ErrCapture `json:",omitempty"`
	Asks       [][]float64 `json:"asks"`
	Bids       [][]float64 `json:"bids"`
	Timestamp  int64       `json:"timestamp"`
}
type TradeResponse struct {
	DateMS int64   `json:"date_ms"`
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	Type   string  `json:"type"`
	TID    string  `json:"tid"`
}
type KlineResponse struct {
	TimeStamp     int64   `json:"timestamp"`
	OpenPrice     float64 `json:"openprice"`
	HigestPrice   float64 `json:"highestprice"`
	LowestPrice   float64 `json:"lowestprice"`
	ClosePrice    float64 `json:"closeprice"`
	TradingVolume float64 `json:"tradingvolume"`
}
type InfoResponse struct {
	Freeze map[string]string `json:"freeze"`
	Asset  map[string]string `json:"asset"`
	Free   map[string]string `json:"Free"`
}
type InfoFinalResponse struct {
	ErrCapture `json:",omitempty"`
	Info       InfoResponse `json:"info"`
}
type CreateOrderResponse struct {
	ErrCapture `json:",omitempty"`
	OrderID    string `json:"order_id"`
}
type RemoveOrderResponse struct {
	ErrCapture `json:",omitempty"`
	Err        string `json:"error"`
	OrderID    string `json:"order_id"`
	Success    string `json:"success"`
}
type OrderResponse struct {
	Symbol     string  `json:"symbol"`
	Amount     float64 `json:"amount"`
	CreateTime int64   `json:"created_time"`
	Price      float64 `json:"price"`
	AvgPrice   float64 `json:"avg_price"`
	Type       string  `json:"type"`
	OrderID    string  `json:"order_id"`
	DealAmount float64 `json:"deal_amount"`
	Status     int64   `json:"status"`
}
type QueryOrderResponse struct {
	ErrCapture `json:",omitempty"`
	Orders     json.RawMessage `json:"orders"`
}
type QueryOrderFinalResponse struct {
	ErrCapture
	Orders []OrderResponse
}
type OrderHistory struct {
	Result      bool            `json:"result,string"`
	Total       string          `json:"total"`
	PageLength  uint8           `json:"page_length"`
	Orders      json.RawMessage `json:"orders"`
	CurrentPage uint8           `json:"current_page"`
	ErrorCode   int64           `json:"error_code"`
}
type OrderHistoryResponse struct {
	ErrCapture  `json:",omitempty"`
	PageLength  uint8           `json:"page_length"`
	Orders      json.RawMessage `json:"orders"`
	CurrentPage uint8           `json:"current_page"`
}
type OrderHistoryFinalResponse struct {
	ErrCapture
	PageLength  uint8
	Orders      []OrderResponse
	CurrentPage uint8
}
type PairInfoResponse struct {
	MinimumQuantity  string `json:"minTranQua"`
	PriceAccuracy    string `json:"priceAccuracy"`
	QuantityAccuracy string `json:"quantityAccuracy"`
	Symbol           string `json:"symbol"`
}
type TransactionTemp struct {
	TxUUID       string  `json:"txUuid"`
	OrderUUID    string  `json:"orderUuid"`
	TradeType    string  `json:"tradeType"`
	DealTime     int64   `json:"dealTime"`
	DealPrice    float64 `json:"dealPrice"`
	DealQuantity float64 `json:"dealQuantity"`
	DealVolPrice float64 `json:"dealVolumePrice"`
	TradeFee     float64 `json:"tradeFee"`
	TradeFeeRate float64 `json:"tradeFeeRate"`
}
type TransactionHistoryResp struct {
	ErrCapture  `json:",omitempty"`
	Transaction []TransactionTemp `json:"transaction"`
}
type OpenOrderResponse struct {
	ErrCapture `json:",omitempty"`
	PageLength uint8           `json:"page_length"`
	PageNumber uint8           `json:"page_number"`
	Total      string          `json:"total"`
	Orders     json.RawMessage `json:"orders"`
}
type OpenOrderFinalResponse struct {
	ErrCapture
	PageLength uint8
	PageNumber uint8
	Total      string
	Orders     []OrderResponse
}
type ExchangeRateResponse struct {
	USD2CNY string `json:"USD2CNY"`
}
type WithdrawConfigResponse struct {
	AssetCode   string `json:"assetCode"`
	Minimum     string `json:"min"`
	CanWithDraw bool   `json:"canWithDraw"`
	Fee         string `json:"fee"`
}
type WithdrawResponse struct {
	ErrCapture `json:",omitempty"`
	WithdrawID string  `json:"withdrawId"`
	Fee        float64 `json:"fee"`
}
type RevokeWithdrawResponse struct {
	ErrCapture `json:",omitempty"`
	WithdrawID string `json:"string"`
}
type ListDataResponse struct {
	ErrCapture `json:",omitempty"`
	Amount     float64 `json:"amount"`
	AssetCode  string  `json:"assetCode"`
	Address    string  `json:"address"`
	Fee        float64 `json:"fee"`
	ID         int64   `json:"id"`
	Time       int64   `json:"time"`
	TXHash     string  `json:"txhash"`
	Status     string  `json:"status"`
}
type WithdrawalResponse struct {
	ErrCapture `json:",omitempty"`
	TotalPages int64              `json:"totalPages"`
	PageSize   int64              `json:"pageSize"`
	PageNo     int64              `json:"pageNo"`
	List       []ListDataResponse `json:"list"`
}
type ErrCapture struct {
	Error  int64 `json:"error_code"`
	Result bool  `json:"result,string"`
}
type GetAllOpenIDResp struct {
	CurrencyPair string
	OrderID      string
}
