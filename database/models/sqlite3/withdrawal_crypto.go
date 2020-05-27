package sqlite3

import (
	"context"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/volatiletech/null"
)

type (
	// WithdrawalCryptoSlice is an alias for a slice of pointers to WithdrawalCrypto.
	// This should generally be used opposed to []WithdrawalCrypto.
	WithdrawalCryptoSlice []*WithdrawalCrypto
	// WithdrawalCryptoHook is the signature for custom WithdrawalCrypto hook methods
	WithdrawalCryptoHook func(context.Context, boil.ContextExecutor, *WithdrawalCrypto) error

	withdrawalCryptoQuery struct {
		*queries.Query
	}
)
type WithdrawalCrypto struct {
	ID                  int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Address             string      `boil:"address" json:"address" toml:"address" yaml:"address"`
	AddressTag          null.String `boil:"address_tag" json:"address_tag,omitempty" toml:"address_tag" yaml:"address_tag,omitempty"`
	Fee                 float64     `boil:"fee" json:"fee" toml:"fee" yaml:"fee"`
	WithdrawalHistoryID string      `boil:"withdrawal_history_id" json:"withdrawal_history_id" toml:"withdrawal_history_id" yaml:"withdrawal_history_id"`

	R *withdrawalCryptoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L withdrawalCryptoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}
type withdrawalCryptoR struct {
	WithdrawalHistory *WithdrawalHistory
}
type withdrawalCryptoL struct{}
