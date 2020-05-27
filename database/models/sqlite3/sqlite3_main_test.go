package sqlite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteTester struct {
	dbConn *sql.DB

	dbName     string
	testDBName string
}
