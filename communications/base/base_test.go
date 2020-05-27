package base

type CommunicationProvider struct {
	ICommunicate

	isEnabled       bool
	isConnected     bool
	ConnectCalled   bool
	PushEventCalled bool
}
