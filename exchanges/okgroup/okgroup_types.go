package okgroup

import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
)
type GetAccountCurrenciesResponse struct {
	Name          string  `json:"name"`
	Currency      string  `json:"currency"`
	CanDeposit    int     `json:"can_deposit,string"`
	CanWithdraw   int     `json:"can_withdraw,string"`
	MinWithdrawal float64 `json:"min_withdrawal,string"`
}
type WalletInformationResponse struct {
	Available float64 `json:"available"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
	Hold      float64 `json:"hold"`
}
type TransferAccountFundsRequest struct {
	Currency     string  `json:"currency"`                // [required] token
	Amount       float64 `json:"amount"`                  // [required] Transfer amount
	From         int64   `json:"from"`                    // [required] the remitting account (0: sub account 1: spot 3: futures 4:C2C 5: margin 6: wallet 7:ETT 8:PiggyBank 9：swap)
	To           int64   `json:"to"`                      // [required] the beneficiary account(0: sub account 1:spot 3: futures 4:C2C 5: margin 6: wallet 7:ETT 8:PiggyBank 9 :swap)
	SubAccountID string  `json:"sub_account,omitempty"`   // [optional] sub account name
	InstrumentID int64   `json:"instrument_id,omitempty"` // [optional] margin token pair ID, for supported pairs only
}
type TransferAccountFundsResponse struct {
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	From       int64   `json:"from"`
	Result     bool    `json:"result"`
	To         int64   `json:"to"`
	TransferID int64   `json:"transfer_id"`
}
type AccountWithdrawRequest struct {
	Amount      float64 `json:"amount"`      // [required] withdrawal amount
	Currency    string  `json:"currency"`    // [required] token
	Destination int64   `json:"destination"` // [required] withdrawal address(2:OKCoin International 3:OKEx 4:others)
	Fee         float64 `json:"fee"`         // [required] Network transaction fee≥0. Withdrawals to OKCoin or OKEx are fee-free, please set as 0. Withdrawal to external digital asset address requires network transaction fee.
	ToAddress   string  `json:"to_address"`  // [required] verified digital asset address, email or mobile number,some digital asset address format is address+tag , eg: "ARDOR-7JF3-8F2E-QUWZ-CAN7F：123456"
	TradePwd    string  `json:"trade_pwd"`   // [required] fund password
}
type AccountWithdrawResponse struct {
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	Result       bool    `json:"result"`
	WithdrawalID int64   `json:"withdrawal_id"`
}
type GetAccountWithdrawalFeeResponse struct {
	Currency string  `json:"currency"`
	MinFee   float64 `json:"min_fee,string"`
	MaxFee   float64 `json:"max_fee,string"`
}
type WithdrawalHistoryResponse struct {
	Amount        float64   `json:"amount,string"`
	Currency      string    `json:"currency"`
	Fee           string    `json:"fee"`
	From          string    `json:"from"`
	Status        int64     `json:"status,string"`
	Timestamp     time.Time `json:"timestamp"`
	To            string    `json:"to"`
	TransactionID string    `json:"txid"`
	PaymentID     string    `json:"payment_id"`
	Tag           string    `json:"tag"`
}
type GetAccountBillDetailsRequest struct {
	Currency string `url:"currency,omitempty"` // [optional] token ,information of all tokens will be returned if the field is left blank
	Type     int64  `url:"type,omitempty"`     // [optional] 1:deposit 2:withdrawal 13:cancel withdrawal 18: into futures account 19: out of futures account 20:into sub account 21:out of sub account 28: claim 29: into ETT account 30: out of ETT account 31: into C2C account 32:out of C2C account 33: into margin account 34: out of margin account 37: into spot account 38: out of spot account
	From     int64  `url:"from,omitempty"`     // [optional] you would request pages after this page.
	To       int64  `url:"to,omitempty"`       // [optional] you would request pages before this page
	Limit    int64  `url:"limit,omitempty"`    // [optional] Number of results per request. Maximum 100. (default 100)
}
type GetAccountBillDetailsResponse struct {
	Amount    float64   `json:"amount"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	Fee       int64     `json:"fee"`
	LedgerID  int64     `json:"ledger_id"`
	Timestamp time.Time `json:"timestamp"`
	Typename  string    `json:"typename"`
}
type GetDepositAddressResponse struct {
	Address   string `json:"address"`
	Tag       string `json:"tag"`
	PaymentID string `json:"payment_id,omitempty"`
	Currency  string `json:"currency"`
}
type GetAccountDepositHistoryResponse struct {
	Amount        float64   `json:"amount,string"`
	Currency      string    `json:"currency"`
	From          string    `json:"from"`
	To            string    `json:"to"`
	Timestamp     time.Time `json:"timestamp"`
	Status        int64     `json:"status,string"`
	TransactionID string    `json:"txid"`
}
type GetSpotTradingAccountResponse struct {
	Available string `json:"available"`
	Balance   string `json:"balance"`
	Currency  string `json:"currency"`
	Frozen    string `json:"frozen"`
	Hold      string `json:"hold"`
	Holds     string `json:"holds"`
	ID        string `json:"id"`
}
type GetSpotBillDetailsForCurrencyRequest struct {
	Currency string `url:"-"`                      // [required] token
	From     int64  `url:"from,string,omitempty"`  // [optional] request page before(newer) this id.
	To       int64  `url:"to,string,omitempty"`    // [optional] request page after(older) this id.
	Limit    int64  `url:"limit,string,omitempty"` // [optional] number of results per request. Maximum 100.(default 100)
}
type GetSpotBillDetailsForCurrencyResponse struct {
	LedgerID         string          `json:"ledger_id"`
	Balance          string          `json:"balance"`
	CurrencyResponse string          `json:"currency"`
	Amount           string          `json:"amount"`
	Type             string          `json:"type"`
	TimeStamp        time.Time       `json:"timestamp"`
	Details          SpotBillDetails `json:"details"`
}
type SpotBillDetails struct {
	OrderID      string `json:"order_id"`
	InstrumentID string `json:"instrument_id"`
}
type PlaceOrderRequest struct {
	ClientOID     string `json:"client_oid,omitempty"` // the order ID customized by yourself
	Type          string `json:"type"`                 // limit / market(default: limit)
	Side          string `json:"side"`                 // buy or sell
	InstrumentID  string `json:"instrument_id"`        // trading pair
	MarginTrading string `json:"margin_trading"`       // margin trading
	OrderType     string `json:"order_type"`           // order type (0: Normal order (Unfilled and 0 imply normal limit order) 1: Post only 2: Fill or Kill 3: Immediate Or Cancel
	Size          string `json:"size"`
	Notional      string `json:"notional,omitempty"` //
	Price         string `json:"price,omitempty"`    // price (Limit order only)
}
type PlaceOrderResponse struct {
	ClientOid string `json:"client_oid"`
	OrderID   string `json:"order_id"`
	Result    bool   `json:"result"`
}
type CancelSpotOrderRequest struct {
	ClientOID    string `json:"client_oid,omitempty"` // the order ID created by yourself
	OrderID      int64  `json:"order_id,string"`      // order ID
	InstrumentID string `json:"instrument_id"`        // By providing this parameter, the corresponding order of a designated trading pair will be cancelled. If not providing this parameter, it will be back to error code.
}
type CancelSpotOrderResponse struct {
	ClientOID string `json:"client_oid"`
	OrderID   int64  `json:"order_id,string"`
	Result    bool   `json:"result"`
}
type CancelMultipleSpotOrdersRequest struct {
	OrderIDs     []int64 `json:"order_ids,omitempty"` // order ID. You may cancel up to 4 orders of a trading pair
	InstrumentID string  `json:"instrument_id"`       // by providing this parameter, the corresponding order of a designated trading pair will be cancelled. If not providing this parameter, it will be back to error code.
}
type CancelMultipleSpotOrdersResponse struct {
	ClientOID string `json:"client_oid"`
	OrderID   int64  `json:"order_id,string"`
	Result    bool   `json:"result"`
	Error     error  // Placeholder to store errors
}
type GetSpotOrdersRequest struct {
	Status string `url:"status"` // list the status of all orders (all, open, part_filled, canceling, filled, cancelled，ordering,failure)
	// （Multiple status separated by '|'，and '|' need encode to ' %7C'）
	InstrumentID string `url:"instrument_id"`          // trading pair ,information of all trading pair will be returned if the field is left blank
	From         int64  `url:"from,string,omitempty"`  // [optional] request page after this id (latest information) (eg. 1, 2, 3, 4, 5. There is only a 5 "from 4", while there are 1, 2, 3 "to 4")
	To           int64  `url:"to,string,omitempty"`    // [optional] request page after (older) this id.
	Limit        int64  `url:"limit,string,omitempty"` // [optional] number of results per request. Maximum 100. (default 100)
}
type GetSpotOrderResponse struct {
	FilledNotional float64   `json:"filled_notional,string"`
	FilledSize     float64   `json:"filled_size,string"`
	InstrumentID   string    `json:"instrument_id"`
	Notional       string    `json:"notional"`
	OrderID        string    `json:"order_id"`
	Price          float64   `json:"price,string"`
	Side           string    `json:"side"`
	Size           float64   `json:"size,string"`
	Status         string    `json:"status"`
	Timestamp      time.Time `json:"timestamp"`
	Type           string    `json:"type"`
}
type GetSpotOpenOrdersRequest struct {
	InstrumentID string `json:"instrument_id"`          // [optional] trading pair ,information of all trading pair will be returned if the field is left blank
	From         int64  `json:"from,string,omitempty"`  // [optional] request page after this id (latest information) (eg. 1, 2, 3, 4, 5. There is only a 5 "from 4", while there are 1, 2, 3 "to 4")
	To           int64  `json:"to,string,omitempty"`    // [optional] request page after (older) this id.
	Limit        int64  `json:"limit,string,omitempty"` // [optional] number of results per request. Maximum 100. (default 100)
}
type GetSpotOrderRequest struct {
	OrderID      string `url:"-"`             // [required] order ID
	InstrumentID string `url:"instrument_id"` // [required]trading pair
}
type GetSpotTransactionDetailsRequest struct {
	InstrumentID string `url:"instrument_id"`          // [required]list all transaction details of this instrument_id.
	OrderID      int64  `url:"order_id,string"`        // [required]list all transaction details of this order_id.
	From         int64  `url:"from,string,omitempty"`  // [optional] request page after this id (latest information) (eg. 1, 2, 3, 4, 5. There is only a 5 "from 4", while there are 1, 2, 3 "to 4")
	To           int64  `url:"to,string,omitempty"`    // [optional] request page after (older) this id.
	Limit        int64  `url:"limit,string,omitempty"` // [optional] number of results per request. Maximum 100. (default 100)
}
type GetSpotTransactionDetailsResponse struct {
	ExecType     string    `json:"exec_type"`
	Fee          string    `json:"fee"`
	InstrumentID string    `json:"instrument_id"`
	LedgerID     string    `json:"ledger_id"`
	OrderID      string    `json:"order_id"`
	Price        string    `json:"price"`
	Side         string    `json:"side"`
	Size         string    `json:"size"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetSpotTokenPairDetailsResponse struct {
	BaseCurrency  string `json:"base_currency"`
	InstrumentID  string `json:"instrument_id"`
	MinSize       string `json:"min_size"`
	QuoteCurrency string `json:"quote_currency"`
	SizeIncrement string `json:"size_increment"`
	TickSize      string `json:"tick_size"`
}
type GetOrderBookRequest struct {
	Size         int64   `url:"size,string,omitempty"`  // [optional] number of results per request. Maximum 200
	Depth        float64 `url:"depth,string,omitempty"` // [optional] the aggregation of the book. e.g . 0.1,0.001
	InstrumentID string  `url:"-"`                      // [required] trading pairs
}
type GetOrderBookResponse struct {
	Timestamp time.Time  `json:"timestamp"`
	Asks      [][]string `json:"asks"` // [[0]: "Price", [1]: "Size", [2]: "Num_orders"], ...
	Bids      [][]string `json:"bids"` // [[0]: "Price", [1]: "Size", [2]: "Num_orders"], ...
}
type GetSpotTokenPairsInformationResponse struct {
	BaseVolume24h  float64       `json:"base_volume_24h,string"`  // 24 trading volume of the base currency
	BestAsk        float64       `json:"best_ask,string"`         // best ask price
	BestBid        float64       `json:"best_bid,string"`         // best bid price
	High24h        float64       `json:"high_24h,string"`         // 24 hour high
	InstrumentID   currency.Pair `json:"instrument_id"`           // trading pair
	Last           float64       `json:"last,string"`             // last traded price
	Low24h         float64       `json:"low_24h,string"`          // 24 hour low
	Open24h        float64       `json:"open_24h,string"`         // 24 hour open
	QuoteVolume24h float64       `json:"quote_volume_24h,string"` // 24 trading volume of the quote currency
	Timestamp      time.Time     `json:"timestamp"`
}
type GetSpotFilledOrdersInformationRequest struct {
	InstrumentID string `url:"-"`                      // [required] trading pairs
	From         int64  `url:"from,string,omitempty"`  // [optional] number of results per request. Maximum 100. (default 100)
	To           int64  `url:"to,string,omitempty"`    // [optional] request page after (older) this id.
	Limit        int64  `url:"limit,string,omitempty"` // [optional] number of results per request. Maximum 100. (default 100)
}
type GetSpotFilledOrdersInformationResponse struct {
	Price     string    `json:"price"`
	Side      string    `json:"side"`
	Size      string    `json:"size"`
	Timestamp time.Time `json:"timestamp"`
	TradeID   string    `json:"trade_id"`
}
type GetSpotMarketDataRequest struct {
	Start        string `url:"start,omitempty"` // [optional] start time in ISO 8601
	End          string `url:"end,omitempty"`   // [optional] end time in ISO 8601
	Granularity  int64  `url:"granularity"`     // The granularity field must be one of the following values: {60, 180, 300, 900, 1800, 3600, 7200, 14400, 43200, 86400, 604800}.
	InstrumentID string `url:"-"`               // [required] trading pairs
}
type GetMarginAccountsResponse struct {
	InstrumentID     string `json:"instrument_id,omitempty"`
	LiquidationPrice string `json:"liquidation_price"`
	ProductID        string `json:"product_id,omitempty"`
	RiskRate         string `json:"risk_rate"`
	Currencies       map[string]MarginAccountInfo
}
type MarginAccountInfo struct {
	Available  float64 `json:"available,string"`
	Balance    float64 `json:"balance,string"`
	Borrowed   float64 `json:"borrowed,string"`
	Frozen     float64 `json:"frozen,string"`
	Hold       float64 `json:"hold,string"`
	Holds      float64 `json:"holds,string"`
	LendingFee float64 `json:"lending_fee,string"`
}
type GetMarginAccountSettingsResponse struct {
	InstrumentID string `json:"instrument_id"`
	ProductID    string `json:"product_id"`
	Currencies   map[string]MarginAccountSettingsInfo
}
type GetMarginBillDetailsRequest struct {
	InstrumentID string `url:"-"`               // [required] trading pair
	Type         int64  `url:"type,omitempty"`  // [optional] 1:deposit 2:withdrawal 13:cancel withdrawal 18: into futures account 19: out of futures account 20:into sub account 21:out of sub account 28: claim 29: into ETT account 30: out of ETT account 31: into C2C account 32:out of C2C account 33: into margin account 34: out of margin account 37: into spot account 38: out of spot account
	From         int64  `url:"from,omitempty"`  // [optional] you would request pages after this page.
	To           int64  `url:"to,omitempty"`    // [optional] you would request pages before this page
	Limit        int64  `url:"limit,omitempty"` // [optional] Number of results per request. Maximum 100. (default 100)
}
type MarginAccountSettingsInfo struct {
	Available     float64 `json:"available,string"`
	Leverage      float64 `json:"leverage,string"`
	LeverageRatio float64 `json:"leverage_ratio,string"`
	Rate          float64 `json:"rate,string"`
}
type GetMarginLoanHistoryRequest struct {
	InstrumentID string // [optional] Used when a specific currency response is desired
	Status       int64  `json:"status,string,omitempty"` // [optional] status(0: outstanding 1: repaid)
	From         int64  `json:"from,string,omitempty"`   // [optional] request page from(newer) this id.
	To           int64  `json:"to,string,omitempty"`     // [optional] request page to(older) this id.
	Limit        int64  `json:"limit,string,omitempty"`  // [optional] number of results per request. Maximum 100.(default 100)
}
type GetMarginLoanHistoryResponse struct {
	Amount           float64   `json:"amount,string"`
	BorrowID         int64     `json:"borrow_id"`
	CreatedAt        string    `json:"created_at"`
	Currency         string    `json:"currency"`
	ForceRepayTime   string    `json:"force_repay_time"`
	InstrumentID     string    `json:"instrument_id"`
	Interest         float64   `json:"interest,string"`
	LastInterestTime string    `json:"last_interest_time"`
	PaidInterest     float64   `json:"paid_interest,string"`
	ProductID        string    `json:"product_id"`
	Rate             float64   `json:"rate,string"`
	RepayAmount      string    `json:"repay_amount"`
	RepayInterest    string    `json:"repay_interest"`
	ReturnedAmount   float64   `json:"returned_amount,string"`
	Timestamp        time.Time `json:"timestamp"`
}
type OpenMarginLoanRequest struct {
	QuoteCurrency string  `json:"currency"`      // [required] Second currency eg BTC-USDT: USDT is quote
	InstrumentID  string  `json:"instrument_id"` // [required] Full pair BTC-USDT
	Amount        float64 `json:"amount,string"` // [required] Amount wanting to borrow
}
type OpenMarginLoanResponse struct {
	BorrowID int64 `json:"borrow_id"`
	Result   bool  `json:"result"`
}
type RepayMarginLoanRequest struct {
	Amount        float64 `json:"amount,string"` // [required] amount repaid
	BorrowID      float64 `json:"borrow_id"`     // [optional] borrow ID . all borrowed token under this trading pair will be repay if the field is left blank
	QuoteCurrency string  `json:"currency"`      // [required] Second currency eg BTC-USDT: USDT is quote
	InstrumentID  string  `json:"instrument_id"` // [required] Full pair BTC-USDT
}
type RepayMarginLoanResponse struct {
	RepaymentID int64 `json:"repayment_id"`
	Result      bool  `json:"result"`
}
type GetFuturesPositionsResponse struct {
	Holding [][]GetFuturePostionsDetails `json:"holding"`
	Result  bool                         `json:"result"`
}
type GetFuturesPositionsForCurrencyResponse struct {
	Holding []GetFuturePostionsDetails `json:"holding"`
	Result  bool                       `json:"result"`
}
type GetFuturePostionsDetails struct {
	CreatedAt            string `json:"created_at"`
	InstrumentID         string `json:"instrument_id"`
	Leverage             string `json:"leverage"`
	LiquidationPrice     string `json:"liquidation_price"`
	LongAvailQty         string `json:"long_avail_qty"`
	LongAvgCost          string `json:"long_avg_cost"`
	LongLeverage         string `json:"long_leverage"`
	LongLiquiPrice       string `json:"long_liqui_price"`
	LongMargin           string `json:"long_margin"`
	LongPnlRatio         string `json:"long_pnl_ratio"`
	LongQty              string `json:"long_qty"`
	LongSettlementPrice  string `json:"long_settlement_price"`
	MarginMode           string `json:"margin_mode"`
	RealisedPnl          string `json:"realised_pnl"`
	ShortAvailQty        string `json:"short_avail_qty"`
	ShortAvgCost         string `json:"short_avg_cost"`
	ShortLeverage        string `json:"short_leverage"`
	ShortLiquiPrice      string `json:"short_liqui_price"`
	ShortMargin          string `json:"short_margin"`
	ShortPnlRatio        string `json:"short_pnl_ratio"`
	ShortQty             string `json:"short_qty"`
	ShortSettlementPrice string `json:"short_settlement_price"`
	UpdatedAt            string `json:"updated_at"`
}
type FuturesAccountForAllCurrenciesResponse struct {
	Info struct {
		Currency map[string]FuturesCurrencyData
	} `json:"info"`
}
type FuturesCurrencyData struct {
	Contracts         []FuturesContractsData `json:"contracts,omitempty"`
	Equity            string                 `json:"equity,omitempty"`
	Margin            string                 `json:"margin,omitempty"`
	MarginMode        string                 `json:"margin_mode,omitempty"`
	MarginRatio       string                 `json:"margin_ratio,omitempty"`
	RealizedPnl       string                 `json:"realized_pnl,omitempty"`
	TotalAvailBalance string                 `json:"total_avail_balance,omitempty"`
	UnrealizedPnl     string                 `json:"unrealized_pnl,omitempty"`
}
type FuturesContractsData struct {
	AvailableQty      string `json:"available_qty"`
	FixedBalance      string `json:"fixed_balance"`
	InstrumentID      string `json:"instrument_id"`
	MarginForUnfilled string `json:"margin_for_unfilled"`
	MarginFrozen      string `json:"margin_frozen"`
	RealizedPnl       string `json:"realized_pnl"`
	UnrealizedPnl     string `json:"unrealized_pnl"`
}
type GetFuturesLeverageResponse struct {
	MarginMode      string `json:"margin_mode,omitempty"`
	Currency        string `json:"currency,omitempty"`
	Leverage        int64  `json:"leverage,omitempty"`
	LeveragePerCoin map[string]GetFuturesLeverageData
}
type GetFuturesLeverageData struct {
	LongLeverage  int64 `json:"long_leverage"`
	ShortLeverage int64 `json:"short_leverage"`
}
type SetFuturesLeverageRequest struct {
	Direction    string `json:"direction,omitempty"`     // opening side (long or short)
	InstrumentID string `json:"instrument_id,omitempty"` //  	Contract ID, e.g. "BTC-USD-180213"
	Leverage     int64  `json:"leverage,omitempty"`      //  	10x or 20x leverage
	Currency     string `json:"currency,omitempty"`
}
type SetFuturesLeverageResponse struct {
	Currency                 string `json:"currency"`
	Leverage                 int64  `json:"leverage"`
	MarginMode               string `json:"margin_mode"`
	Result                   string `json:"result"`
	Direction                string `json:"direction"`
	ShortLongDataPerContract map[string]SetFutureLeverageShortLongData
}
type SetFutureLeverageShortLongData struct {
	Long  int `json:"long"`
	Short int `json:"short"`
}
type PlaceFuturesOrderRequest struct {
	ClientOid    string  `json:"client_oid,omitempty"`         // [optional] 	the order ID customized by yourself
	InstrumentID string  `json:"instrument_id"`                // [required]   	Contract ID,e.g. "TC-USD-180213"
	Type         int64   `json:"type,string"`                  //  [required] 	1:open long 2:open short 3:close long 4:close short
	Price        float64 `json:"price,string"`                 //  [required] 	Price of each contract
	Size         int64   `json:"size,string"`                  //  [required] The buying or selling quantity
	MatchPrice   int64   `json:"match_price,string,omitempty"` // [optional] 	Order at best counter party price? (0:no 1:yes) the parameter is defaulted as 0. If it is set as 1, the price parameter will be ignored
	Leverage     int64   `json:"leverage,string"`              // [required]  	 	10x or 20x leverage
}
type PlaceFuturesOrderResponse struct {
	ClientOid     string `json:"client_oid"`
	ErrorCode     int    `json:"error_code"`
	ErrorMesssage string `json:"error_messsage"`
	OrderID       string `json:"order_id"`
	Result        bool   `json:"result"`
}
type PlaceFuturesOrderBatchRequest struct {
	InstrumentID string                                 `json:"instrument_id"` // [required] Contract ID, e.g."BTC-USD-180213"
	Leverage     int                                    `json:"leverage"`      // [required] 10x or 20x leverage
	OrdersData   []PlaceFuturesOrderBatchRequestDetails `json:"orders_data"`   // [required] the JSON word string for placing multiple orders, include：{client_oid type price size match_price}
}
type PlaceFuturesOrderBatchRequestDetails struct {
	ClientOid  string `json:"client_oid"`  // [required] To identify your order with the order ID set by you
	MatchPrice string `json:"match_price"` // undocumented
	Price      string `json:"price"`       // undocumented
	Size       string `json:"size"`        // undocumented
	Type       string `json:"type"`        // undocumented
}
type PlaceFuturesOrderBatchResponse struct {
	OrderInfo []PlaceFuturesOrderBatchResponseData `json:"order_info"`
	Result    bool                                 `json:"result"`
}
type PlaceFuturesOrderBatchResponseData struct {
	ClientOid    string  `json:"client_oid"`
	ErrorCode    int     `json:"error_code"`
	ErrorMessage string  `json:"error_message"`
	OrderID      float64 `json:"order_id"`
}
type CancelFuturesOrderRequest struct {
	OrderID      string `json:"order_id"`      // [required] Order ID
	InstrumentID string `json:"instrument_id"` // [required] Contract ID,e.g. "BTC-USD-180213"
}
type CancelFuturesOrderResponse struct {
	InstrumentID string `json:"instrument_id"`
	OrderID      string `json:"order_id"`
	Result       bool   `json:"result"`
}
type GetFuturesOrdersListRequest struct {
	InstrumentID string `url:"-"`                      // [required] Contract ID, e.g. "BTC-USD-180213"
	Status       int64  `url:"status,string"`          // [required] Order Status （-1 canceled; 0: pending, 1: partially filled, 2: fully filled, 6: open (pending partially + fully filled), 7: completed (canceled + fully filled))
	From         int64  `url:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `url:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `url:"limit,string,omitempty"` // [optional] Number of results per request. Maximum 100. (default 100)
}
type GetFuturesOrderListResponse struct {
	OrderInfo []GetFuturesOrderDetailsResponseData `json:"order_info"`
	Result    bool                                 `json:"result"`
}
type GetFuturesOrderDetailsResponseData struct {
	ContractVal  float64   `json:"contract_val,string"`
	Fee          float64   `json:"fee,string"`
	FilledQty    float64   `json:"filled_qty,string"`
	InstrumentID string    `json:"instrument_id"`
	Leverage     int64     `json:"leverage,string"` //  	Leverage value:10\20 default:10
	OrderID      int64     `json:"order_id,string"`
	Price        float64   `json:"price,string"`
	PriceAvg     float64   `json:"price_avg,string"`
	Size         float64   `json:"size,string"`
	Status       int64     `json:"status,string"` // Order Status （-1 canceled; 0: pending, 1: partially filled, 2: fully filled)
	Timestamp    time.Time `json:"timestamp"`
	Type         int64     `json:"type,string"` //  	Type (1: open long 2: open short 3: close long 4: close short)
}
type GetFuturesOrderDetailsRequest struct {
	OrderID      int64  `json:"order_id,string"` // [required] Order ID
	InstrumentID string `json:"instrument_id"`   // [required] Contract ID, e.g. "BTC-USD-180213"
}
type GetFuturesTransactionDetailsRequest struct {
	OrderID      int64  `json:"order_id,string"`        // [required] Order ID
	InstrumentID string `json:"instrument_id"`          // [required] Contract ID, e.g. "BTC-USD-180213"
	Status       int64  `json:"status,string"`          // [required] Order Status （-1 canceled; 0: pending, 1: partially filled, 2: fully filled, 6: open (pending partially + fully filled), 7: completed (canceled + fully filled))
	From         int64  `json:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `json:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `json:"limit,string,omitempty"` // [optional]  	Number of results per request. Maximum 100. (default 100)
}
type GetFuturesTransactionDetailsResponse struct {
	CreatedAt    string `json:"created_at"`
	ExecType     string `json:"exec_type"`
	Fee          string `json:"fee"`
	InstrumentID string `json:"instrument_id"`
	OrderID      string `json:"order_id"`
	OrderQty     string `json:"order_qty"`
	Price        string `json:"price"`
	Side         string `json:"side"`
	TradeID      string `json:"trade_id"`
}
type GetFuturesContractInformationResponse struct {
	ContractValue         float64 `json:"contract_val,string"`
	Alias                 string  `json:"alias"`
	BaseCurrency          string  `json:"base_currency"`
	SettlementCurrency    string  `json:"settlement_currency"`
	ContractValueCurrency string  `json:"contract_val_currency"`
	Delivery              string  `json:"delivery"`
	InstrumentID          string  `json:"instrument_id"`
	Listing               string  `json:"listing"`
	QuoteCurrency         string  `json:"quote_currency"`
	IsInverse             bool    `json:"is_inverse,string"`
	TickSize              float64 `json:"tick_size,string"`
	TradeIncrement        int64   `json:"trade_increment,string"`
	Underlying            string  `json:"underlying"`
	UnderlyingIndex       string  `json:"underlying_index"`
}
type GetFuturesTokenInfoResponse struct {
	BestAsk      float64   `json:"best_ask,string"`
	BestBid      float64   `json:"best_bid,string"`
	High24h      float64   `json:"high_24h,string"`
	InstrumentID string    `json:"instrument_id"`
	Last         float64   `json:"last,string"`
	Low24h       float64   `json:"low_24h,string"`
	Timestamp    time.Time `json:"timestamp"`
	Volume24h    float64   `json:"volume_24h,string"`
}
type GetFuturesFilledOrderRequest struct {
	InstrumentID string `url:"-"`                      // [required] Contract ID, e.g. "BTC-USD-180213"
	From         int64  `url:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `url:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `url:"limit,string,omitempty"` // [optional]  	Number of results per request. Maximum 100. (default 100)
}
type GetFuturesFilledOrdersResponse struct {
	Price     float64   `json:"price,string"`
	Qty       int64     `json:"qty,string"`
	Side      string    `json:"side"`
	Timestamp time.Time `json:"timestamp"`
	TradeID   string    `json:"trade_id"`
}
type GetFuturesMarketDateRequest struct {
	Start        string `url:"start,omitempty"`       // [optional] start time in ISO 8601
	End          string `url:"end,omitempty"`         // [optional] end time in ISO 8601
	Granularity  int64  `url:"granularity,omitempty"` // [optional] The granularity field must be one of the following values: {60, 180, 300, 900, 1800, 3600, 7200, 14400, 43200, 86400, 604800}.
	InstrumentID string `url:"-"`                     // [required] trading pairs
}
type GetFuturesHoldAmountResponse struct {
	Amount       float64   `json:"amount,string"`
	InstrumentID string    `json:"instrument_id"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetFuturesIndicesResponse struct {
	Index        float64   `json:"index,string"`
	InstrumentID string    `json:"instrument_id"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetFuturesExchangeRatesResponse struct {
	InstrumentID string    `json:"instrument_id"`
	Rate         float64   `json:"rate,string"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetFuturesEstimatedDeliveryPriceResponse struct {
	InstrumentID    string    `json:"instrument_id"`
	SettlementPrice float64   `json:"settlement_price,string"`
	Timestamp       time.Time `json:"timestamp"`
}
type GetFuturesOpenInterestsResponse struct {
	Amount       float64   `json:"amount,string"`
	InstrumentID string    `json:"instrument_id"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetFuturesCurrentPriceLimitResponse struct {
	Highest      float64   `json:"highest,string"`
	InstrumentID string    `json:"instrument_id"`
	Lowest       float64   `json:"lowest,string"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetFuturesCurrentMarkPriceResponse struct {
	MarkPrice    float64   `json:"mark_price,string"`
	InstrumentID string    `json:"instrument_id"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetFuturesForceLiquidatedOrdersRequest struct {
	InstrumentID string `url:"-"`                      // [required] Contract ID, e.g. "BTC-USD-180213"
	From         int64  `url:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `url:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `url:"limit,string,omitempty"` // [optional]  	Number of results per request. Maximum 100. (default 100)
	Status       string `url:"status,omitempty"`       // [optional] Status (0:unfilled orders in the recent 7 days 1:filled orders in the recent 7 days)
}
type GetFuturesForceLiquidatedOrdersResponse struct {
	Loss         float64 `json:"loss,string"`
	Size         int64   `json:"size,string"`
	Price        float64 `json:"price,string"`
	CreatedAt    string  `json:"created_at"`
	InstrumentID string  `json:"instrument_id"`
	Type         int64   `json:"type,string"`
}
type GetFuturesTagPriceResponse struct {
	MarkPrice    float64   `json:"mark_price"`
	InstrumentID string    `json:"instrument_id"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetSwapPostionsResponse struct {
	MarginMode string                           `json:"margin_mode"`
	Holding    []GetSwapPostionsResponseHolding `json:"holding"`
}
type GetSwapPostionsResponseHolding struct {
	AvailPosition    string    `json:"avail_position"`
	AvgCost          string    `json:"avg_cost"`
	InstrumentID     string    `json:"instrument_id"`
	Leverage         string    `json:"leverage"`
	LiquidationPrice string    `json:"liquidation_price"`
	Margin           string    `json:"margin"`
	Position         string    `json:"position"`
	RealizedPnl      string    `json:"realized_pnl"`
	SettlementPrice  string    `json:"settlement_price"`
	Side             string    `json:"side"`
	Timestamp        time.Time `json:"timestamp"`
}
type GetSwapAccountOfAllCurrencyResponse struct {
	Info []GetSwapAccountOfAllCurrencyResponseInfo `json:"info"`
}
type GetSwapAccountOfAllCurrencyResponseInfo struct {
	Equity            string    `json:"equity"`
	FixedBalance      string    `json:"fixed_balance"`
	TotalAvailBalance string    `json:"total_avail_balance"`
	Margin            string    `json:"margin"`
	RealizedPnl       string    `json:"realized_pnl"`
	UnrealizedPnl     string    `json:"unrealized_pnl"`
	MarginRatio       string    `json:"margin_ratio"`
	InstrumentID      string    `json:"instrument_id"`
	MarginFrozen      string    `json:"margin_frozen"`
	Timestamp         time.Time `json:"timestamp"`
	MarginMode        string    `json:"margin_mode"`
}
type GetSwapAccountSettingsOfAContractResponse struct {
	LongLeverage  float64 `json:"long_leverage,string"`
	MarginMode    string  `json:"margin_mode"`
	ShortLeverage float64 `json:"short_leverage,string"`
	InstrumentID  string  `json:"instrument_id"`
}
type SetSwapLeverageLevelOfAContractRequest struct {
	InstrumentID string `json:"instrument_id,omitempty"` // [required] Contract ID, e.g. BTC-USD-SWAP
	Leverage     int64  `json:"leverage,string"`         // [required] New leverage level from 1-100
	Side         int64  `json:"side,string"`             // [required] Side: 1.FIXED-LONG 2.FIXED-SHORT 3.CROSSED
}
type SetSwapLeverageLevelOfAContractResponse struct {
	InstrumentID  string `json:"instrument_id"`
	LongLeverage  int64  `json:"long_leverage,string"`
	MarginMode    string `json:"margin_mode"`
	ShortLeverage int64  `json:"short_leverage,string"`
}
type GetSwapBillDetailsResponse struct {
	LedgerID     string    `json:"ledger_id"`
	Amount       string    `json:"amount"`
	Type         string    `json:"type"`
	Fee          string    `json:"fee"`
	Timestamp    time.Time `json:"timestamp"`
	InstrumentID string    `json:"instrument_id"`
}
type PlaceSwapOrderRequest struct {
	ClientOID    string  `json:"client_oid,omitempty"`         // [optional] the order ID customized by yourself. 1-32 with digits and letter，The type of client_oid should be comprised of alphabets + numbers or only alphabets within 1 – 32 characters,Both uppercase and lowercase letters are supported
	Size         float64 `json:"size,string"`                  // [required] The buying or selling quantity
	Type         int64   `json:"type,string"`                  // [required] 1:open long 2:open short 3:close long 4:close short
	MatchPrice   int64   `json:"match_price,string,omitempty"` // [optional] Order at best counter party price? (0:no 1:yes)
	Price        float64 `json:"price,string"`                 // [required] Price of each contract
	InstrumentID string  `json:"instrument_id"`                // [required] Contract ID, e.g. BTC-USD-SWAP
}
type PlaceSwapOrderResponse struct {
	OrderID      string `json:"order_id"`
	ClientOID    int64  `json:"client_oid,string"`
	ErrorCode    int64  `json:"error_code,string"`
	ErrorMessage string `json:"error_message"`
	Result       bool   `json:"result,string"`
}
type PlaceMultipleSwapOrdersRequest struct {
	InstrumentID string                       `json:"instrument_id"` // [required] Contract ID, e.g. BTC-USD-SWAP
	Leverage     int64                        `json:"leverage"`      // [required] 10x or 20x leverage
	OrdersData   []PlaceMultipleSwapOrderData `json:"orders_data"`   // [required] the JSON word string for placing multiple orders, include：{client_oid type price size match_price}
}
type PlaceMultipleSwapOrderData struct {
	ClientOID  string `json:"client_oid"`  // [required] To identify your order with the order ID set by you
	Type       string `json:"type"`        // Undocumented
	Price      string `json:"price"`       // Undocumented
	Size       string `json:"size"`        // Undocumented
	MatchPrice string `json:"match_price"` // Undocumented
}
type PlaceMultipleSwapOrdersResponse struct {
	Result    bool                                  `json:"result,string"`
	OrderInfo []PlaceMultipleSwapOrdersResponseInfo `json:"order_info"`
}
type PlaceMultipleSwapOrdersResponseInfo struct {
	ErrorMessage string `json:"error_message"`
	ErrorCode    int64  `json:"error_code"`
	ClientOID    string `json:"client_oid"`
	OrderID      string `json:"order_id"`
}
type CancelSwapOrderRequest struct {
	OrderID      string `json:"order_id"`      // [required] Order ID
	InstrumentID string `json:"instrument_id"` // [required] Contract ID,e.g. BTC-USD-SWAP
}
type CancelSwapOrderResponse struct {
	Result       bool   `json:"result,string"`
	OrderID      string `json:"order_id"`
	InstrumentID string `json:"instrument_id"`
}
type CancelMultipleSwapOrdersRequest struct {
	InstrumentID string  `json:"instrument_id,omitempty"` // [required] The contract of the orders to be cancelled
	OrderIDs     []int64 `json:"order_ids"`               // [required] ID's of the orders canceled
}
type CancelMultipleSwapOrdersResponse struct {
	Result       bool     `json:"result,string"`
	OrderIDS     []string `json:"order_ids"`
	InstrumentID string   `json:"instrument_id"`
}
type GetSwapOrderListRequest struct {
	InstrumentID string `url:"-"`                      // [required] Contract ID, e.g. "BTC-USD-180213"
	Status       int64  `url:"status,string"`          // [required] Order Status （-1 canceled; 0: pending, 1: partially filled, 2: fully filled, 6: open (pending partially + fully filled), 7: completed (canceled + fully filled))
	From         int64  `url:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `url:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `url:"limit,string,omitempty"` // [optional] Number of results per request. Maximum 100. (default 100)
}
type GetSwapOrderListResponse struct {
	Result    bool                           `json:"result,string"`
	OrderInfo []GetSwapOrderListResponseData `json:"order_info"`
}
type GetSwapOrderListResponseData struct {
	ContractVal  float64   `json:"contract_val,string"`
	Fee          float64   `json:"fee,string"`
	FilledQty    float64   `json:"filled_qty,string"`
	InstrumentID string    `json:"instrument_id"`
	Leverage     int64     `json:"leverage,string"` //  	Leverage value:10\20 default:10
	OrderID      int64     `json:"order_id,string"`
	Price        float64   `json:"price,string"`
	PriceAvg     float64   `json:"price_avg,string"`
	Size         float64   `json:"size,string"`
	Status       int64     `json:"status,string"` // Order Status （-1 canceled; 0: pending, 1: partially filled, 2: fully filled)
	Timestamp    time.Time `json:"timestamp"`
	Type         int64     `json:"type,string"` //  	Type (1: open long 2: open short 3: close long 4: close short)
}
type GetSwapOrderDetailsRequest struct {
	InstrumentID string `json:"instrument_id"` // [required] Contract ID,e.g. BTC-USD-SWAP
	OrderID      string `json:"order_id"`      // [required] Order ID
}
type GetSwapTransactionDetailsRequest struct {
	InstrumentID string `json:"instrument_id"`          // [required] Contract ID, e.g. BTC-USD-SWAP
	OrderID      string `json:"order_id"`               // [required] Order ID
	From         int64  `json:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `json:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `json:"limit,string,omitempty"` // [optional] number of results per request. Maximum 100. (default 100)
}
type GetSwapTransactionDetailsResponse struct {
	TradeID      string    `json:"trade_id"`
	InstrumentID string    `json:"instrument_id"`
	OrderID      string    `json:"order_id"`
	Price        string    `json:"price"`
	OrderQty     string    `json:"order_qty"`
	Fee          string    `json:"fee"`
	Timestamp    time.Time `json:"timestamp"`
	ExecType     string    `json:"exec_type"`
	Side         string    `json:"side"`
}
type GetSwapContractInformationResponse struct {
	InstrumentID    string  `json:"instrument_id"`
	UnderlyingIndex string  `json:"underlying_index"`
	QuoteCurrency   string  `json:"quote_currency"`
	Coin            string  `json:"coin"`
	ContractVal     float64 `json:"contract_val,string"`
	Listing         string  `json:"listing"`
	Delivery        string  `json:"delivery"`
	SizeIncrement   float64 `json:"size_increment,string"`
	TickSize        float64 `json:"tick_size,string"`
}
type GetSwapOrderBookRequest struct {
	InstrumentID string  `url:"-"`
	Size         float64 `url:"size,string,omitempty"`
}
type GetSwapOrderBookResponse struct {
	Asks      [][]interface{} `json:"asks"` // eg [["411.3","16",5,4]] [[0: Price, 1: Size price, 2: number of force liquidated orders, 3: number of orders on the price]]
	Bids      [][]interface{} `json:"bids"` // eg [["411.3","16",5,4]] [[0: Price, 1: Size price, 2: number of force liquidated orders, 3: number of orders on the price]]
	Timestamp time.Time       `json:"timestamp"`
}
type GetAllSwapTokensInformationResponse struct {
	InstrumentID string    `json:"instrument_id"`
	Last         float64   `json:"last,string"`
	High24H      float64   `json:"high_24h,string"`
	Low24H       float64   `json:"low_24h,string"`
	BestBid      float64   `json:"best_bid,string"`
	BestAsk      float64   `json:"best_ask,string"`
	Volume24H    float64   `json:"volume_24h,string"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetSwapFilledOrdersDataRequest struct {
	InstrumentID string `url:"-"`                      // [required] Contract ID, e.g. "BTC-USD-SWAP
	From         int64  `url:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `url:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `url:"limit,string,omitempty"` // [optional] Number of results per request. Maximum 100. (default 100)
}
type GetSwapFilledOrdersDataResponse struct {
	TradeID   string    `json:"trade_id"`
	Price     float64   `json:"price,string"`
	Size      float64   `json:"size,string"`
	Side      string    `json:"side"`
	Timestamp time.Time `json:"timestamp"`
}
type GetSwapMarketDataRequest struct {
	Start        string `url:"start,omitempty"`       // [optional] start time in ISO 8601
	End          string `url:"end,omitempty"`         // [optional] end time in ISO 8601
	Granularity  int64  `url:"granularity,omitempty"` // The granularity field must be one of the following values: {60, 180, 300, 900, 1800, 3600, 7200, 14400, 43200, 86400, 604800}.
	InstrumentID string `url:"-"`                     // [required] trading pairs
}
type GetSwapIndecesResponse struct {
	InstrumentID string    `json:"instrument_id"`
	Index        float64   `json:"index,string"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetSwapExchangeRatesResponse struct {
	InstrumentID string    `json:"instrument_id"`
	Rate         float64   `json:"rate,string"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetSwapOpenInterestResponse struct {
	InstrumentID string    `json:"instrument_id"`
	Amount       float64   `json:"amount,string"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetSwapCurrentPriceLimitsResponse struct {
	InstrumentID string    `json:"instrument_id"`
	Highest      float64   `json:"highest,string"`
	Lowest       float64   `json:"lowest,string"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetSwapForceLiquidatedOrdersRequest struct {
	InstrumentID string `url:"-"`                      // [required] Contract ID, e.g. "BTC-USD-180213"
	From         int64  `url:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `url:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `url:"limit,string,omitempty"` // [optional]  	Number of results per request. Maximum 100. (default 100)
	Status       string `url:"status,omitempty"`       // [optional] Status (0:unfilled orders in the recent 7 days 1:filled orders in the recent 7 days)
}
type GetSwapForceLiquidatedOrdersResponse struct {
	Loss         float64 `json:"loss,string"`
	Size         int64   `json:"size,string"`
	Price        float64 `json:"price,string"`
	CreatedAt    string  `json:"created_at"`
	InstrumentID string  `json:"instrument_id"`
	Type         int64   `json:"type,string"`
}
type GetSwapOnHoldAmountForOpenOrdersResponse struct {
	InstrumentID string    `json:"instrument_id"`
	Amount       float64   `json:"amount,string"`
	Timestamp    time.Time `json:"timestamp"`
}
type GetSwapNextSettlementTimeResponse struct {
	InstrumentID string `json:"instrument_id"`
	FundingTime  string `json:"funding_time"`
}
type GetSwapMarkPriceResponse struct {
	InstrumentID string `json:"instrument_id"`
	MarkPrice    string `json:"mark_price"`
	Timstamp     string `json:"timstamp"`
}
type GetSwapFundingRateHistoryRequest struct {
	InstrumentID string `url:"ins-trument_id"`         // [required] Contract ID, e.g. "BTC-USD-SWAP
	From         int64  `url:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To           int64  `url:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit        int64  `url:"limit,string,omitempty"` // [optional] Number of results per request. Maximum 100.
}
type GetSwapFundingRateHistoryResponse struct {
	InstrumentID string  `json:"instrument_id"`
	FundingRate  float64 `json:"funding_rate,string,omitempty"`
	RealizedRate float64 `json:"realized_rate,string"`
	InterestRate float64 `json:"interest_rate,string"`
	FundingTime  string  `json:"funding_time"`
	FundingFee   float64 `json:"funding_fee,string,omitempty"`
}
type GetETTResponse struct {
	Currency  string  `json:"currency"`
	Balance   float64 `json:"balance"`
	Holds     float64 `json:"holds"`
	Available float64 `json:"available"`
}
type GetETTBillsDetailsResponse struct {
	LedgerID  int64   `json:"ledger_id"`
	Currency  string  `json:"currency"`
	Balance   float64 `json:"balance"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	CreatedAt string  `json:"created_at"`
	Details   int64   `json:"details"`
}
type PlaceETTOrderRequest struct {
	ClientOID     string  `json:"client_oid"`     // [optional]the order ID customized by yourself
	Type          int64   `json:"type"`           // Type of order (0:ETT subscription 1:subscribe with USDT 2:Redeem in USDT 3:Redeem in underlying)
	QuoteCurrency string  `json:"quote_currency"` // Subscription/redemption currency
	Amount        float64 `json:"amount"`         // Subscription amount. Required for usdt subscription
	Size          string  `json:"size"`           // Redemption size. Required for ETT subscription and redemption
	ETT           string  `json:"ett"`            // ETT name
}
type PlaceETTOrderResponse struct {
	ClientOID string `json:"client_oid"`
	OrderID   string `json:"type"`
	Result    bool   `json:"quote_currency"`
}
type GetETTOrderListRequest struct {
	ETT    string `url:"ett"`                    //  	[required] list specific ETT order
	Type   int64  `url:"type"`                   //  	[required]（1: subscription 2: redemption）
	Status int64  `url:"status,omitempty"`       // [optional]  	List the orders of the status (0:All 1:Unfilled 2:Filled 3:Canceled)
	From   int64  `url:"from,string,omitempty"`  // [optional] Request paging content for this page number.（Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	To     int64  `url:"to,string,omitempty"`    // [optional] Request page after (older) this pagination id. （Example: 1,2,3,4,5. From 4 we only have 4, to 4 we only have 3）
	Limit  int64  `url:"limit,string,omitempty"` // [optional] Number of results per request. Maximum 100.
}
type GetETTOrderListResponse struct {
	OrderID       string `json:"order_id"`
	Price         string `json:"price"`
	Size          string `json:"size"`
	Amount        string `json:"amount"`
	QuoteCurrency string `json:"quote_currency"`
	Ett           string `json:"ett"`
	Type          int64  `json:"type"`
	CreatedAt     string `json:"created_at"`
	Status        string `json:"status"`
}
type GetETTConstituentsResponse struct {
	NetValue     float64           `json:"net_value"`
	Ett          string            `json:"ett"`
	Constituents []ConstituentData `json:"constituents"`
}
type ConstituentData struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
type GetETTSettlementPriceHistoryResponse struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}
type SpotInstrument struct {
	BaseCurrency   string  `json:"base_currency"`
	BaseIncrement  float64 `json:"base_increment,string"`
	BaseMinSize    float64 `json:"base_min_size,string"`
	InstrumentID   string  `json:"instrument_id"`
	MinSize        float64 `json:"min_size,string"`
	ProductID      string  `json:"product_id"`
	QuoteCurrency  string  `json:"quote_currency"`
	QuoteIncrement float64 `json:"quote_increment,string"`
	SizeIncrement  float64 `json:"size_increment,string"`
	TickSize       float64 `json:"tick_size,string"`
}
type WebsocketEventRequest struct {
	Operation string   `json:"op"`   // 1--subscribe 2--unsubscribe 3--login
	Arguments []string `json:"args"` // args: the value is the channel name, which can be one or more channels
}
type WebsocketEventResponse struct {
	Event   string `json:"event"`
	Channel string `json:"channel,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type WebsocketDataResponse struct {
	Table  string        `json:"table"`
	Action string        `json:"action,omitempty"`
	Data   []interface{} `json:"data"`
}
type WebsocketTickerData struct {
	Table string `json:"table"`
	Data  []struct {
		BaseVolume24h  float64   `json:"base_volume_24h,string"`
		BestAsk        float64   `json:"best_ask,string"`
		BestAskSize    float64   `json:"best_ask_size,string"`
		BestBid        float64   `json:"best_bid,string"`
		BestBidSize    float64   `json:"best_bid_size,string"`
		High24h        float64   `json:"high_24h,string"`
		InstrumentID   string    `json:"instrument_id"`
		Last           float64   `json:"last,string"`
		LastQty        float64   `json:"last_qty,string"`
		Low24h         float64   `json:"low_24h,string"`
		Open24h        float64   `json:"open_24h,string"`
		QuoteVolume24h float64   `json:"quote_volume_24h,string"`
		Timestamp      time.Time `json:"timestamp"`
	} `json:"data"`
}
type WebsocketTradeResponse struct {
	Table string `json:"table"`
	Data  []struct {
		Price        float64   `json:"price,string"`
		Size         float64   `json:"size,string"`
		InstrumentID string    `json:"instrument_id"`
		Side         string    `json:"side"`
		Timestamp    time.Time `json:"timestamp"`
		TradeID      string    `json:"trade_id"`
	} `json:"data"`
}
type WebsocketCandleResponse struct {
	Table string `json:"table"`
	Data  []struct {
		Candle       []string `json:"candle"` // [0]timestamp, [1]open, [2]high, [3]low, [4]close, [5]volume, [6]currencyVolume
		InstrumentID string   `json:"instrument_id"`
	} `json:"data"`
}
type WebsocketOrderBooksData struct {
	Table  string               `json:"table"`
	Action string               `json:"action"`
	Data   []WebsocketOrderBook `json:"data"`
}
type WebsocketOrderBook struct {
	Checksum     int32           `json:"checksum,omitempty"`
	InstrumentID string          `json:"instrument_id"`
	Timestamp    time.Time       `json:"timestamp,omitempty"`
	Asks         [][]interface{} `json:"asks,omitempty"` // [0] Price, [1] Size, [2] Number of orders
	Bids         [][]interface{} `json:"bids,omitempty"` // [0] Price, [1] Size, [2] Number of orders
}
type WebsocketUserSwapPositionResponse struct {
	InstrumentID string                                 `json:"instrument_id"`
	Timestamp    time.Time                              `json:"timestamp,omitempty"`
	Holding      []WebsocketUserSwapPositionHoldingData `json:"holding,omitempty"`
}
type WebsocketUserSwapPositionHoldingData struct {
	AvailablePosition float64   `json:"avail_position,string,omitempty"`
	AverageCost       float64   `json:"avg_cost,string,omitempty"`
	Leverage          float64   `json:"leverage,string,omitempty"`
	LiquidationPrice  float64   `json:"liquidation_price,string,omitempty"`
	Margin            float64   `json:"margin,string,omitempty"`
	Position          float64   `json:"position,string,omitempty"`
	RealizedPnl       float64   `json:"realized_pnl,string,omitempty"`
	SettlementPrice   float64   `json:"settlement_price,string,omitempty"`
	Side              string    `json:"side,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}
type WebsocketSpotOrderResponse struct {
	Table string `json:"table"`
	Data  []struct {
		ClientOid      string    `json:"client_oid"`
		CreatedAt      time.Time `json:"created_at"`
		FilledNotional float64   `json:"filled_notional,string"`
		FilledSize     float64   `json:"filled_size,string"`
		InstrumentID   string    `json:"instrument_id"`
		LastFillPx     float64   `json:"last_fill_px,string"`
		LastFillQty    float64   `json:"last_fill_qty,string"`
		LastFillTime   time.Time `json:"last_fill_time"`
		MarginTrading  int64     `json:"margin_trading,string"`
		Notional       string    `json:"notional"`
		OrderID        string    `json:"order_id"`
		OrderType      int64     `json:"order_type,string"`
		Price          float64   `json:"price,string"`
		Side           string    `json:"side"`
		Size           float64   `json:"size,string"`
		State          int64     `json:"state,string"`
		Status         string    `json:"status"`
		Timestamp      time.Time `json:"timestamp"`
		Type           string    `json:"type"`
	} `json:"data"`
}
type WebsocketErrorResponse struct {
	Event     string `json:"event"`
	Message   string `json:"message"`
	ErrorCode int64  `json:"errorCode"`
}
