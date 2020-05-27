package base

import (
	"sync"
)

type FXHandler struct {
	Primary Provider
	Support []Provider
	mtx     sync.Mutex
}
type Provider struct {
	Provider            IFXProvider
	SupportedCurrencies []string
}
