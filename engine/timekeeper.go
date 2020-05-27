package engine

import (
	"errors"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/thrasher-corp/gocryptotrader/log"
	ntpclient "github.com/thrasher-corp/gocryptotrader/ntpclient"
)
type ntpManager struct {
	started       int32
	stopped       int32
	inititalCheck bool
	shutdown      chan struct{}
}
