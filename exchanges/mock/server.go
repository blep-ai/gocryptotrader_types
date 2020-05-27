package mock

type VCRMock struct {
	Routes map[string]map[string][]HTTPResponse `json:"routes"`
}
