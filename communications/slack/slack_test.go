package slack

import (
	"encoding/json"
	"testing"

	"github.com/blep-ai/gocryptotrader_types/communications/base"
	"github.com/blep-ai/gocryptotrader_types/config"
)
type group struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
}
