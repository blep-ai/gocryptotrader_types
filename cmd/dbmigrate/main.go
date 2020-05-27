package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/core"
	"github.com/blep-ai/gocryptotrader_types/database"
	dbPSQL "github.com/blep-ai/gocryptotrader_types/database/drivers/postgres"
	dbsqlite3 "github.com/blep-ai/gocryptotrader_types/database/drivers/sqlite3"
	"github.com/blep-ai/gocryptotrader_types/database/repository"
	"github.com/thrasher-corp/goose"
)
