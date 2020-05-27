package request

import (
	"golang.org/x/time/rate"
)

type BasicLimit struct {
	r *rate.Limiter
}
