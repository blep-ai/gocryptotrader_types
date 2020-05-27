package engine

import (
	"github.com/blep-ai/gocryptotrader_types/connchecker"
)

type connectionManager struct {
	started int32
	stopped int32
	conn    *connchecker.Checker
}
