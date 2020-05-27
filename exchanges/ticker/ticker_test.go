package ticker
import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/dispatch"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
)
