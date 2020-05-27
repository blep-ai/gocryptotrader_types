package currency
import (
	"errors"
	"fmt"
	"sync"

	"github.com/blep-ai/gocryptotrader_types/log"
)
type ConversionRates struct {
	m   map[*Item]map[*Item]*float64
	mtx sync.Mutex
}
type Conversion struct {
	From        Code
	To          Code
	rate        *float64
	inverseRate *float64
	mtx         *sync.Mutex
}
