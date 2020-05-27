package engine
import (
	"fmt"
	"path/filepath"
	"sync/atomic"

	"github.com/blep-ai/gocryptotrader_types/gctscript/vm"
	"github.com/blep-ai/gocryptotrader_types/log"
)
type gctScriptManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
}
