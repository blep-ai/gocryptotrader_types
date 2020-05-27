package request

import (
	"io"
	"net/http"
	"time"

	"github.com/thrasher-corp/gocryptotrader/common/timedmutex"
	"github.com/thrasher-corp/gocryptotrader/exchanges/nonce"
)
type Requester struct {
	HTTPClient         *http.Client
	limiter            Limiter
	Name               string
	UserAgent          string
	maxRetries         int
	jobs               int32
	Nonce              nonce.Nonce
	disableRateLimiter int32
	backoff            Backoff
	retryPolicy        RetryPolicy
	timedLock          *timedmutex.TimedMutex
}
type Item struct {
	Method        string
	Path          string
	Headers       map[string]string
	Body          io.Reader
	Result        interface{}
	AuthRequest   bool
	NonceEnabled  bool
	Verbose       bool
	HTTPDebugging bool
	HTTPRecording bool
	IsReserved    bool
	Endpoint      EndpointLimit
}
