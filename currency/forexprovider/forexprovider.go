package forexprovider

import (
	"errors"

	"github.com/blep-ai/gocryptotrader_types/currency/forexprovider/base"
	currencyconverter "github.com/blep-ai/gocryptotrader_types/currency/forexprovider/currencyconverterapi"
	"github.com/blep-ai/gocryptotrader_types/currency/forexprovider/currencylayer"
	exchangerates "github.com/blep-ai/gocryptotrader_types/currency/forexprovider/exchangeratesapi.io"
	fixer "github.com/blep-ai/gocryptotrader_types/currency/forexprovider/fixer.io"
	"github.com/blep-ai/gocryptotrader_types/currency/forexprovider/openexchangerates"
)
type ForexProviders struct {
	base.FXHandler
}
