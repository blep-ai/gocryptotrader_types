package currencyconverter

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/currency/forexprovider/base"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"github.com/blep-ai/gocryptotrader_types/log"
)
