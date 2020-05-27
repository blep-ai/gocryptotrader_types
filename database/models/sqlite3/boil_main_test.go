package sqlite3

import (
	"database/sql"
)

type tester interface {
	setup() error
	conn() (*sql.DB, error)
	teardown() error
}
