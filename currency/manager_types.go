package currency

import (
	"sync"

	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
)
type PairsManager struct {
	RequestFormat   *PairFormat               `json:"requestFormat,omitempty"`
	ConfigFormat    *PairFormat               `json:"configFormat,omitempty"`
	UseGlobalFormat bool                      `json:"useGlobalFormat,omitempty"`
	LastUpdated     int64                     `json:"lastUpdated,omitempty"`
	AssetTypes      asset.Items               `json:"assetTypes"`
	Pairs           map[asset.Item]*PairStore `json:"pairs"`
	m               sync.Mutex
}
type PairStore struct {
	Enabled       Pairs       `json:"enabled"`
	Available     Pairs       `json:"available"`
	RequestFormat *PairFormat `json:"requestFormat,omitempty"`
	ConfigFormat  *PairFormat `json:"configFormat,omitempty"`
}
type PairFormat struct {
	Uppercase bool   `json:"uppercase"`
	Delimiter string `json:"delimiter,omitempty"`
	Separator string `json:"separator,omitempty"`
	Index     string `json:"index,omitempty"`
}
