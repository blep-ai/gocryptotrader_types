package engine

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/blep-ai/gocryptotrader_types/log"
	"github.com/blep-ai/gocryptotrader_types/portfolio"
)
type portfolioManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
}
