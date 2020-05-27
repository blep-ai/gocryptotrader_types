package gct

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	objects "github.com/d5/tengo/v2"
	"github.com/blep-ai/gocryptotrader_types/common/file"
	"github.com/blep-ai/gocryptotrader_types/gctscript/modules/ta/indicators"
	"github.com/blep-ai/gocryptotrader_types/log"
)
