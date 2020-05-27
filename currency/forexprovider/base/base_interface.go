package base

import (
	"sync"
)

type IFXProvider interface {
	Setup(config Settings) error
	GetRates(baseCurrency, symbols string) (map[string]float64, error)
	GetName() string
	IsEnabled() bool
	IsPrimaryProvider() bool
	GetSupportedCurrencies() ([]string, error)
}
type FXHandler struct {
	Primary Provider
	Support []Provider
	mtx     sync.Mutex
}
type Provider struct {
	Provider            IFXProvider
	SupportedCurrencies []string
}
