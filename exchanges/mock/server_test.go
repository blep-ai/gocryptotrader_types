package mock

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/thrasher-corp/gocryptotrader/common"
)
type responsePayload struct {
	Price    float64 `json:"price"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
