package connchecker

import (
	"sync"
	"time"
)

type Checker struct {
	DNSList       []string
	DomainList    []string
	CheckInterval time.Duration
	shutdown      chan struct{}
	wg            sync.WaitGroup
	connected     bool
	sync.Mutex
}
