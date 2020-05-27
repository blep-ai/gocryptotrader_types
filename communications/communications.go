package communications

import (
	"errors"

	"github.com/blep-ai/gocryptotrader_types/communications/base"
	"github.com/blep-ai/gocryptotrader_types/communications/slack"
	"github.com/blep-ai/gocryptotrader_types/communications/smsglobal"
	"github.com/blep-ai/gocryptotrader_types/communications/smtpservice"
	"github.com/blep-ai/gocryptotrader_types/communications/telegram"
	"github.com/blep-ai/gocryptotrader_types/config"
)
type Communications struct {
	base.IComm
}
