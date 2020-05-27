package currency
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/blep-ai/gocryptotrader_types/common/file"
	"github.com/blep-ai/gocryptotrader_types/currency/coinmarketcap"
	"github.com/blep-ai/gocryptotrader_types/currency/forexprovider"
	"github.com/blep-ai/gocryptotrader_types/currency/forexprovider/base"
	"github.com/blep-ai/gocryptotrader_types/log"
)
