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
	"github.com/blep-ai/gocryptotrader_types/common"
	"github.com/blep-ai/gocryptotrader_types/communications/base"
	"github.com/blep-ai/gocryptotrader_types/config"
	"github.com/blep-ai/gocryptotrader_types/log"
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
