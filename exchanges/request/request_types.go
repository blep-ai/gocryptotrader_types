package request

import (
	"io"
	"net/http"
	"time"

	"github.com/blep-ai/gocryptotrader_types/common/timedmutex"
	"github.com/blep-ai/gocryptotrader_types/exchanges/nonce"
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
type Backoff func(n int) time.Duration
type RetryPolicy func(resp *http.Response, err error) (bool, error)
type RequesterOption func(*Requester)
