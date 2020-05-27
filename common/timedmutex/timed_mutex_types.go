package timedmutex

import (
	"sync"
	"time"
)

type TimedMutex struct {
	mtx       sync.Mutex
	timerLock sync.RWMutex
	timer     *time.Timer
	duration  time.Duration
}
