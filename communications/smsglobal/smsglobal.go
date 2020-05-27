package smsglobal

import (
	"github.com/blep-ai/gocryptotrader_types/communications/base"
)

type SMSGlobal struct {
	base.Base
	Contacts []Contact
	Username string
	Password string
	SendFrom string
}
