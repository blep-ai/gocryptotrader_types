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
type WithdrawalCrypto struct {
	ID                 int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	WithdrawalCryptoID null.String `boil:"withdrawal_crypto_id" json:"withdrawal_crypto_id,omitempty" toml:"withdrawal_crypto_id" yaml:"withdrawal_crypto_id,omitempty"`
	Address            string      `boil:"address" json:"address" toml:"address" yaml:"address"`
	AddressTag         null.String `boil:"address_tag" json:"address_tag,omitempty" toml:"address_tag" yaml:"address_tag,omitempty"`
	Fee                float64     `boil:"fee" json:"fee" toml:"fee" yaml:"fee"`

	R *withdrawalCryptoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L withdrawalCryptoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}
type withdrawalCryptoR struct {
	WithdrawalCrypto *WithdrawalHistory
}
type withdrawalCryptoL struct{}