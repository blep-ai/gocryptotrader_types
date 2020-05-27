package order
import (
	"errors"
	"fmt"
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
)
type Submit struct {
	ImmediateOrCancel bool
	HiddenOrder       bool
	FillOrKill        bool
	PostOnly          bool
	Leverage          string
	Price             float64
	Amount            float64
	LimitPriceUpper   float64
	LimitPriceLower   float64
	TriggerPrice      float64
	TargetAmount      float64
	ExecutedAmount    float64
	RemainingAmount   float64
	Fee               float64
	Exchange          string
	InternalOrderID   string
	ID                string
	AccountID         string
	ClientID          string
	WalletAddress     string
	Type              Type
	Side              Side
	Status            Status
	AssetType         asset.Item
	Date              time.Time
	LastUpdated       time.Time
	Pair              currency.Pair
	Trades            []TradeHistory
}
type SubmitResponse struct {
	IsOrderPlaced bool
	FullyMatched  bool
	OrderID       string
}
type Modify struct {
	ImmediateOrCancel bool
	HiddenOrder       bool
	FillOrKill        bool
	PostOnly          bool
	Leverage          string
	Price             float64
	Amount            float64
	LimitPriceUpper   float64
	LimitPriceLower   float64
	TriggerPrice      float64
	TargetAmount      float64
	ExecutedAmount    float64
	RemainingAmount   float64
	Fee               float64
	Exchange          string
	InternalOrderID   string
	ID                string
	AccountID         string
	ClientID          string
	WalletAddress     string
	Type              Type
	Side              Side
	Status            Status
	AssetType         asset.Item
	Date              time.Time
	LastUpdated       time.Time
	Pair              currency.Pair
	Trades            []TradeHistory
}
type ModifyResponse struct {
	OrderID string
}
type Detail struct {
	ImmediateOrCancel bool
	HiddenOrder       bool
	FillOrKill        bool
	PostOnly          bool
	Leverage          string
	Price             float64
	Amount            float64
	LimitPriceUpper   float64
	LimitPriceLower   float64
	TriggerPrice      float64
	TargetAmount      float64
	ExecutedAmount    float64
	RemainingAmount   float64
	Fee               float64
	Exchange          string
	InternalOrderID   string
	ID                string
	AccountID         string
	ClientID          string
	WalletAddress     string
	Type              Type
	Side              Side
	Status            Status
	AssetType         asset.Item
	Date              time.Time
	LastUpdated       time.Time
	Pair              currency.Pair
	Trades            []TradeHistory
}
type Cancel struct {
	Price         float64
	Amount        float64
	Exchange      string
	ID            string
	AccountID     string
	ClientID      string
	WalletAddress string
	Type          Type
	Side          Side
	Status        Status
	AssetType     asset.Item
	Date          time.Time
	Pair          currency.Pair
	Trades        []TradeHistory
}
type CancelAllResponse struct {
	Status map[string]string
}
type TradeHistory struct {
	Price       float64
	Amount      float64
	Fee         float64
	Exchange    string
	TID         string
	Description string
	Type        Type
	Side        Side
	Timestamp   time.Time
	IsMaker     bool
}
type GetOrdersRequest struct {
	Type       Type
	Side       Side
	StartTicks time.Time
	EndTicks   time.Time
	// Currencies Empty array = all currencies. Some endpoints only support
	// singular currency enquiries
	Pairs []currency.Pair
}
type ClassificationError struct {
	Exchange string
	OrderID  string
	Err      error
}
