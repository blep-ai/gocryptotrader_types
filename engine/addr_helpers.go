package engine

import (
	"sync"
)

type DepositAddressStore struct {
	m     sync.Mutex
	Store map[string]map[string]string
}
type DepositAddressManager struct {
	Store DepositAddressStore
}
