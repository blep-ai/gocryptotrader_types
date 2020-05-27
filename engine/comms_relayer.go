package engine

import (
	"github.com/blep-ai/gocryptotrader_types/communications"
	"github.com/blep-ai/gocryptotrader_types/communications/base"
)

type commsManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
	relayMsg chan base.Event
	comms    *communications.Communications
}
