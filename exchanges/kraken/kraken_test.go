package kraken
import (
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/core"
	"github.com/blep-ai/gocryptotrader_types/currency"
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/order"
	"github.com/blep-ai/gocryptotrader_types/exchanges/sharedtestvalues"
	"github.com/blep-ai/gocryptotrader_types/exchanges/websocket/wshandler"
	"github.com/blep-ai/gocryptotrader_types/portfolio/withdraw"
)
