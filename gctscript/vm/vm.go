package vm

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/d5/tengo/v2"
	"github.com/gofrs/uuid"
	"github.com/thrasher-corp/gocryptotrader/common/crypto"
	scriptevent "github.com/thrasher-corp/gocryptotrader/database/repository/script"
	"github.com/thrasher-corp/gocryptotrader/gctscript/modules/loader"
	"github.com/thrasher-corp/gocryptotrader/gctscript/wrappers/validator"
	"github.com/thrasher-corp/gocryptotrader/log"
	"github.com/volatiletech/null"
)
