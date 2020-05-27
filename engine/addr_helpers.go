package engine
import (
	"errors"
	"strings"
	"sync"

	"github.com/blep-ai/gocryptotrader_types/currency"
)
type DepositAddressStore struct {
	m     sync.Mutex
	Store map[string]map[string]string
}
type DepositAddressManager struct {
	Store DepositAddressStore
}
