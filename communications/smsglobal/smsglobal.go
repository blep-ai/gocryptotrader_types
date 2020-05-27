package smsglobal

import (
	"errors"
	"flag"
	"net/http"
	"net/url"
	"strings"

	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/communications/base"
	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/log"
)
type SMSGlobal struct {
	base.Base
	Contacts []Contact
	Username string
	Password string
	SendFrom string
}
