package engine

import (
	"errors"
	"sync/atomic"

	"github.com/thrasher-corp/gocryptotrader/connchecker"
	"github.com/thrasher-corp/gocryptotrader/log"
)
type connectionManager struct {
	started int32
	stopped int32
	conn    *connchecker.Checker
}
