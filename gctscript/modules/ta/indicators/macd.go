package indicators

import (
	objects "github.com/d5/tengo/v2"
)

type MACD struct {
	objects.Array
	Period, PeriodSlow, PeriodFast int
}
