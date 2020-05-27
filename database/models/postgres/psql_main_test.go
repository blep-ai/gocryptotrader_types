package postgres

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kat-co/vala"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/thrasher-corp/goose"
	"github.com/thrasher-corp/sqlboiler/drivers/sqlboiler-psql/driver"
	"github.com/thrasher-corp/sqlboiler/randomize"
)
type pgTester struct {
	dbConn *sql.DB

	dbName  string
	host    string
	user    string
	pass    string
	sslmode string
	port    int

	pgPassFile string

	testDBName string
	skipSQLCmd bool
}
