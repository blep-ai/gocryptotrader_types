package coinbene

import (
	"golang.org/x/time/rate"
)

type RateLimit struct {
	ContractOrderbook             *rate.Limiter
	ContractTickers               *rate.Limiter
	ContractKline                 *rate.Limiter
	ContractTrades                *rate.Limiter
	ContractAccountInfo           *rate.Limiter
	ContractPositionInfo          *rate.Limiter
	ContractPlaceOrder            *rate.Limiter
	ContractCancelOrder           *rate.Limiter
	ContractGetOpenOrders         *rate.Limiter
	ContractOpenOrdersByPage      *rate.Limiter
	ContractGetOrderInfo          *rate.Limiter
	ContractGetClosedOrders       *rate.Limiter
	ContractGetClosedOrdersbyPage *rate.Limiter
	ContractCancelMultipleOrders  *rate.Limiter
	ContractGetOrderFills         *rate.Limiter
	ContractGetFundingRates       *rate.Limiter
	SpotPairs                     *rate.Limiter
	SpotPairInfo                  *rate.Limiter
	SpotOrderbook                 *rate.Limiter
	SpotTickerList                *rate.Limiter
	SpotSpecificTicker            *rate.Limiter
	SpotMarketTrades              *rate.Limiter
	// spotKline        // Not implemented yet
	// spotExchangeRate // Not implemented yet
	SpotAccountInfo       *rate.Limiter
	SpotAccountAssetInfo  *rate.Limiter
	SpotPlaceOrder        *rate.Limiter
	SpotBatchOrder        *rate.Limiter
	SpotQueryOpenOrders   *rate.Limiter
	SpotQueryClosedOrders *rate.Limiter
	SpotQuerySpecficOrder *rate.Limiter
	SpotQueryTradeFills   *rate.Limiter
	SpotCancelOrder       *rate.Limiter
	SpotCancelOrdersBatch *rate.Limiter
}
