package okgroup

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/thrasher-corp/gocryptotrader/common/crypto"
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/request"
	"github.com/thrasher-corp/gocryptotrader/exchanges/websocket/wshandler"
	"github.com/thrasher-corp/gocryptotrader/log"
)
type OKGroup struct {
	exchange.Base
	ExchangeName  string
	WebsocketConn *wshandler.WebsocketConnection
	// Spot and contract market error codes as per https://www.okex.com/rest_request.html
	ErrorCodes map[string]error
	// Stores for corresponding variable checks
	ContractTypes         []string
	CurrencyPairsDefaults []string
	ContractPosition      []string
	Types                 []string
	// URLs to be overridden by implementations of OKGroup
	APIURL       string
	APIVersion   string
	WebsocketURL string
}
