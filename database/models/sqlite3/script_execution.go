package sqlite3

import (
	"context"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
)

type (
	// ScriptExecutionSlice is an alias for a slice of pointers to ScriptExecution.
	// This should generally be used opposed to []ScriptExecution.
	ScriptExecutionSlice []*ScriptExecution
	// ScriptExecutionHook is the signature for custom ScriptExecution hook methods
	ScriptExecutionHook func(context.Context, boil.ContextExecutor, *ScriptExecution) error

	scriptExecutionQuery struct {
		*queries.Query
	}
)
type ScriptExecution struct {
	ID              int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	ScriptID        string `boil:"script_id" json:"script_id" toml:"script_id" yaml:"script_id"`
	ExecutionType   string `boil:"execution_type" json:"execution_type" toml:"execution_type" yaml:"execution_type"`
	ExecutionStatus string `boil:"execution_status" json:"execution_status" toml:"execution_status" yaml:"execution_status"`
	ExecutionTime   string `boil:"execution_time" json:"execution_time" toml:"execution_time" yaml:"execution_time"`

	R *scriptExecutionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L scriptExecutionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}
type scriptExecutionR struct {
	Script *Script
}
type scriptExecutionL struct{}
