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
	"github.com/blep-ai/gocryptotrader_types/common/crypto"
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
	"github.com/blep-ai/gocryptotrader_types/log"
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
