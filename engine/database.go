package engine
import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/blep-ai/gocryptotrader_types/database"
	dbpsql "github.com/blep-ai/gocryptotrader_types/database/drivers/postgres"
	dbsqlite3 "github.com/blep-ai/gocryptotrader_types/database/drivers/sqlite3"
	"github.com/blep-ai/gocryptotrader_types/log"
	"github.com/thrasher-corp/sqlboiler/boil"
)
type databaseManager struct {
	started  int32
	stopped  int32
	shutdown chan struct{}
}
