package currency

import (
	"encoding/json"
	"fmt"
	"strings"
)
type Pair struct {
	Delimiter string `json:"delimiter"`
	Base      Code   `json:"base"`
	Quote     Code   `json:"quote"`
}
