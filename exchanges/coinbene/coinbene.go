package coinbene
import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/common/crypto"
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
)
type Coinbene struct {
	exchange.Base
	WebsocketConn *wshandler.WebsocketConnection
}
