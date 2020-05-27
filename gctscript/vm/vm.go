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
	"github.com/blep-ai/gocryptotrader_types/common/crypto"
	scriptevent "github.com/blep-ai/gocryptotrader_types/database/repository/script"
	"github.com/blep-ai/gocryptotrader_types/gctscript/modules/loader"
	"github.com/blep-ai/gocryptotrader_types/gctscript/wrappers/validator"
	"github.com/blep-ai/gocryptotrader_types/log"
	"github.com/volatiletech/null"
)
