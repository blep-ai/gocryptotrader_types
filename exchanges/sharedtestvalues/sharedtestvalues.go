package sharedtestvalues

import "time"

// This package is only to be referenced in test files
const (
	// WebsocketResponseDefaultTimeout used in websocket testing
	// Defines wait time for receiving websocket response before cancelling
	WebsocketResponseDefaultTimeout = (3 * time.Second)
