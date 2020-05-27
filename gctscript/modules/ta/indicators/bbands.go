package indicators

import (
	objects "github.com/d5/tengo/v2"
	"github.com/thrasher-corp/gct-ta/indicators"
)

type BBands struct {
	objects.Array
	Period               int
	STDDevUp, STDDevDown float64
	MAType               indicators.MaType
}
