package mock

import (
	"encoding/json"
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
