package request

import (
	"golang.org/x/time/rate"
)

type GlobalLimitTest struct {
	Auth   *rate.Limiter
	UnAuth *rate.Limiter
}
