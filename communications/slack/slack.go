package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/thrasher-corp/gocryptotrader/common"
	"github.com/thrasher-corp/gocryptotrader/communications/base"
	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/log"
)
type Slack struct {
	base.Base

	TargetChannel     string
	VerificationToken string

	TargetChannelID string
	Details         Response
	ReconnectURL    string
	WebsocketConn   *websocket.Conn
	Connected       bool
	Shutdown        bool
	sync.Mutex
}
