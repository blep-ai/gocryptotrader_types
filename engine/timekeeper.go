package engine

import (
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/blep-ai/gocryptotrader_types/log"
	ntpclient "github.com/blep-ai/gocryptotrader_types/ntpclient"
)
type ntpManager struct {
	started       int32
	stopped       int32
	inititalCheck bool
	shutdown      chan struct{}
}
