package bitfinex
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/common/crypto"
	"github.com/blep-ai/gocryptotrader_types/currency"
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
	"github.com/blep-ai/gocryptotrader_types/log"
	"github.com/blep-ai/gocryptotrader_types/portfolio/withdraw"
)
type Bitfinex struct {
	exchange.Base
	WebsocketConn              *wshandler.WebsocketConnection
	AuthenticatedWebsocketConn *wshandler.WebsocketConnection
	WebsocketSubdChannels      map[int]WebsocketChanInfo
}
