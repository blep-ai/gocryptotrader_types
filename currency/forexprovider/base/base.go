package base

import (
	"time"
)

type Settings struct {
	Name             string        `json:"name"`
	Enabled          bool          `json:"enabled"`
	Verbose          bool          `json:"verbose"`
	RESTPollingDelay time.Duration `json:"restPollingDelay"`
	APIKey           string        `json:"apiKey"`
	APIKeyLvl        int           `json:"apiKeyLvl"`
	PrimaryProvider  bool          `json:"primaryProvider"`
}
type Base struct {
	Settings `json:"settings"`
}
