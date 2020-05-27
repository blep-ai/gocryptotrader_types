package engine
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/common/file"
	"github.com/blep-ai/gocryptotrader_types/currency"
	"github.com/blep-ai/gocryptotrader_types/dispatch"
	exchange "github.com/blep-ai/gocryptotrader_types/exchanges"
	"github.com/blep-ai/gocryptotrader_types/exchanges/account"
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
	"github.com/blep-ai/gocryptotrader_types/exchanges/orderbook"
	"github.com/blep-ai/gocryptotrader_types/exchanges/stats"
	"github.com/blep-ai/gocryptotrader_types/exchanges/ticker"
	"github.com/blep-ai/gocryptotrader_types/gctscript/vm"
	"github.com/blep-ai/gocryptotrader_types/log"
	"github.com/blep-ai/gocryptotrader_types/portfolio"
)
type RPCEndpoint struct {
	Started    bool
	ListenAddr string
}
