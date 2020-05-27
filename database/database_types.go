package database

import (
	"database/sql"
	"sync"

	"github.com/blep-ai/gocryptotrader_types/database/drivers"
)

type Instance struct {
	SQL       *sql.DB
	DataPath  string
	Config    *Config
	Connected bool
	Mu        sync.RWMutex
}
type Config struct {
	Enabled                   bool   `json:"enabled"`
	Verbose                   bool   `json:"verbose"`
	Driver                    string `json:"driver"`
	drivers.ConnectionDetails `json:"connectionDetails"`
}
