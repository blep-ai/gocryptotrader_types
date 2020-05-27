package indicators
import (
	"errors"
	"fmt"
	"math"
	"strings"

	objects "github.com/d5/tengo/v2"
	"github.com/thrasher-corp/gct-ta/indicators"
	"github.com/blep-ai/gocryptotrader_types/gctscript/modules"
	"github.com/blep-ai/gocryptotrader_types/gctscript/wrappers/validator"
)
type BBands struct {
	objects.Array
	Period               int
	STDDevUp, STDDevDown float64
	MAType               indicators.MaType
}
