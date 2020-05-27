package nonce

import (
	"strconv"
	"sync"
)
type Nonce struct {
	n int64
	m sync.Mutex
}
