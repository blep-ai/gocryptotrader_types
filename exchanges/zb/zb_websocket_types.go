package zb

import (
	"encoding/json"

	"github.com/blep-ai/gocryptotrader_types/currency"
)

type Subscription struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
	No      int64  `json:"no,string,omitempty"`
}
type Generic struct {
	Code    int64           `json:"code"`
	Channel string          `json:"channel"`
	Message interface{}     `json:"message"`
	No      int64           `json:"no,string,omitempty"`
	Data    json.RawMessage `json:"data"`
}
type WsTicker struct {
	Date int64 `json:"date,string"`
	Data struct {
		Volume24Hr float64 `json:"vol,string"`
		High       float64 `json:"high,string"`
		Low        float64 `json:"low,string"`
		Last       float64 `json:"last,string"`
		Buy        float64 `json:"buy,string"`
		Sell       float64 `json:"sell,string"`
	} `json:"ticker"`
}
type WsDepth struct {
	Timestamp int64           `json:"timestamp"`
	Asks      [][]interface{} `json:"asks"`
	Bids      [][]interface{} `json:"bids"`
}
type WsTrades struct {
	Data []struct {
		Amount    float64     `json:"amount,string"`
		Price     float64     `json:"price,string"`
		TID       interface{} `json:"tid"`
		Date      int64       `json:"date"`
		Type      string      `json:"type"`
		TradeType string      `json:"trade_type"`
	} `json:"data"`
}
type WsAuthenticatedRequest struct {
	Accesskey string `json:"accesskey"`
	Channel   string `json:"channel"`
	Event     string `json:"event"`
	No        int64  `json:"no,string,omitempty"`
	Sign      string `json:"sign,omitempty"`
}
type WsAddSubUserRequest struct {
	Accesskey   string `json:"accesskey"`
	Channel     string `json:"channel"`
	Event       string `json:"event"`
	Memo        string `json:"memo"`
	Password    string `json:"password"`
	SubUserName string `json:"subUserName"`
	No          int64  `json:"no,string,omitempty"`
	Sign        string `json:"sign,omitempty"`
}
type WsCreateSubUserKeyRequest struct {
	Accesskey   string `json:"accesskey"`
	AssetPerm   bool   `json:"assetPerm,string"`
	Channel     string `json:"channel"`
	EntrustPerm bool   `json:"entrustPerm,string"`
	Event       string `json:"event"`
	KeyName     string `json:"keyName"`
	LeverPerm   bool   `json:"leverPerm,string"`
	MoneyPerm   bool   `json:"moneyPerm,string"`
	No          int64  `json:"no,string,omitempty"`
	Sign        string `json:"sign,omitempty"`
	ToUserID    string `json:"toUserId"`
}
type WsDoTransferFundsRequest struct {
	Accesskey    string        `json:"accesskey"`
	Amount       float64       `json:"amount,string"`
	Channel      string        `json:"channel"`
	Currency     currency.Code `json:"currency"`
	Event        string        `json:"event"`
	FromUserName string        `json:"fromUserName"`
	No           int64         `json:"no,string"`
	Sign         string        `json:"sign,omitempty"`
	ToUserName   string        `json:"toUserName"`
}
type WsGetSubUserListResponse struct {
	Success bool                           `json:"success"`
	Code    int64                          `json:"code"`
	Channel string                         `json:"channel"`
	Message []WsGetSubUserListResponseData `json:"message"`
	No      int64                          `json:"no,string"`
}
type WsGetSubUserListResponseData struct {
	IsOpenAPI bool   `json:"isOpenApi,omitempty"`
	Memo      string `json:"memo,omitempty"`
	UserName  string `json:"userName,omitempty"`
	UserID    int64  `json:"userId,omitempty"`
	IsFreez   bool   `json:"isFreez,omitempty"`
}
type WsRequestResponse struct {
	Success bool        `json:"success"`
	Code    int64       `json:"code"`
	Channel string      `json:"channel"`
	Message interface{} `json:"message"`
	No      int64       `json:"no,string"`
}
type WsSubmitOrderRequest struct {
	Accesskey string  `json:"accesskey"`
	Amount    float64 `json:"amount,string"`
	Channel   string  `json:"channel"`
	Event     string  `json:"event"`
	No        int64   `json:"no,string,omitempty"`
	Price     float64 `json:"price,string"`
	Sign      string  `json:"sign,omitempty"`
	TradeType int64   `json:"tradeType,string"`
}
type WsSubmitOrderResponse struct {
	Message string `json:"message"`
	No      int64  `json:"no,string"`
	Data    struct {
		EntrustID int64 `json:"intrustID"`
	} `json:"data"`
	Code    int64  `json:"code"`
	Channel string `json:"channel"`
	Success bool   `json:"success"`
}
type WsCancelOrderRequest struct {
	Accesskey string `json:"accesskey"`
	Channel   string `json:"channel"`
	Event     string `json:"event"`
	ID        int64  `json:"id"`
	Sign      string `json:"sign,omitempty"`
	No        int64  `json:"no,string"`
}
type WsCancelOrderResponse struct {
	Message string `json:"message"`
	No      int64  `json:"no,string"`
	Code    int64  `json:"code"`
	Channel string `json:"channel"`
	Success bool   `json:"success"`
}
type WsGetOrderRequest struct {
	Accesskey string `json:"accesskey"`
	Channel   string `json:"channel"`
	Event     string `json:"event"`
	ID        int64  `json:"id"`
	Sign      string `json:"sign,omitempty"`
	No        int64  `json:"no,string"`
}
type WsGetOrderResponse struct {
	Message string  `json:"message"`
	No      int64   `json:"no,string"`
	Code    int64   `json:"code"`
	Channel string  `json:"channel"`
	Success bool    `json:"success"`
	Data    []Order `json:"data"`
}
type WsGetOrdersRequest struct {
	Accesskey string `json:"accesskey"`
	Channel   string `json:"channel"`
	Event     string `json:"event"`
	No        int64  `json:"no,string"`
	PageIndex int64  `json:"pageIndex"`
	TradeType int64  `json:"tradeType"`
	Sign      string `json:"sign,omitempty"`
}
type WsGetOrdersResponse struct {
	Message string  `json:"message"`
	No      int64   `json:"no,string"`
	Code    int64   `json:"code"`
	Channel string  `json:"channel"`
	Success bool    `json:"success"`
	Data    []Order `json:"data"`
}
type WsGetOrdersIgnoreTradeTypeRequest struct {
	Accesskey string `json:"accesskey"`
	Channel   string `json:"channel"`
	Event     string `json:"event"`
	No        int64  `json:"no,string"`
	PageIndex int64  `json:"pageIndex"`
	PageSize  int64  `json:"pageSize"`
	Sign      string `json:"sign,omitempty"`
}
type WsGetOrdersIgnoreTradeTypeResponse struct {
	Message string  `json:"message"`
	No      int64   `json:"no,string"`
	Code    int64   `json:"code"`
	Channel string  `json:"channel"`
	Success bool    `json:"success"`
	Data    []Order `json:"data"`
}
type WsGetAccountInfoResponse struct {
	Message string `json:"message"`
	No      int64  `json:"no,string"`
	Data    struct {
		Coins []AccountsResponseCoin `json:"coins"`
		Base  AccountsBaseResponse   `json:"base"`
	} `json:"data"`
	Code    int64  `json:"code"`
	Channel string `json:"channel"`
	Success bool   `json:"success"`
}
