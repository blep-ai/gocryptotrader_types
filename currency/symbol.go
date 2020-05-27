package currency

import "errors"

// GetSymbolByCurrencyName returns a currency symbol
func GetSymbolByCurrencyName(currency Code) (string, error) {
	result, ok := symbols[currency.Item]
	if !ok {
		return "", errors.New("currency symbol not found")
