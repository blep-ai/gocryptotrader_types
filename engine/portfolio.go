package engine

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/thrasher-corp/gocryptotrader/log"
	"github.com/thrasher-corp/gocryptotrader/portfolio"
)
type portfolioManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
}
