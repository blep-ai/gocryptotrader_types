package sqlite3
import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/strmangle"
)
type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}
type updateCache struct {
	query        string
	valueMapping []uint64
}
