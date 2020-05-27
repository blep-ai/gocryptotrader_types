package communications

import (
	"errors"

	"github.com/thrasher-corp/gocryptotrader/communications/base"
	"github.com/thrasher-corp/gocryptotrader/communications/slack"
	"github.com/thrasher-corp/gocryptotrader/communications/smsglobal"
	"github.com/thrasher-corp/gocryptotrader/communications/smtpservice"
	"github.com/thrasher-corp/gocryptotrader/communications/telegram"
	"github.com/thrasher-corp/gocryptotrader/config"
)
type Communications struct {
	base.IComm
}
