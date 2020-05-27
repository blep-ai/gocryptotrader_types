package engine

import (
	"fmt"
	"path/filepath"
	"sync/atomic"

	"github.com/thrasher-corp/gocryptotrader/gctscript/vm"
	"github.com/thrasher-corp/gocryptotrader/log"
)
type gctScriptManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
}
