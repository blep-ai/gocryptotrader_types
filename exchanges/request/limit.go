package request
import (
	"errors"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)
type BasicLimit struct {
	r *rate.Limiter
}
