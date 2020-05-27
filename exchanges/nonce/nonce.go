package nonce

import (
	"sync"
)

type Nonce struct {
	n int64
	m sync.Mutex
}
