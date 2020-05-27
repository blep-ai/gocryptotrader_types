package base

import (
	"testing"
)
type CommunicationProvider struct {
	ICommunicate

	isEnabled       bool
	isConnected     bool
	ConnectCalled   bool
	PushEventCalled bool
}
