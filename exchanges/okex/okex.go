package okex
import (
	"fmt"
	"net/http"
	"time"

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/exchanges/okgroup"
)
type OKEX struct {
	okgroup.OKGroup
}
