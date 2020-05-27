package currency

import "testing"

func TestGetSymbolByCurrencyName(t *testing.T) {
	expected := "₩"
	actual, err := GetSymbolByCurrencyName(KPW)
