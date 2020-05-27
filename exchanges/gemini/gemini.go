package gemini
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/common/crypto"
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
	"github.com/blep-ai/gocryptotrader_types/log"
)
type Gemini struct {
	WebsocketConn              *wshandler.WebsocketConnection
	AuthenticatedWebsocketConn *wshandler.WebsocketConnection
	exchange.Base
	Role              string
	RequiresHeartBeat bool
}
