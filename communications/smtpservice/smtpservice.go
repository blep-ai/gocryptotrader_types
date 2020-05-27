package smtpservice

import (
	"github.com/blep-ai/gocryptotrader_types/communications/base"
)

type SMTPservice struct {
	base.Base
	Host            string
	Port            string
	AccountName     string
	AccountPassword string
	From            string
	RecipientList   string
}
