package gemini

import (
	"golang.org/x/time/rate"
)

type RateLimit struct {
	Auth   *rate.Limiter
	UnAuth *rate.Limiter
}
