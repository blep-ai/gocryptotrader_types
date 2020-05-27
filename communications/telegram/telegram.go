package telegram

import (
	"github.com/blep-ai/gocryptotrader_types/communications/base"
)

type Telegram struct {
	base.Base
	initConnected     bool
	Token             string
	Offset            int64
	AuthorisedClients []int64
}
