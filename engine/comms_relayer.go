package engine

import (
	"errors"
	"sync/atomic"

	"github.com/thrasher-corp/gocryptotrader/communications"
	"github.com/thrasher-corp/gocryptotrader/communications/base"
	"github.com/thrasher-corp/gocryptotrader/log"
)
type commsManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
	relayMsg chan base.Event
	comms    *communications.Communications
}
