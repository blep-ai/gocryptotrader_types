package sqlite3

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/thrasher-corp/sqlboiler/queries/qmhelper"
	"github.com/thrasher-corp/sqlboiler/strmangle"
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
