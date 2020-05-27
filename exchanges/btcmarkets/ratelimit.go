package btcmarkets
import (
	"time"

	"github.com/blep-ai/gocryptotrader_types/exchanges/request"
	"golang.org/x/time/rate"
)
type RateLimit struct {
	Auth            *rate.Limiter
	UnAuth          *rate.Limiter
	OrderPlacement  *rate.Limiter
	BatchOrders     *rate.Limiter
	WithdrawRequest *rate.Limiter
	CreateNewReport *rate.Limiter
}
