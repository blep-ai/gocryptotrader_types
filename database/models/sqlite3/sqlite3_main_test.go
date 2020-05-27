package sqlite3
import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thrasher-corp/goose"
)
type sqliteTester struct {
	dbConn *sql.DB

	dbName     string
	testDBName string
}
