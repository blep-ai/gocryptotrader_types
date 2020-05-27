package base

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/blep-ai/gocryptotrader_types/common"
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
