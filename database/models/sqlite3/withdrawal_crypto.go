package sqlite3

import (
	"github.com/volatiletech/null"
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
