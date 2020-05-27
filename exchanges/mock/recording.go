package mock
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/blep-ai/gocryptotrader_types/common/crypto"
	"github.com/blep-ai/gocryptotrader_types/common/file"
)
type HTTPResponse struct {
	Data        json.RawMessage     `json:"data"`
	QueryString string              `json:"queryString"`
	BodyParams  string              `json:"bodyParams"`
	Headers     map[string][]string `json:"headers"`
}
type Exclusion struct {
	Headers   []string `json:"headers"`
	Variables []string `json:"variables"`
}
