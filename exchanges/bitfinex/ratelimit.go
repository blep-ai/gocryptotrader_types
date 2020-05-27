package bitfinex

import (
	"errors"
	"time"

	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	"golang.org/x/time/rate"
)
type RateLimit struct {
	PlatformStatus       *rate.Limiter
	TickerBatch          *rate.Limiter
	Ticker               *rate.Limiter
	Trade                *rate.Limiter
	Orderbook            *rate.Limiter
	Stats                *rate.Limiter
	Candle               *rate.Limiter
	Configs              *rate.Limiter
	Status               *rate.Limiter
	Liquid               *rate.Limiter
	LeaderBoard          *rate.Limiter
	MarketAveragePrice   *rate.Limiter
	Fx                   *rate.Limiter
	AccountWalletBalance *rate.Limiter
	AccountWalletHistory *rate.Limiter
	// Orders -
	RetrieveOrder  *rate.Limiter
	SubmitOrder    *rate.Limiter
	UpdateOrder    *rate.Limiter
	CancelOrder    *rate.Limiter
	OrderBatch     *rate.Limiter
	CancelBatch    *rate.Limiter
	OrderHistory   *rate.Limiter
	GetOrderTrades *rate.Limiter
	GetTrades      *rate.Limiter
	GetLedgers     *rate.Limiter
	// Positions -
	GetAccountMarginInfo       *rate.Limiter
	GetActivePositions         *rate.Limiter
	ClaimPosition              *rate.Limiter
	GetPositionHistory         *rate.Limiter
	GetPositionAudit           *rate.Limiter
	UpdateCollateralOnPosition *rate.Limiter
	// Margin funding -
	GetActiveFundingOffers   *rate.Limiter
	SubmitFundingOffer       *rate.Limiter
	CancelFundingOffer       *rate.Limiter
	CancelAllFundingOffer    *rate.Limiter
	CloseFunding             *rate.Limiter
	FundingAutoRenew         *rate.Limiter
	KeepFunding              *rate.Limiter
	GetOffersHistory         *rate.Limiter
	GetFundingLoans          *rate.Limiter
	GetFundingLoanHistory    *rate.Limiter
	GetFundingCredits        *rate.Limiter
	GetFundingCreditsHistory *rate.Limiter
	GetFundingTrades         *rate.Limiter
	GetFundingInfo           *rate.Limiter
	// Account actions
	GetUserInfo               *rate.Limiter
	TransferBetweenWallets    *rate.Limiter
	GetDepositAddress         *rate.Limiter
	Withdrawal                *rate.Limiter
	GetMovements              *rate.Limiter
	GetAlertList              *rate.Limiter
	SetPriceAlert             *rate.Limiter
	DeletePriceAlert          *rate.Limiter
	GetBalanceForOrdersOffers *rate.Limiter
	UserSettingsWrite         *rate.Limiter
	UserSettingsRead          *rate.Limiter
	UserSettingsDelete        *rate.Limiter
	// Account V1 endpoints
	GetAccountFees    *rate.Limiter
	GetWithdrawalFees *rate.Limiter
	GetAccountSummary *rate.Limiter
	NewDepositAddress *rate.Limiter
	GetKeyPermissions *rate.Limiter
	GetMarginInfo     *rate.Limiter
	GetAccountBalance *rate.Limiter
	WalletTransfer    *rate.Limiter
	WithdrawV1        *rate.Limiter
	OrderV1           *rate.Limiter
	OrderMulti        *rate.Limiter
	StatsV1           *rate.Limiter
	Fundingbook       *rate.Limiter
	Lends             *rate.Limiter
}
