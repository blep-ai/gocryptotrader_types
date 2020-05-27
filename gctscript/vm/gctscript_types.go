package vm

import (
	"time"
)

type Config struct {
	Enabled            bool          `json:"enabled"`
	ScriptTimeout      time.Duration `json:"timeout"`
	MaxVirtualMachines uint8         `json:"max_virtual_machines"`
	AllowImports       bool          `json:"allow_imports"`
	AutoLoad           []string      `json:"auto_load"`
	Verbose            bool          `json:"verbose"`
}
type Error struct {
	Script string
	Action string
	Cause  error
}
