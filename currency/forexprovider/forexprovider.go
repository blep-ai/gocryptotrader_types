package forexprovider

import (
	"errors"

	"github.com/thrasher-corp/gocryptotrader/currency/forexprovider/base"
	currencyconverter "github.com/thrasher-corp/gocryptotrader/currency/forexprovider/currencyconverterapi"
	"github.com/thrasher-corp/gocryptotrader/currency/forexprovider/currencylayer"
	exchangerates "github.com/thrasher-corp/gocryptotrader/currency/forexprovider/exchangeratesapi.io"
	fixer "github.com/thrasher-corp/gocryptotrader/currency/forexprovider/fixer.io"
	"github.com/thrasher-corp/gocryptotrader/currency/forexprovider/openexchangerates"
)
type ForexProviders struct {
	base.FXHandler
}
