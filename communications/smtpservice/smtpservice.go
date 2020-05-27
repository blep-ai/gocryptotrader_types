package smtpservice

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/thrasher-corp/gocryptotrader/communications/base"
	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/log"
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
