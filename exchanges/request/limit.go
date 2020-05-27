package request

import (
	"golang.org/x/time/rate"
)

type BasicLimit struct {
	r *rate.Limiter
}
type Limiter interface {
	Limit(EndpointLimit) error
}
type EndpointLimit int
