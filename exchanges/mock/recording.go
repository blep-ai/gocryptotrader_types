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

	"github.com/thrasher-corp/gocryptotrader/common/crypto"
	"github.com/thrasher-corp/gocryptotrader/common/file"
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
