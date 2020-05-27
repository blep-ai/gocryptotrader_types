package postgres

import (
	"context"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/volatiletech/null"
)

type (
	// ScriptSlice is an alias for a slice of pointers to Script.
	// This should generally be used opposed to []Script.
	ScriptSlice []*Script
	// ScriptHook is the signature for custom Script hook methods
	ScriptHook func(context.Context, boil.ContextExecutor, *Script) error

	scriptQuery struct {
		*queries.Query
	}
)
type Script struct {
	ID             string     `boil:"id" json:"id" toml:"id" yaml:"id"`
	ScriptID       string     `boil:"script_id" json:"script_id" toml:"script_id" yaml:"script_id"`
	ScriptName     string     `boil:"script_name" json:"script_name" toml:"script_name" yaml:"script_name"`
	ScriptPath     string     `boil:"script_path" json:"script_path" toml:"script_path" yaml:"script_path"`
	ScriptData     null.Bytes `boil:"script_data" json:"script_data,omitempty" toml:"script_data" yaml:"script_data,omitempty"`
	LastExecutedAt null.Time  `boil:"last_executed_at" json:"last_executed_at,omitempty" toml:"last_executed_at" yaml:"last_executed_at,omitempty"`
	CreatedAt      null.Time  `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *scriptR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L scriptL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}
type scriptR struct {
	ScriptExecutions ScriptExecutionSlice
}
type scriptL struct{}
