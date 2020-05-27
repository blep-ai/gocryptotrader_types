package wshandler
import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/log"
)
