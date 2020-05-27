package gct

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	objects "github.com/d5/tengo/v2"
	"github.com/thrasher-corp/gocryptotrader/common/file"
	"github.com/thrasher-corp/gocryptotrader/gctscript/modules/ta/indicators"
	"github.com/thrasher-corp/gocryptotrader/log"
)
