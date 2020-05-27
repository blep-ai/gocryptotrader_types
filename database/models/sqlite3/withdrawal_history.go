package sqlite3

import (
	"github.com/volatiletech/null"
)

type WithdrawalHistory struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Exchange     string      `boil:"exchange" json:"exchange" toml:"exchange" yaml:"exchange"`
	ExchangeID   string      `boil:"exchange_id" json:"exchange_id" toml:"exchange_id" yaml:"exchange_id"`
	Status       string      `boil:"status" json:"status" toml:"status" yaml:"status"`
	Currency     string      `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	Amount       float64     `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Description  null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	WithdrawType int64       `boil:"withdraw_type" json:"withdraw_type" toml:"withdraw_type" yaml:"withdraw_type"`
	CreatedAt    string      `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    string      `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *withdrawalHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L withdrawalHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}
type withdrawalHistoryR struct {
	WithdrawalCryptos WithdrawalCryptoSlice
	WithdrawalFiats   WithdrawalFiatSlice
}
type withdrawalHistoryL struct{}
