package currency
import (
	"sync"
	"time"
)
type BaseCodes struct {
	Items          []*Item
	LastMainUpdate time.Time
	mtx            sync.Mutex
}
type Code struct {
	Item      *Item
	UpperCase bool
}
type Item struct {
	ID            int      `json:"id"`
	FullName      string   `json:"fullName"`
	Symbol        string   `json:"symbol"`
	Role          Role     `json:"role"`
	AssocChain    string   `json:"associatedBlockchain"`
	AssocExchange []string `json:"associatedExchanges"`
}
