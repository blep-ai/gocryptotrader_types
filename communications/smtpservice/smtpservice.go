package smtpservice
import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/blep-ai/gocryptotrader_types/communications/base"
	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/log"
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
