package bitmex

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/thrasher-corp/gocryptotrader/common"
)
type APIKeyParams struct {
	// API Key ID (public component).
	APIKeyID string `json:"apiKeyID,omitempty"`
}
type ChatGetParams struct {
	// ChannelID - [Optional] Leave blank for all.
	ChannelID float64 `json:"channelID,omitempty"`

	// Count - [Optional] Number of results to fetch.
	Count int32 `json:"count,omitempty"`

	// Reverse - [Optional] If true, will sort results newest first.
	Reverse bool `json:"reverse,omitempty"`

	// Start - [Optional] Starting ID for results.
	Start int32 `json:"start,omitempty"`
}
type ChatSendParams struct {
	// ChannelID - Channel to post to. Default 1 (English).
	ChannelID float64 `json:"channelID,omitempty"`

	// Message to send
	Message string `json:"message,omitempty"`
}
type GenericRequestParams struct {
	// Columns - [Optional] Array of column names to fetch. If omitted, will
	// return all columns.
	// NOTE that this method will always return item keys, even when not
	// specified, so you may receive more columns that you expect.
	Columns string `json:"columns,omitempty"`

	// Count - Number of results to fetch.
	Count int32 `json:"count,omitempty"`

	// EndTime - Ending date filter for results.
	EndTime string `json:"endTime,omitempty"`

	// Filter - Generic table filter. Send JSON key/value pairs, such as
	// `{"key": "value"}`. You can key on individual fields, and do more advanced
	// querying on timestamps. See the
	// [Timestamp Docs](https://testnet.bitmex.com/app/restAPI#Timestamp-Filters)
	// for more details.
	Filter string `json:"filter,omitempty"`

	// Reverse - If true, will sort results newest first.
	Reverse bool `json:"reverse,omitempty"`

	// Start - Starting point for results.
	Start int32 `json:"start,omitempty"`

	// StartTime - Starting date filter for results.
	StartTime string `json:"startTime,omitempty"`

	// Symbol - Instrument symbol. Send a bare series (e.g. XBU) to get data for
	// the nearest expiring contract in that series.
	// You can also send a timeframe, e.g. `XBU:monthly`. Timeframes are `daily`,
	// `weekly`, `monthly`, `quarterly`, and `biquarterly`.
	Symbol string `json:"symbol,omitempty"`
}
type LeaderboardGetParams struct {
	// MethodRanking - [Optional] type. Options: "notional", "ROE"
	Method string `json:"method,omitempty"`
}
type OrderNewParams struct {
	// ClientOrderID - [Optional] Client Order ID. This clOrdID will come back on the
	// order and any related executions.
	ClientOrderID string `json:"clOrdID,omitempty"`

	// ClientOrderLinkID - [Optional] Client Order Link ID for contingent orders.
	ClientOrderLinkID string `json:"clOrdLinkID,omitempty"`

	// ContingencyType - [Optional] contingency type for use with `clOrdLinkID`.
	// Valid options: OneCancelsTheOther, OneTriggersTheOther,
	// OneUpdatesTheOtherAbsolute, OneUpdatesTheOtherProportional.
	ContingencyType string `json:"contingencyType,omitempty"`

	// DisplayQuantity- [Optional] quantity to display in the book. Use 0 for a fully
	// hidden order.
	DisplayQuantity float64 `json:"displayQty,omitempty"`

	// ExecutionInstance - [Optional] execution instructions. Valid options:
	// ParticipateDoNotInitiate, AllOrNone, MarkPrice, IndexPrice, LastPrice,
	// Close, ReduceOnly, Fixed. 'AllOrNone' instruction requires `displayQty`
	// to be 0. 'MarkPrice', 'IndexPrice' or 'LastPrice' instruction valid for
	// 'Stop', 'StopLimit', 'MarketIfTouched', and 'LimitIfTouched' orders.
	ExecInst string `json:"execInst,omitempty"`

	// OrderType - Order type. Valid options: Market, Limit, Stop, StopLimit,
	// MarketIfTouched, LimitIfTouched, MarketWithLeftOverAsLimit, Pegged.
	// Defaults to 'Limit' when `price` is specified. Defaults to 'Stop' when
	// `stopPx` is specified. Defaults to 'StopLimit' when `price` and `stopPx`
	// are specified.
	OrderType string `json:"ordType,omitempty"`

	// OrderQuantity Order quantity in units of the instrument (i.e. contracts).
	OrderQuantity float64 `json:"orderQty,omitempty"`

	// PegOffsetValue - [Optional] trailing offset from the current price for
	// 'Stop', 'StopLimit', 'MarketIfTouched', and 'LimitIfTouched' orders; use a
	// negative offset for stop-sell orders and buy-if-touched orders. [Optional]
	// offset from the peg price for 'Pegged' orders.
	PegOffsetValue float64 `json:"pegOffsetValue,omitempty"`

	// PegPriceType - [Optional] peg price type. Valid options: LastPeg,
	// MidPricePeg, MarketPeg, PrimaryPeg, TrailingStopPeg.
	PegPriceType string `json:"pegPriceType,omitempty"`

	// Price - [Optional] limit price for 'Limit', 'StopLimit', and
	// 'LimitIfTouched' orders.
	Price float64 `json:"price,omitempty"`

	// Side - Order side. Valid options: Buy, Sell. Defaults to 'Buy' unless
	// `orderQty` or `simpleOrderQty` is negative.
	Side string `json:"side,omitempty"`

	// SimpleOrderQuantity - Order quantity in units of the underlying instrument
	// (i.e. Bitcoin).
	SimpleOrderQuantity float64 `json:"simpleOrderQty,omitempty"`

	// StopPrice - [Optional] trigger price for 'Stop', 'StopLimit',
	// 'MarketIfTouched', and 'LimitIfTouched' orders. Use a price below the
	// current price for stop-sell orders and buy-if-touched orders. Use
	// `execInst` of 'MarkPrice' or 'LastPrice' to define the current price used
	// for triggering.
	StopPx float64 `json:"stopPx,omitempty"`

	// Symbol - Instrument symbol. e.g. 'XBTUSD'.
	Symbol string `json:"symbol,omitempty"`

	// Text - [Optional] order annotation. e.g. 'Take profit'.
	Text string `json:"text,omitempty"`

	// TimeInForce - Valid options: Day, GoodTillCancel, ImmediateOrCancel,
	// FillOrKill. Defaults to 'GoodTillCancel' for 'Limit', 'StopLimit',
	// 'LimitIfTouched', and 'MarketWithLeftOverAsLimit' orders.
	TimeInForce string `json:"timeInForce,omitempty"`
}
type OrderAmendParams struct {
	// ClientOrderID - [Optional] new Client Order ID, requires `origClOrdID`.
	ClientOrderID string `json:"clOrdID,omitempty"`

	// LeavesQuantity - [Optional] leaves quantity in units of the instrument
	// (i.e. contracts). Useful for amending partially filled orders.
	LeavesQuantity int32 `json:"leavesQty,omitempty"`

	OrderID string `json:"orderID,omitempty"`

	// OrderQuantity - [Optional] order quantity in units of the instrument
	// (i.e. contracts).
	OrderQty int32 `json:"orderQty,omitempty"`

	// OrigClOrdID - Client Order ID. See POST /order.
	OrigClOrdID string `json:"origClOrdID,omitempty"`

	// PegOffsetValue - [Optional] trailing offset from the current price for
	// 'Stop', 'StopLimit', 'MarketIfTouched', and 'LimitIfTouched' orders; use a
	// negative offset for stop-sell orders and buy-if-touched orders. [Optional]
	// offset from the peg price for 'Pegged' orders.
	PegOffsetValue float64 `json:"pegOffsetValue,omitempty"`

	// Price - [Optional] limit price for 'Limit', 'StopLimit', and
	// 'LimitIfTouched' orders.
	Price float64 `json:"price,omitempty"`

	// SimpleLeavesQuantity - [Optional] leaves quantity in units of the underlying
	// instrument (i.e. Bitcoin). Useful for amending partially filled orders.
	SimpleLeavesQuantity float64 `json:"simpleLeavesQty,omitempty"`

	// SimpleOrderQuantity - [Optional] order quantity in units of the underlying
	// instrument (i.e. Bitcoin).
	SimpleOrderQuantity float64 `json:"simpleOrderQty,omitempty"`

	// StopPrice - [Optional] trigger price for 'Stop', 'StopLimit',
	// 'MarketIfTouched', and 'LimitIfTouched' orders. Use a price below the
	// current price for stop-sell orders and buy-if-touched orders.
	StopPx float64 `json:"stopPx,omitempty"`

	// Text - [Optional] amend annotation. e.g. 'Adjust skew'.
	Text string `json:"text,omitempty"`
}
type OrderCancelParams struct {
	// ClientOrderID - Client Order ID(s). See POST /order.
	ClientOrderID string `json:"clOrdID,omitempty"`

	// OrderID - Order ID(s).
	OrderID string `json:"orderID,omitempty"`

	// Text - [Optional] cancellation annotation. e.g. 'Spread Exceeded'.
	Text string `json:"text,omitempty"`
}
type OrderCancelAllParams struct {
	// Filter - [Optional] filter for cancellation. Use to only cancel some
	// orders, e.g. `{"side": "Buy"}`.
	Filter string `json:"filter,omitempty"`

	// Symbol - [Optional] symbol. If provided, only cancels orders for that
	// symbol.
	Symbol string `json:"symbol,omitempty"`

	// Text - [Optional] cancellation annotation. e.g. 'Spread Exceeded'
	Text string `json:"text,omitempty"`
}
type OrderAmendBulkParams struct {
	// Orders - An array of orders.
	Orders []OrderAmendParams `json:"orders,omitempty"`
}
type OrderNewBulkParams struct {
	// Orders - An array of orders.
	Orders []OrderNewParams `json:"orders,omitempty"`
}
type OrderCancelAllAfterParams struct {
	// Timeout in ms. Set to 0 to cancel this timer.
	Timeout float64 `json:"timeout,omitempty"`
}
type OrderClosePositionParams struct {
	// Price - [Optional] limit price.
	Price float64 `json:"price,omitempty"`

	// Symbol of position to close.
	Symbol string `json:"symbol,omitempty"`
}
type PositionGetParams struct {

	// Columns - Which columns to fetch. For example, send ["columnName"].
	Columns string `json:"columns,omitempty"`

	// Count - Number of rows to fetch.
	Count int32 `json:"count,omitempty"`

	// Filter - Table filter. For example, send {"symbol": "XBTUSD"}.
	Filter string `json:"filter,omitempty"`
}
type PositionIsolateMarginParams struct {
	// Enabled - True for isolated margin, false for cross margin.
	Enabled bool `json:"enabled,omitempty"`

	// Symbol - Position symbol to isolate.
	Symbol string `json:"symbol,omitempty"`
}
type PositionUpdateLeverageParams struct {
	// Leverage - Leverage value. Send a number between 0.01 and 100 to enable
	// isolated margin with a fixed leverage. Send 0 to enable cross margin.
	Leverage float64 `json:"leverage,omitempty"`

	// Symbol - Symbol of position to adjust.
	Symbol string `json:"symbol,omitempty"`
}
type PositionUpdateRiskLimitParams struct {
	// RiskLimit - New Risk Limit, in Satoshis.
	RiskLimit int64 `json:"riskLimit,omitempty"`

	// Symbol - Symbol of position to update risk limit on.
	Symbol string `json:"symbol,omitempty"`
}
type PositionTransferIsolatedMarginParams struct {
	// Amount - Amount to transfer, in Satoshis. May be negative.
	Amount int64 `json:"amount,omitempty"`

	// Symbol - Symbol of position to isolate.
	Symbol string `json:"symbol,omitempty"`
}
type QuoteGetBucketedParams struct {
	// BinSize - Time interval to bucket by. Available options: [1m,5m,1h,1d].
	BinSize string `json:"binSize,omitempty"`

	// Columns - Array of column names to fetch. If omitted, will return all
	// columns. NOTE that this method will always return item keys, even when not
	// specified, so you may receive more columns that you expect.
	Columns string `json:"columns,omitempty"`

	// Count - Number of results to fetch.
	Count int32 `json:"count,omitempty"`

	// EndTime - Ending date filter for results.
	EndTime string `json:"endTime,omitempty"`

	// Filter - Generic table filter. Send JSON key/value pairs, such as
	// `{"key": "value"}`. You can key on individual fields, and do more advanced
	// querying on timestamps. See the
	// [Timestamp Docs](https://testnet.bitmex.com/app/restAPI#Timestamp-Filters)
	// for more details.
	Filter string `json:"filter,omitempty"`

	// Partial - If true, will send in-progress (incomplete) bins for the current
	// time period.
	Partial bool `json:"partial,omitempty"`

	// Reverse - If true, will sort results newest first.
	Reverse bool `json:"reverse,omitempty"`

	// Start - Starting point for results.
	Start int32 `json:"start,omitempty"`

	// StartTime - Starting date filter for results.
	StartTime string `json:"startTime,omitempty"`

	// Symbol - Instrument symbol. Send a bare series (e.g. XBU) to get data for
	// the nearest expiring contract in that series.You can also send a timeframe,
	// e.g. `XBU:monthly`. Timeframes are `daily`, `weekly`, `monthly`,
	// `quarterly`, and `biquarterly`.
	Symbol string `json:"symbol,omitempty"`
}
type TradeGetBucketedParams struct {
	// BinSize - Time interval to bucket by. Available options: [1m,5m,1h,1d].
	BinSize string `json:"binSize,omitempty"`

	// Columns - Array of column names to fetch. If omitted, will return all
	// columns.
	// Note that this method will always return item keys, even when not
	// specified, so you may receive more columns that you expect.
	Columns string `json:"columns,omitempty"`

	// Count - Number of results to fetch.
	Count int32 `json:"count,omitempty"`

	// EndTime - Ending date filter for results.
	EndTime string `json:"endTime,omitempty"`

	// Filter - Generic table filter. Send JSON key/value pairs, such as
	// `{"key": "value"}`. You can key on individual fields, and do more advanced
	// querying on timestamps. See the
	// [Timestamp Docs](https://testnet.bitmex.com/app/restAPI#Timestamp-Filters)
	// for more details.
	Filter string `json:"filter,omitempty"`

	// Partial - If true, will send in-progress (incomplete) bins for the current
	// time period.
	Partial bool `json:"partial,omitempty"`

	// Reverse - If true, will sort results newest first.
	Reverse bool `json:"reverse,omitempty"`

	// Start - Starting point for results.
	Start int32 `json:"start,omitempty"`

	// StartTime - Starting date filter for results.
	StartTime string `json:"startTime,omitempty"`

	// Symbol - Instrument symbol. Send a bare series (e.g. XBU) to get data for
	// the nearest expiring contract in that series.You can also send a timeframe,
	// e.g. `XBU:monthly`. Timeframes are `daily`, `weekly`, `monthly`,
	// `quarterly`, and `biquarterly`.
	Symbol string `json:"symbol,omitempty"`
}
type UserUpdateParams struct {
	// Country - Country of residence.
	Country string `json:"country,omitempty"`

	// New Password string
	NewPassword string `json:"newPassword,omitempty"`

	// Confirmation string - must match
	NewPasswordConfirm string `json:"newPasswordConfirm,omitempty"`

	// old password string
	OldPassword string `json:"oldPassword,omitempty"`

	// PGP Public Key. If specified, automated emails will be sent with this key.
	PgpPubKey string `json:"pgpPubKey,omitempty"`

	// Username can only be set once. To reset, email support.
	Username string `json:"username,omitempty"`
}
type UserTokenParams struct {
	Token string `json:"token,omitempty"`
}
type UserCheckReferralCodeParams struct {
	ReferralCode string `json:"referralCode,omitempty"`
}
type UserConfirmTFAParams struct {
	// Token - Token from your selected TFA type.
	Token string `json:"token,omitempty"`

	// Type - Two-factor auth type. Supported types: 'GA' (Google Authenticator),
	// 'Yubikey'
	Type string `json:"type,omitempty"`
}
type UserCurrencyParams struct {
	Currency string `json:"currency,omitempty"`
}
type UserPreferencesParams struct {
	// Overwrite - If true, will overwrite all existing preferences.
	Overwrite bool `json:"overwrite,omitempty"`
	// Prefs - preferences
	Prefs string `json:"prefs,omitempty"`
}
type UserRequestWithdrawalParams struct {
	// Address - Destination Address.
	Address string `json:"address,omitempty"`

	// Amount - Amount of withdrawal currency.
	Amount float64 `json:"amount,omitempty"`

	// Currency - Currency you're withdrawing. Options: `XBt`
	Currency string `json:"currency,omitempty"`

	// Fee - Network fee for Bitcoin withdrawals. If not specified, a default
	// value will be calculated based on Bitcoin network conditions. You will have
	// a chance to confirm this via email.
	Fee float64 `json:"fee,omitempty"`

	// OtpToken - 2FA token. Required if 2FA is enabled on your account.
	OtpToken int64 `json:"otpToken,omitempty"`
}
type OrdersRequest struct {
	Symbol    string  `json:"symbol,omitempty"`
	Filter    string  `json:"filter,omitempty"`
	Columns   string  `json:"columns,omitempty"`
	Count     float64 `json:"count,omitempty"`
	Start     float64 `json:"start,omitempty"`
	Reverse   bool    `json:"reverse,omitempty"`
	StartTime string  `json:"startTime,omitempty"`
	EndTime   string  `json:"endTime,omitempty"`
}
