package alphapoint

type Response struct {
	IsAccepted    bool    `json:"isAccepted"`
	RejectReason  string  `json:"rejectReason"`
	Fee           float64 `json:"fee"`
	FeeProduct    string  `json:"feeProduct"`
	CancelOrderID int64   `json:"cancelOrderId"`
	ServerOrderID int64   `json:"serverOrderId"`
	DateTimeUTC   float64 `json:"dateTimeUtc"`
	ModifyOrderID int64   `json:"modifyOrderId"`
	Addresses     []DepositAddresses
}
type Ticker struct {
	High               float64 `json:"high"`
	Last               float64 `json:"last"`
	Bid                float64 `json:"bid"`
	Volume             float64 `json:"volume"`
	Low                float64 `json:"low"`
	Ask                float64 `json:"ask"`
	Total24HrQtyTraded float64 `json:"Total24HrQtyTraded"`
	Total24HrNumTrades float64 `json:"Total24HrNumTrades"`
	SellOrderCount     float64 `json:"sellOrderCount"`
	BuyOrderCount      float64 `json:"buyOrderCount"`
	NumOfCreateOrders  float64 `json:"numOfCreateOrders"`
	IsAccepted         bool    `json:"isAccepted"`
	RejectReason       string  `json:"rejectReason"`
}
type Trades struct {
	IsAccepted   bool    `json:"isAccepted"`
	RejectReason string  `json:"rejectReason"`
	DateTimeUTC  int64   `json:"dateTimeUtc"`
	Instrument   string  `json:"ins"`
	StartIndex   int     `json:"startIndex"`
	Count        int     `json:"count"`
	StartDate    int64   `json:"startDate"`
	EndDate      int64   `json:"endDate"`
	Trades       []Trade `json:"trades"`
}
type Trade struct {
	TID                   int64   `json:"tid"`
	Price                 float64 `json:"px"`
	Quantity              float64 `json:"qty"`
	Unixtime              int     `json:"unixtime"`
	UTCTicks              int64   `json:"utcticks"`
	IncomingOrderSide     int     `json:"incomingOrderSide"`
	IncomingServerOrderID int     `json:"incomingServerOrderId"`
	BookServerOrderID     int     `json:"bookServerOrderId"`
}
type Orderbook struct {
	Bids         []OrderbookEntry `json:"bids"`
	Asks         []OrderbookEntry `json:"asks"`
	IsAccepted   bool             `json:"isAccepted"`
	RejectReason string           `json:"rejectReason"`
}
type OrderbookEntry struct {
	Quantity float64 `json:"qty"`
	Price    float64 `json:"px"`
}
type ProductPairs struct {
	ProductPairs []ProductPair `json:"productPairs"`
	IsAccepted   bool          `json:"isAccepted"`
	RejectReason string        `json:"rejectReason"`
}
type ProductPair struct {
	Name                  string `json:"name"`
	Productpaircode       int    `json:"productPairCode"`
	Product1Label         string `json:"product1Label"`
	Product1Decimalplaces int    `json:"product1DecimalPlaces"`
	Product2Label         string `json:"product2Label"`
	Product2Decimalplaces int    `json:"product2DecimalPlaces"`
}
type Products struct {
	Products     []Product `json:"products"`
	IsAccepted   bool      `json:"isAccepted"`
	RejectReason string    `json:"rejectReason"`
}
type Product struct {
	Name          string `json:"name"`
	IsDigital     bool   `json:"isDigital"`
	ProductCode   int    `json:"productCode"`
	DecimalPlaces int    `json:"decimalPlaces"`
	FullName      string `json:"fullName"`
}
type UserInfo struct {
	UserInforKVPs []UserInfoKVP `json:"userInfoKVP"`
	IsAccepted    bool          `json:"isAccepted"`
	RejectReason  string        `json:"rejectReason"`
}
type UserInfoKVP struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type UserInfoSet struct {
	IsAccepted        string `json:"isAccepted"`
	RejectReason      string `json:"rejectReason"`
	RequireAuthy2FA   bool   `json:"requireAuthy2FA"`
	Val2FaRequestCode string `json:"val2FaRequestCode"`
}
type AccountInfo struct {
	Currencies []struct {
		Name    string `json:"name"`
		Balance int    `json:"balance"`
		Hold    int    `json:"hold"`
	} `json:"currencies"`
	ProductPairs []struct {
		ProductPairName string `json:"productPairName"`
		ProductPairCode int    `json:"productPairCode"`
		TradeCount      int    `json:"tradeCount"`
		TradeVolume     int    `json:"tradeVolume"`
	} `json:"productPairs"`
	IsAccepted   bool   `json:"isAccepted"`
	RejectReason string `json:"rejectReason"`
}
type Order struct {
	ServerOrderID int     `json:"ServerOrderId"`
	AccountID     int     `json:"AccountId"`
	Price         float64 `json:"Price"`
	QtyTotal      float64 `json:"QtyTotal"`
	QtyRemaining  float64 `json:"QtyRemaining"`
	ReceiveTime   int64   `json:"ReceiveTime"`
	Side          int64   `json:"Side"`
	State         int     `json:"orderState"`
	OrderType     int     `json:"orderType"`
}
type OpenOrders struct {
	Instrument string  `json:"ins"`
	OpenOrders []Order `json:"openOrders"`
}
type OrderInfo struct {
	OpenOrders   []OpenOrders `json:"openOrdersInfo"`
	IsAccepted   bool         `json:"isAccepted"`
	DateTimeUTC  int64        `json:"dateTimeUtc"`
	RejectReason string       `json:"rejectReason"`
}
type DepositAddresses struct {
	Name           string `json:"name"`
	DepositAddress string `json:"depositAddress"`
}
type WebsocketTicker struct {
	MessageType             string  `json:"messageType"`
	ProductPair             string  `json:"prodPair"`
	High                    float64 `json:"high"`
	Low                     float64 `json:"low"`
	Last                    float64 `json:"last"`
	Volume                  float64 `json:"volume"`
	Volume24Hrs             float64 `json:"volume24hrs"`
	Volume24HrsProduct2     float64 `json:"volume24hrsProduct2"`
	Total24HrQtyTraded      float64 `json:"Total24HrQtyTraded"`
	Total24HrProduct2Traded float64 `json:"Total24HrProduct2Traded"`
	Total24HrNumTrades      float64 `json:"Total24HrNumTrades"`
	Bid                     float64 `json:"bid"`
	Ask                     float64 `json:"ask"`
	BuyOrderCount           int     `json:"buyOrderCount"`
	SellOrderCount          int     `json:"sellOrderCount"`
}
