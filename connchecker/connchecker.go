package connchecker
import (
	"net"
	"strings"
	"sync"
	"time"

	"github.com/blep-ai/gocryptotrader_types/log"
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
