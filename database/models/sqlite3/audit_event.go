package sqlite3

import (
	"context"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
)

type (
	// AuditEventSlice is an alias for a slice of pointers to AuditEvent.
	// This should generally be used opposed to []AuditEvent.
	AuditEventSlice []*AuditEvent
	// AuditEventHook is the signature for custom AuditEvent hook methods
	AuditEventHook func(context.Context, boil.ContextExecutor, *AuditEvent) error

	auditEventQuery struct {
		*queries.Query
	}
)
type AuditEvent struct {
	ID         int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Type       string `boil:"type" json:"type" toml:"type" yaml:"type"`
	Identifier string `boil:"identifier" json:"identifier" toml:"identifier" yaml:"identifier"`
	Message    string `boil:"message" json:"message" toml:"message" yaml:"message"`
	CreatedAt  string `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *auditEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L auditEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}
type whereHelperstring struct{ field string }
type auditEventR struct {
}
type auditEventL struct{}
