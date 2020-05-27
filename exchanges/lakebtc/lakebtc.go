package lakebtc
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/blep-ai/gocryptotrader_types/common/crypto"
	"github.com/blep-ai/gocryptotrader_types/currency"
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"github.com/blep-ai/gocryptotrader_types/log"
)
type LakeBTC struct {
	exchange.Base
	WebsocketConn
}
