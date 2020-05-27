package engine
import (
	"errors"
	"sync/atomic"

	"github.com/blep-ai/gocryptotrader_types/connchecker"
	"github.com/blep-ai/gocryptotrader_types/log"
)
type connectionManager struct {
	started int32
	stopped int32
	conn    *connchecker.Checker
}
