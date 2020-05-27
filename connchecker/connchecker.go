package connchecker

import (
	"net"
	"strings"
	"sync"
	"time"

	"github.com/thrasher-corp/gocryptotrader/log"
)
type Checker struct {
	DNSList       []string
	DomainList    []string
	CheckInterval time.Duration
	shutdown      chan struct{}
	wg            sync.WaitGroup
	connected     bool
	sync.Mutex
}
