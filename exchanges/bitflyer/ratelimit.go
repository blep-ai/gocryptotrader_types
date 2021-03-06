package bitflyer

import (
	"golang.org/x/time/rate"
)

type RateLimit struct {
	Auth   *rate.Limiter
	UnAuth *rate.Limiter

	// Send a New Order
	// Submit New Parent Order (Special order)
	// Cancel All Orders
	Order     *rate.Limiter
	LowVolume *rate.Limiter
}
