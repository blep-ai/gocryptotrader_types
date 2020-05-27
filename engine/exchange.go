package engine

import (
	"sync"

	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
)

type exchangeManager struct {
	m         sync.Mutex
	exchanges map[string]exchange.IBotExchange
}
