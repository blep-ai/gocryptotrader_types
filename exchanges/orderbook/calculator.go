package orderbook

type WhaleBombResult struct {
	Amount               float64
	MinimumPrice         float64
	MaximumPrice         float64
	PercentageGainOrLoss float64
	Orders               orderSummary
	Status               string
}
type OrderSimulationResult WhaleBombResult
type orderSummary []Item
type ByPrice orderSummary
