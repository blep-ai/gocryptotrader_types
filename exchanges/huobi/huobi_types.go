package huobi

type Response struct {
	Status       string `json:"status"`
	Channel      string `json:"ch"`
	Timestamp    int64  `json:"ts"`
	ErrorCode    string `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
}
type KlineItem struct {
	ID     int64   `json:"id"`
	Open   float64 `json:"open"`
	Close  float64 `json:"close"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Amount float64 `json:"amount"`
	Volume float64 `json:"vol"`
	Count  int     `json:"count"`
}
type CancelOpenOrdersBatch struct {
	Data struct {
		FailedCount  int `json:"failed-count"`
		NextID       int `json:"next-id"`
		SuccessCount int `json:"success-count"`
	} `json:"data"`
	Status       string `json:"status"`
	ErrorMessage string `json:"err-msg"`
}
type DetailMerged struct {
	Detail
	Version int64     `json:"version"`
	Ask     []float64 `json:"ask"`
	Bid     []float64 `json:"bid"`
}
type Tickers struct {
	Data []Ticker `json:"data"`
}
type Ticker struct {
	Amount float64 `json:"amount"`
	Close  float64 `json:"close"`
	Count  int64   `json:"count"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Open   float64 `json:"open"`
	Symbol string  `json:"symbol"`
	Volume float64 `json:"vol"`
}
type OrderBookDataRequestParams struct {
	Symbol string                         `json:"symbol"` // Required; example LTCBTC,BTCUSDT
	Type   OrderBookDataRequestParamsType `json:"type"`   // step0, step1, step2, step3, step4, step5 (combined depth 0-5); when step0, no depth is merged
}
type Orderbook struct {
	ID         int64       `json:"id"`
	Timetstamp int64       `json:"ts"`
	Bids       [][]float64 `json:"bids"`
	Asks       [][]float64 `json:"asks"`
}
type Trade struct {
	ID        float64 `json:"id"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Direction string  `json:"direction"`
	Timestamp int64   `json:"ts"`
}
type TradeHistory struct {
	ID        int64   `json:"id"`
	Timestamp int64   `json:"ts"`
	Trades    []Trade `json:"data"`
}
type Detail struct {
	Amount    float64 `json:"amount"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Timestamp int64   `json:"timestamp"`
	ID        int64   `json:"id"`
	Count     int     `json:"count"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"vol"`
}
type Symbol struct {
	BaseCurrency       string  `json:"base-currency"`
	QuoteCurrency      string  `json:"quote-currency"`
	PricePrecision     int     `json:"price-precision"`
	AmountPrecision    int     `json:"amount-precision"`
	SymbolPartition    string  `json:"symbol-partition"`
	Innovation         string  `json:"innovation"`
	State              string  `json:"state"`
	ValuePrecision     int     `json:"value-precision"`
	MinimumOrderAmount float64 `json:"min-order-amt"`
	MaximumOrderAmount float64 `json:"max-order-amt"`
	MinimumOrderValue  float64 `json:"min-order-value"`
}
type Account struct {
	ID     int64  `json:"id"`
	Type   string `json:"type"`
	State  string `json:"state"`
	UserID int64  `json:"user-id"`
}
type AccountBalance struct {
	ID                    int64                  `json:"id"`
	Type                  string                 `json:"type"`
	State                 string                 `json:"state"`
	AccountBalanceDetails []AccountBalanceDetail `json:"list"`
}
type AccountBalanceDetail struct {
	Currency string  `json:"currency"`
	Type     string  `json:"type"`
	Balance  float64 `json:"balance,string"`
}
type AggregatedBalance struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance,string"`
}
type CancelOrderBatch struct {
	Success []string `json:"success"`
	Failed  []struct {
		OrderID      int64  `json:"order-id,string"`
		ErrorCode    string `json:"err-code"`
		ErrorMessage string `json:"err-msg"`
	} `json:"failed"`
}
type OrderInfo struct {
	ID               int64   `json:"id"`
	Symbol           string  `json:"symbol"`
	AccountID        int64   `json:"account-id"`
	Amount           float64 `json:"amount,string"`
	Price            float64 `json:"price,string"`
	CreatedAt        int64   `json:"created-at"`
	Type             string  `json:"type"`
	FieldAmount      float64 `json:"field-amount,string"`
	FieldCashAmount  float64 `json:"field-cash-amount,string"`
	FilledAmount     float64 `json:"filled-amount,string"`
	FilledCashAmount float64 `json:"filled-cash-amount,string"`
	FilledFees       float64 `json:"filled-fees,string"`
	FinishedAt       int64   `json:"finished-at"`
	UserID           int64   `json:"user-id"`
	Source           string  `json:"source"`
	State            string  `json:"state"`
	CanceledAt       int64   `json:"canceled-at"`
	Exchange         string  `json:"exchange"`
	Batch            string  `json:"batch"`
}
type OrderMatchInfo struct {
	ID           int    `json:"id"`
	OrderID      int    `json:"order-id"`
	MatchID      int    `json:"match-id"`
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	Source       string `json:"source"`
	Price        string `json:"price"`
	FilledAmount string `json:"filled-amount"`
	FilledFees   string `json:"filled-fees"`
	CreatedAt    int64  `json:"created-at"`
}
type MarginOrder struct {
	Currency        string `json:"currency"`
	Symbol          string `json:"symbol"`
	AccruedAt       int64  `json:"accrued-at"`
	LoanAmount      string `json:"loan-amount"`
	LoanBalance     string `json:"loan-balance"`
	InterestBalance string `json:"interest-balance"`
	CreatedAt       int64  `json:"created-at"`
	InterestAmount  string `json:"interest-amount"`
	InterestRate    string `json:"interest-rate"`
	AccountID       int    `json:"account-id"`
	UserID          int    `json:"user-id"`
	UpdatedAt       int64  `json:"updated-at"`
	ID              int    `json:"id"`
	State           string `json:"state"`
}
type MarginAccountBalance struct {
	ID       int              `json:"id"`
	Type     string           `json:"type"`
	State    string           `json:"state"`
	Symbol   string           `json:"symbol"`
	FlPrice  string           `json:"fl-price"`
	FlType   string           `json:"fl-type"`
	RiskRate string           `json:"risk-rate"`
	List     []AccountBalance `json:"list"`
}
type SpotNewOrderRequestParams struct {
	AccountID int                           `json:"account-id,string"` // Account ID, obtained using the accounts method. Curency trades use the accountid of the ‘spot’ account; for loan asset transactions, please use the accountid of the ‘margin’ account.
	Amount    float64                       `json:"amount"`            // The limit price indicates the quantity of the order, the market price indicates how much to buy when the order is paid, and the market price indicates how much the coin is sold when the order is sold.
	Price     float64                       `json:"price"`             // Order price, market price does not use  this parameter
	Source    string                        `json:"source"`            // Order source, api: API call, margin-api: loan asset transaction
	Symbol    string                        `json:"symbol"`            // The symbol to use; example btcusdt, bccbtc......
	Type      SpotNewOrderRequestParamsType `json:"type"`              // 订单类型, buy-market: 市价买, sell-market: 市价卖, buy-limit: 限价买, sell-limit: 限价卖
}
type DepositAddress struct {
	Currency   string `json:"currency"`
	Address    string `json:"address"`
	AddressTag string `json:"addressTag"`
	Chain      string `json:"chain"`
}
type ChainQuota struct {
	Chain                         string  `json:"chain"`
	MaxWithdrawAmount             float64 `json:"maxWithdrawAmt,string"`
	WithdrawQuotaPerDay           float64 `json:"withdrawQuotaPerDay,string"`
	RemainingWithdrawQuotaPerDay  float64 `json:"remainWithdrawQuotaPerDay,string"`
	WithdrawQuotaPerYear          float64 `json:"withdrawQuotaPerYear,string"`
	RemainingWithdrawQuotaPerYear float64 `json:"remainWithdrawQuotaPerYear,string"`
	WithdrawQuotaTotal            float64 `json:"withdrawQuotaTotal,string"`
	RemainingWithdrawQuotaTotal   float64 `json:"remainWithdrawQuotaTotal,string"`
}
type WithdrawQuota struct {
	Currency string       `json:"currency"`
	Chains   []ChainQuota `json:"chains"`
}
type KlinesRequestParams struct {
	Symbol string       // Symbol to be used; example btcusdt, bccbtc......
	Period TimeInterval // Kline time interval; 1min, 5min, 15min......
	Size   int          // Size; [1-2000]
}
type WsRequest struct {
	Topic       string `json:"req,omitempty"`
	Subscribe   string `json:"sub,omitempty"`
	Unsubscribe string `json:"unsub,omitempty"`
	ClientID    int64  `json:"cid,string,omitempty"`
}
type WsResponse struct {
	Op           string `json:"op"`
	TS           int64  `json:"ts"`
	Status       string `json:"status"`
	ErrorCode    int64  `json:"err-code"`
	ErrorMessage string `json:"err-msg"`
	Ping         int64  `json:"ping"`
	Channel      string `json:"ch"`
	Rep          string `json:"rep"`
	Topic        string `json:"topic"`
	Subscribed   string `json:"subbed"`
	UnSubscribed string `json:"unsubbed"`
	ClientID     int64  `json:"cid,string"`
}
type WsHeartBeat struct {
	ClientNonce int64 `json:"ping"`
}
type WsDepth struct {
	Channel   string `json:"ch"`
	Timestamp int64  `json:"ts"`
	Tick      struct {
		Bids      [][]interface{} `json:"bids"`
		Asks      [][]interface{} `json:"asks"`
		Timestamp int64           `json:"ts"`
		Version   int64           `json:"version"`
	} `json:"tick"`
}
type WsKline struct {
	Channel   string `json:"ch"`
	Timestamp int64  `json:"ts"`
	Tick      struct {
		ID     int64   `json:"id"`
		Open   float64 `json:"open"`
		Close  float64 `json:"close"`
		Low    float64 `json:"low"`
		High   float64 `json:"high"`
		Amount float64 `json:"amount"`
		Volume float64 `json:"vol"`
		Count  int64   `json:"count"`
	}
}
type WsTick struct {
	Channel   string `json:"ch"`
	Rep       string `json:"rep"`
	Timestamp int64  `json:"ts"`
	Tick      struct {
		Amount    float64 `json:"amount"`
		Close     float64 `json:"close"`
		Count     float64 `json:"count"`
		High      float64 `json:"high"`
		ID        float64 `json:"id"`
		Low       float64 `json:"low"`
		Open      float64 `json:"open"`
		Timestamp float64 `json:"ts"`
		Volume    float64 `json:"vol"`
	} `json:"tick"`
}
type WsTrade struct {
	Channel   string `json:"ch"`
	Timestamp int64  `json:"ts"`
	Tick      struct {
		ID        int64 `json:"id"`
		Timestamp int64 `json:"ts"`
		Data      []struct {
			Amount    float64 `json:"amount"`
			Timestamp int64   `json:"ts"`
			ID        float64 `json:"id"`
			Price     float64 `json:"price"`
			Direction string  `json:"direction"`
		} `json:"data"`
	}
}
type WsAuthenticationRequest struct {
	Op               string `json:"op"`
	AccessKeyID      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
	ClientID         int64  `json:"cid,string,omitempty"`
}
type WsMessage struct {
	Raw []byte
	URL string
}
type WsAuthenticatedSubscriptionRequest struct {
	Op               string `json:"op"`
	AccessKeyID      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
	Topic            string `json:"topic"`
	ClientID         int64  `json:"cid,string,omitempty"`
}
type WsAuthenticatedAccountsListRequest struct {
	Op               string `json:"op"`
	AccessKeyID      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
	Topic            string `json:"topic"`
	Symbol           string `json:"symbol"`
	ClientID         int64  `json:"cid,string,omitempty"`
}
type WsAuthenticatedOrderDetailsRequest struct {
	Op               string `json:"op"`
	AccessKeyID      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
	Topic            string `json:"topic"`
	OrderID          string `json:"order-id"`
	ClientID         int64  `json:"cid,string,omitempty"`
}
type WsAuthenticatedOrdersListRequest struct {
	Op               string `json:"op"`
	AccessKeyID      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
	Topic            string `json:"topic"`
	States           string `json:"states"`
	AccountID        int64  `json:"account-id"`
	Symbol           string `json:"symbol"`
	ClientID         int64  `json:"cid,string,omitempty"`
}
type WsAuthenticatedAccountsResponse struct {
	WsResponse
	Data WsAuthenticatedAccountsResponseData `json:"data"`
}
type WsAuthenticatedAccountsResponseData struct {
	Event string                                    `json:"event"`
	List  []WsAuthenticatedAccountsResponseDataList `json:"list"`
}
type WsAuthenticatedAccountsResponseDataList struct {
	AccountID int64   `json:"account-id"`
	Currency  string  `json:"currency"`
	Type      string  `json:"type"`
	Balance   float64 `json:"balance,string"`
}
type WsAuthenticatedOrdersUpdateResponse struct {
	WsResponse
	Data WsAuthenticatedOrdersUpdateResponseData `json:"data"`
}
type WsAuthenticatedOrdersUpdateResponseData struct {
	UnfilledAmount   float64 `json:"unfilled-amount,string"`
	FilledAmount     float64 `json:"filled-amount,string"`
	Price            float64 `json:"price,string"`
	OrderID          int64   `json:"order-id"`
	Symbol           string  `json:"symbol"`
	MatchID          int64   `json:"match-id"`
	FilledCashAmount float64 `json:"filled-cash-amount,string"`
	Role             string  `json:"role"`
	OrderState       string  `json:"order-state"`
	OrderType        string  `json:"order-type"`
}
type WsAuthenticatedOrdersResponse struct {
	WsResponse
	Data []WsAuthenticatedOrdersResponseData `json:"data"`
}
type WsOldOrderUpdate struct {
	WsResponse
	Data WsAuthenticatedOrdersResponseData `json:"data"`
}
type WsAuthenticatedOrdersResponseData struct {
	SeqID            int64   `json:"seq-id"`
	OrderID          int64   `json:"order-id"`
	Symbol           string  `json:"symbol"`
	AccountID        int64   `json:"account-id"`
	OrderAmount      float64 `json:"order-amount,string"`
	OrderPrice       float64 `json:"order-price,string"`
	CreatedAt        int64   `json:"created-at"`
	OrderType        string  `json:"order-type"`
	OrderSource      string  `json:"order-source"`
	OrderState       string  `json:"order-state"`
	Role             string  `json:"role"`
	Price            float64 `json:"price,string"`
	FilledAmount     float64 `json:"filled-amount,string"`
	UnfilledAmount   float64 `json:"unfilled-amount,string"`
	FilledCashAmount float64 `json:"filled-cash-amount,string"`
	FilledFees       float64 `json:"filled-fees,string"`
}
type WsAuthenticatedAccountsListResponse struct {
	WsResponse
	Data []WsAuthenticatedAccountsListResponseData `json:"data"`
}
type WsAuthenticatedAccountsListResponseData struct {
	ID    int64                                         `json:"id"`
	Type  string                                        `json:"type"`
	State string                                        `json:"state"`
	List  []WsAuthenticatedAccountsListResponseDataList `json:"list"`
}
type WsAuthenticatedAccountsListResponseDataList struct {
	Currency string  `json:"currency"`
	Type     string  `json:"type"`
	Balance  float64 `json:"balance,string"`
}
type WsAuthenticatedOrdersListResponse struct {
	WsResponse
	Data []OrderInfo `json:"data"`
}
type WsAuthenticatedOrderDetailResponse struct {
	WsResponse
	Data OrderInfo `json:"data"`
}
type WsPong struct {
	Pong int64 `json:"pong"`
}
type wsKLineResponseThing struct {
	Data []struct {
		Amount float64 `json:"amount"`
		Close  float64 `json:"close"`
		Count  float64 `json:"count"`
		High   float64 `json:"high"`
		ID     int64   `json:"id"`
		Low    float64 `json:"low"`
		Open   float64 `json:"open"`
		Volume float64 `json:"vol"`
	} `json:"data"`
	Rep    string `json:"rep"`
	Status string `json:"status"`
}
type OrderBookDataRequestParamsType string
type SpotNewOrderRequestParamsType string
type TimeInterval string
