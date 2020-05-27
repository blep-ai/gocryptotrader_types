package smsglobal
import (
	"errors"
	"flag"
	"net/http"
	"net/url"
	"strings"

	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/communications/base"
	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/log"
)
type SMSGlobal struct {
	base.Base
	Contacts []Contact
	Username string
	Password string
	SendFrom string
}
