package postgres
import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/queries/qm"
	"github.com/thrasher-corp/sqlboiler/queries/qmhelper"
	"github.com/thrasher-corp/sqlboiler/strmangle"
	"github.com/volatiletech/null"
)
type WithdrawalFiat struct {
	ID                int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	WithdrawalFiatID  null.String `boil:"withdrawal_fiat_id" json:"withdrawal_fiat_id,omitempty" toml:"withdrawal_fiat_id" yaml:"withdrawal_fiat_id,omitempty"`
	BankName          string      `boil:"bank_name" json:"bank_name" toml:"bank_name" yaml:"bank_name"`
	BankAddress       string      `boil:"bank_address" json:"bank_address" toml:"bank_address" yaml:"bank_address"`
	BankAccountName   string      `boil:"bank_account_name" json:"bank_account_name" toml:"bank_account_name" yaml:"bank_account_name"`
	BankAccountNumber string      `boil:"bank_account_number" json:"bank_account_number" toml:"bank_account_number" yaml:"bank_account_number"`
	BSB               string      `boil:"bsb" json:"bsb" toml:"bsb" yaml:"bsb"`
	SwiftCode         string      `boil:"swift_code" json:"swift_code" toml:"swift_code" yaml:"swift_code"`
	Iban              string      `boil:"iban" json:"iban" toml:"iban" yaml:"iban"`
	BankCode          float64     `boil:"bank_code" json:"bank_code" toml:"bank_code" yaml:"bank_code"`

	R *withdrawalFiatR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L withdrawalFiatL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}
type withdrawalFiatR struct {
	WithdrawalFiat *WithdrawalHistory
}
type withdrawalFiatL struct{}
