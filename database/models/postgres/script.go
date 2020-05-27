package postgres
import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/thrasher-corp/sqlboiler/queries/qmhelper"
	"github.com/thrasher-corp/sqlboiler/strmangle"
	"github.com/volatiletech/null"
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
