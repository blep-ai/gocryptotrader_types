package base
import (
	"time"
)
type Base struct {
	Name      string
	Enabled   bool
	Verbose   bool
	Connected bool
}
type Event struct {
	Type    string
	Message string
}
type CommsStatus struct {
	Enabled   bool `json:"enabled"`
	Connected bool `json:"connected"`
}
