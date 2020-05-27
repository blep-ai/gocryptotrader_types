package sqlite3
import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"regexp"

	"github.com/thrasher-corp/sqlboiler/boil"
)
type fKeyDestroyer struct {
	reader io.Reader
	buf    *bytes.Buffer
	rgx    *regexp.Regexp
}
