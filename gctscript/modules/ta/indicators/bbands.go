package indicators

import (
	"errors"
	"fmt"
	"math"
	"strings"

	objects "github.com/d5/tengo/v2"
	"github.com/thrasher-corp/gct-ta/indicators"
	"github.com/thrasher-corp/gocryptotrader/gctscript/modules"
	"github.com/thrasher-corp/gocryptotrader/gctscript/wrappers/validator"
)
type BBands struct {
	objects.Array
	Period               int
	STDDevUp, STDDevDown float64
	MAType               indicators.MaType
}