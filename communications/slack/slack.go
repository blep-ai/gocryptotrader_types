package slack

import (
	"sync"

	"github.com/blep-ai/gocryptotrader_types/communications/base"
	"github.com/gorilla/websocket"
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
