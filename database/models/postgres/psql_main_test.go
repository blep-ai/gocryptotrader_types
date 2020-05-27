package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
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
