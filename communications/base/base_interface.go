package base

import (
	"github.com/blep-ai/gocryptotrader_types/config"
)

type ICommunicate interface {
	Setup(config *config.CommunicationsConfig)
	Connect() error
	PushEvent(Event) error
	IsEnabled() bool
	IsConnected() bool
	GetName() string
}
