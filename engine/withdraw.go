package engine

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"
	withdrawDataStore "github.com/thrasher-corp/gocryptotrader/database/repository/withdraw"
	"github.com/thrasher-corp/gocryptotrader/gctrpc"
	"github.com/thrasher-corp/gocryptotrader/log"
	"github.com/thrasher-corp/gocryptotrader/portfolio/withdraw"
)
