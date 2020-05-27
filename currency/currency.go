package currency

// GetDefaultExchangeRates returns the currency exchange rates based off the
// default fiat values
func GetDefaultExchangeRates() (Conversions, error) {
	return storage.GetDefaultForeignExchangeRates()
