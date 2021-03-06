package wshandler

type testStruct struct {
	Error error
	WC    WebsocketConnection
}
type testRequest struct {
	Event        string          `json:"event"`
	RequestID    int64           `json:"reqid,omitempty"`
	Pairs        []string        `json:"pair"`
	Subscription testRequestData `json:"subscription,omitempty"`
}
type testRequestData struct {
	Name     string `json:"name,omitempty"`
	Interval int64  `json:"interval,omitempty"`
	Depth    int64  `json:"depth,omitempty"`
}
type testResponse struct {
	RequestID int64 `json:"reqid,omitempty"`
}
