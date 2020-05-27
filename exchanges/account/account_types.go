package account

import (
	"sync"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/dispatch"
	"github.com/gofrs/uuid"
)

type Service struct {
	accounts map[string]*Account
	mux      *dispatch.Mux
	sync.Mutex
}
type Account struct {
	h  *Holdings
	ID uuid.UUID
}
type Holdings struct {
	Exchange string
	Accounts []SubAccount
}
type SubAccount struct {
	ID         string
	Currencies []Balance
}
type Balance struct {
	CurrencyName currency.Code
	TotalValue   float64
	Hold         float64
}
