package engine
import "github.com/gorilla/websocket"
type WebsocketClient struct {
	Hub           *WebsocketHub
	Conn          *websocket.Conn
	Authenticated bool
	authFailures  int
	Send          chan []byte
}
type WebsocketHub struct {
	Clients    map[*WebsocketClient]bool
	Broadcast  chan []byte
	Register   chan *WebsocketClient
	Unregister chan *WebsocketClient
}
type WebsocketEvent struct {
	Exchange  string `json:"exchange,omitempty"`
	AssetType string `json:"assetType,omitempty"`
	Event     string
	Data      interface{}
}
type WebsocketEventResponse struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
type WebsocketOrderbookTickerRequest struct {
	Exchange  string `json:"exchangeName"`
	Currency  string `json:"currency"`
	AssetType string `json:"assetType"`
}
type WebsocketAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
