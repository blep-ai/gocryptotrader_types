package main

import (
	"github.com/blep-ai/gocryptotrader_types/exchanges/asset"
)

type WebsocketEvent struct {
	Exchange  string `json:"exchange,omitempty"`
	AssetType string `json:"assetType,omitempty"`
	Event     string
	Data      interface{}
}
type WebsocketAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type WebsocketEventResponse struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
type WebsocketOrderbookTickerRequest struct {
	Exchange  string     `json:"exchangeName"`
	Currency  string     `json:"currency"`
	AssetType asset.Item `json:"assetType"`
}
