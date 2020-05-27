package gctrpc

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

type GetInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetInfoResponse struct {
	Uptime               string                  `protobuf:"bytes,1,opt,name=uptime,proto3" json:"uptime,omitempty"`
	AvailableExchanges   int64                   `protobuf:"varint,2,opt,name=available_exchanges,json=availableExchanges,proto3" json:"available_exchanges,omitempty"`
	EnabledExchanges     int64                   `protobuf:"varint,3,opt,name=enabled_exchanges,json=enabledExchanges,proto3" json:"enabled_exchanges,omitempty"`
	DefaultForexProvider string                  `protobuf:"bytes,4,opt,name=default_forex_provider,json=defaultForexProvider,proto3" json:"default_forex_provider,omitempty"`
	DefaultFiatCurrency  string                  `protobuf:"bytes,5,opt,name=default_fiat_currency,json=defaultFiatCurrency,proto3" json:"default_fiat_currency,omitempty"`
	SubsystemStatus      map[string]bool         `protobuf:"bytes,6,rep,name=subsystem_status,json=subsystemStatus,proto3" json:"subsystem_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	RpcEndpoints         map[string]*RPCEndpoint `protobuf:"bytes,7,rep,name=rpc_endpoints,json=rpcEndpoints,proto3" json:"rpc_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}
type GetCommunicationRelayersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type CommunicationRelayer struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Connected            bool     `protobuf:"varint,2,opt,name=connected,proto3" json:"connected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetCommunicationRelayersResponse struct {
	CommunicationRelayers map[string]*CommunicationRelayer `protobuf:"bytes,1,rep,name=communication_relayers,json=communicationRelayers,proto3" json:"communication_relayers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral  struct{}                         `json:"-"`
	XXX_unrecognized      []byte                           `json:"-"`
	XXX_sizecache         int32                            `json:"-"`
}
type GenericSubsystemRequest struct {
	Subsystem            string   `protobuf:"bytes,1,opt,name=subsystem,proto3" json:"subsystem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GenericSubsystemResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetSubsystemsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetSusbsytemsResponse struct {
	SubsystemsStatus     map[string]bool `protobuf:"bytes,1,rep,name=subsystems_status,json=subsystemsStatus,proto3" json:"subsystems_status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}
type GetRPCEndpointsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type RPCEndpoint struct {
	Started              bool     `protobuf:"varint,1,opt,name=started,proto3" json:"started,omitempty"`
	ListenAddress        string   `protobuf:"bytes,2,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetRPCEndpointsResponse struct {
	Endpoints            map[string]*RPCEndpoint `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}
type GenericExchangeNameRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GenericExchangeNameResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetExchangesRequest struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetExchangesResponse struct {
	Exchanges            string   `protobuf:"bytes,1,opt,name=exchanges,proto3" json:"exchanges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetExchangeOTPReponse struct {
	OtpCode              string   `protobuf:"bytes,1,opt,name=otp_code,json=otpCode,proto3" json:"otp_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetExchangeOTPsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetExchangeOTPsResponse struct {
	OtpCodes             map[string]string `protobuf:"bytes,1,rep,name=otp_codes,json=otpCodes,proto3" json:"otp_codes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
type DisableExchangeRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type PairsSupported struct {
	AvailablePairs       string   `protobuf:"bytes,1,opt,name=available_pairs,json=availablePairs,proto3" json:"available_pairs,omitempty"`
	EnabledPairs         string   `protobuf:"bytes,2,opt,name=enabled_pairs,json=enabledPairs,proto3" json:"enabled_pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetExchangeInfoResponse struct {
	Name                 string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled              bool                       `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Verbose              bool                       `protobuf:"varint,3,opt,name=verbose,proto3" json:"verbose,omitempty"`
	UsingSandbox         bool                       `protobuf:"varint,4,opt,name=using_sandbox,json=usingSandbox,proto3" json:"using_sandbox,omitempty"`
	HttpTimeout          string                     `protobuf:"bytes,5,opt,name=http_timeout,json=httpTimeout,proto3" json:"http_timeout,omitempty"`
	HttpUseragent        string                     `protobuf:"bytes,6,opt,name=http_useragent,json=httpUseragent,proto3" json:"http_useragent,omitempty"`
	HttpProxy            string                     `protobuf:"bytes,7,opt,name=http_proxy,json=httpProxy,proto3" json:"http_proxy,omitempty"`
	BaseCurrencies       string                     `protobuf:"bytes,8,opt,name=base_currencies,json=baseCurrencies,proto3" json:"base_currencies,omitempty"`
	SupportedAssets      map[string]*PairsSupported `protobuf:"bytes,9,rep,name=supported_assets,json=supportedAssets,proto3" json:"supported_assets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AuthenticatedApi     bool                       `protobuf:"varint,10,opt,name=authenticated_api,json=authenticatedApi,proto3" json:"authenticated_api,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}
type GetTickerRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	AssetType            string        `protobuf:"bytes,3,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type CurrencyPair struct {
	Delimiter            string   `protobuf:"bytes,1,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	Base                 string   `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Quote                string   `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type TickerResponse struct {
	Pair                 *CurrencyPair `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	LastUpdated          int64         `protobuf:"varint,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	CurrencyPair         string        `protobuf:"bytes,3,opt,name=currency_pair,json=currencyPair,proto3" json:"currency_pair,omitempty"`
	Last                 float64       `protobuf:"fixed64,4,opt,name=last,proto3" json:"last,omitempty"`
	High                 float64       `protobuf:"fixed64,5,opt,name=high,proto3" json:"high,omitempty"`
	Low                  float64       `protobuf:"fixed64,6,opt,name=low,proto3" json:"low,omitempty"`
	Bid                  float64       `protobuf:"fixed64,7,opt,name=bid,proto3" json:"bid,omitempty"`
	Ask                  float64       `protobuf:"fixed64,8,opt,name=ask,proto3" json:"ask,omitempty"`
	Volume               float64       `protobuf:"fixed64,9,opt,name=volume,proto3" json:"volume,omitempty"`
	PriceAth             float64       `protobuf:"fixed64,10,opt,name=price_ath,json=priceAth,proto3" json:"price_ath,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type GetTickersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type Tickers struct {
	Exchange             string            `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Tickers              []*TickerResponse `protobuf:"bytes,2,rep,name=tickers,proto3" json:"tickers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
type GetTickersResponse struct {
	Tickers              []*Tickers `protobuf:"bytes,1,rep,name=tickers,proto3" json:"tickers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}
type GetOrderbookRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	AssetType            string        `protobuf:"bytes,3,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type OrderbookItem struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type OrderbookResponse struct {
	Pair                 *CurrencyPair    `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	CurrencyPair         string           `protobuf:"bytes,2,opt,name=currency_pair,json=currencyPair,proto3" json:"currency_pair,omitempty"`
	Bids                 []*OrderbookItem `protobuf:"bytes,3,rep,name=bids,proto3" json:"bids,omitempty"`
	Asks                 []*OrderbookItem `protobuf:"bytes,4,rep,name=asks,proto3" json:"asks,omitempty"`
	LastUpdated          int64            `protobuf:"varint,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	AssetType            string           `protobuf:"bytes,6,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}
type GetOrderbooksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type Orderbooks struct {
	Exchange             string               `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Orderbooks           []*OrderbookResponse `protobuf:"bytes,2,rep,name=orderbooks,proto3" json:"orderbooks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}
type GetOrderbooksResponse struct {
	Orderbooks           []*Orderbooks `protobuf:"bytes,1,rep,name=orderbooks,proto3" json:"orderbooks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type GetAccountInfoRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type Account struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Currencies           []*AccountCurrencyInfo `protobuf:"bytes,2,rep,name=currencies,proto3" json:"currencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}
type AccountCurrencyInfo struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	TotalValue           float64  `protobuf:"fixed64,2,opt,name=total_value,json=totalValue,proto3" json:"total_value,omitempty"`
	Hold                 float64  `protobuf:"fixed64,3,opt,name=hold,proto3" json:"hold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetAccountInfoResponse struct {
	Exchange             string     `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Accounts             []*Account `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}
type GetConfigRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetConfigResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type PortfolioAddress struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CoinType             string   `protobuf:"bytes,2,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Balance              float64  `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetPortfolioRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetPortfolioResponse struct {
	Portfolio            []*PortfolioAddress `protobuf:"bytes,1,rep,name=portfolio,proto3" json:"portfolio,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}
type GetPortfolioSummaryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type Coin struct {
	Coin                 string   `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin,omitempty"`
	Balance              float64  `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Percentage           float64  `protobuf:"fixed64,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type OfflineCoinSummary struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance              float64  `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Percentage           float64  `protobuf:"fixed64,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type OnlineCoinSummary struct {
	Balance              float64  `protobuf:"fixed64,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Percentage           float64  `protobuf:"fixed64,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type OfflineCoins struct {
	Addresses            []*OfflineCoinSummary `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}
type OnlineCoins struct {
	Coins                map[string]*OnlineCoinSummary `protobuf:"bytes,1,rep,name=coins,proto3" json:"coins,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}
type GetPortfolioSummaryResponse struct {
	CoinTotals           []*Coin                  `protobuf:"bytes,1,rep,name=coin_totals,json=coinTotals,proto3" json:"coin_totals,omitempty"`
	CoinsOffline         []*Coin                  `protobuf:"bytes,2,rep,name=coins_offline,json=coinsOffline,proto3" json:"coins_offline,omitempty"`
	CoinsOfflineSummary  map[string]*OfflineCoins `protobuf:"bytes,3,rep,name=coins_offline_summary,json=coinsOfflineSummary,proto3" json:"coins_offline_summary,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CoinsOnline          []*Coin                  `protobuf:"bytes,4,rep,name=coins_online,json=coinsOnline,proto3" json:"coins_online,omitempty"`
	CoinsOnlineSummary   map[string]*OnlineCoins  `protobuf:"bytes,5,rep,name=coins_online_summary,json=coinsOnlineSummary,proto3" json:"coins_online_summary,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}
type AddPortfolioAddressRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CoinType             string   `protobuf:"bytes,2,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Balance              float64  `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty"`
	SupportedExchanges   string   `protobuf:"bytes,5,opt,name=supported_exchanges,json=supportedExchanges,proto3" json:"supported_exchanges,omitempty"`
	ColdStorage          bool     `protobuf:"varint,6,opt,name=cold_storage,json=coldStorage,proto3" json:"cold_storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type AddPortfolioAddressResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type RemovePortfolioAddressRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CoinType             string   `protobuf:"bytes,2,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type RemovePortfolioAddressResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetForexProvidersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type ForexProvider struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Verbose              bool     `protobuf:"varint,3,opt,name=verbose,proto3" json:"verbose,omitempty"`
	RestPollingDelay     string   `protobuf:"bytes,4,opt,name=rest_polling_delay,json=restPollingDelay,proto3" json:"rest_polling_delay,omitempty"`
	ApiKey               string   `protobuf:"bytes,5,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiKeyLevel          int64    `protobuf:"varint,6,opt,name=api_key_level,json=apiKeyLevel,proto3" json:"api_key_level,omitempty"`
	PrimaryProvider      bool     `protobuf:"varint,7,opt,name=primary_provider,json=primaryProvider,proto3" json:"primary_provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetForexProvidersResponse struct {
	ForexProviders       []*ForexProvider `protobuf:"bytes,1,rep,name=forex_providers,json=forexProviders,proto3" json:"forex_providers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}
type GetForexRatesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type ForexRatesConversion struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Rate                 float64  `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
	InverseRate          float64  `protobuf:"fixed64,4,opt,name=inverse_rate,json=inverseRate,proto3" json:"inverse_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetForexRatesResponse struct {
	ForexRates           []*ForexRatesConversion `protobuf:"bytes,1,rep,name=forex_rates,json=forexRates,proto3" json:"forex_rates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}
type OrderDetails struct {
	Exchange             string          `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	BaseCurrency         string          `protobuf:"bytes,3,opt,name=base_currency,json=baseCurrency,proto3" json:"base_currency,omitempty"`
	QuoteCurrency        string          `protobuf:"bytes,4,opt,name=quote_currency,json=quoteCurrency,proto3" json:"quote_currency,omitempty"`
	AssetType            string          `protobuf:"bytes,5,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	OrderSide            string          `protobuf:"bytes,6,opt,name=order_side,json=orderSide,proto3" json:"order_side,omitempty"`
	OrderType            string          `protobuf:"bytes,7,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	CreationTime         int64           `protobuf:"varint,8,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	Status               string          `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	Price                float64         `protobuf:"fixed64,10,opt,name=price,proto3" json:"price,omitempty"`
	Amount               float64         `protobuf:"fixed64,11,opt,name=amount,proto3" json:"amount,omitempty"`
	OpenVolume           float64         `protobuf:"fixed64,12,opt,name=open_volume,json=openVolume,proto3" json:"open_volume,omitempty"`
	Fee                  float64         `protobuf:"fixed64,13,opt,name=fee,proto3" json:"fee,omitempty"`
	Trades               []*TradeHistory `protobuf:"bytes,14,rep,name=trades,proto3" json:"trades,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}
type TradeHistory struct {
	CreationTime         int64    `protobuf:"varint,1,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Amount               float64  `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Exchange             string   `protobuf:"bytes,5,opt,name=exchange,proto3" json:"exchange,omitempty"`
	AssetType            string   `protobuf:"bytes,6,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	OrderSide            string   `protobuf:"bytes,7,opt,name=order_side,json=orderSide,proto3" json:"order_side,omitempty"`
	Fee                  float64  `protobuf:"fixed64,8,opt,name=fee,proto3" json:"fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetOrdersRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	AssetType            string        `protobuf:"bytes,2,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type GetOrdersResponse struct {
	Orders               []*OrderDetails `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}
type GetOrderRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type SubmitOrderRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	Side                 string        `protobuf:"bytes,3,opt,name=side,proto3" json:"side,omitempty"`
	OrderType            string        `protobuf:"bytes,4,opt,name=order_type,json=orderType,proto3" json:"order_type,omitempty"`
	Amount               float64       `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Price                float64       `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	ClientId             string        `protobuf:"bytes,7,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type SubmitOrderResponse struct {
	OrderPlaced          bool     `protobuf:"varint,1,opt,name=order_placed,json=orderPlaced,proto3" json:"order_placed,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type SimulateOrderRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	Amount               float64       `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Side                 string        `protobuf:"bytes,4,opt,name=side,proto3" json:"side,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type SimulateOrderResponse struct {
	Orders               []*OrderbookItem `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Amount               float64          `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	MinimumPrice         float64          `protobuf:"fixed64,3,opt,name=minimum_price,json=minimumPrice,proto3" json:"minimum_price,omitempty"`
	MaximumPrice         float64          `protobuf:"fixed64,4,opt,name=maximum_price,json=maximumPrice,proto3" json:"maximum_price,omitempty"`
	PercentageGainLoss   float64          `protobuf:"fixed64,5,opt,name=percentage_gain_loss,json=percentageGainLoss,proto3" json:"percentage_gain_loss,omitempty"`
	Status               string           `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}
type WhaleBombRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	PriceTarget          float64       `protobuf:"fixed64,3,opt,name=price_target,json=priceTarget,proto3" json:"price_target,omitempty"`
	Side                 string        `protobuf:"bytes,4,opt,name=side,proto3" json:"side,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type CancelOrderRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	AccountId            string        `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	OrderId              string        `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,4,opt,name=pair,proto3" json:"pair,omitempty"`
	AssetType            string        `protobuf:"bytes,5,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	WalletAddress        string        `protobuf:"bytes,6,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	Side                 string        `protobuf:"bytes,7,opt,name=side,proto3" json:"side,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type CancelOrderResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type CancelAllOrdersRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type CancelAllOrdersResponse struct {
	Orders               []*CancelAllOrdersResponse_Orders `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}
type GetEventsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type ConditionParams struct {
	Condition            string   `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	CheckBids            bool     `protobuf:"varint,3,opt,name=check_bids,json=checkBids,proto3" json:"check_bids,omitempty"`
	CheckBidsAndAsks     bool     `protobuf:"varint,4,opt,name=check_bids_and_asks,json=checkBidsAndAsks,proto3" json:"check_bids_and_asks,omitempty"`
	OrderbookAmount      float64  `protobuf:"fixed64,5,opt,name=orderbook_amount,json=orderbookAmount,proto3" json:"orderbook_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetEventsResponse struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             string           `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Item                 string           `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	ConditionParams      *ConditionParams `protobuf:"bytes,4,opt,name=condition_params,json=conditionParams,proto3" json:"condition_params,omitempty"`
	Pair                 *CurrencyPair    `protobuf:"bytes,5,opt,name=pair,proto3" json:"pair,omitempty"`
	Action               string           `protobuf:"bytes,6,opt,name=action,proto3" json:"action,omitempty"`
	Executed             bool             `protobuf:"varint,7,opt,name=executed,proto3" json:"executed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}
type AddEventRequest struct {
	Exchange             string           `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Item                 string           `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	ConditionParams      *ConditionParams `protobuf:"bytes,3,opt,name=condition_params,json=conditionParams,proto3" json:"condition_params,omitempty"`
	Pair                 *CurrencyPair    `protobuf:"bytes,4,opt,name=pair,proto3" json:"pair,omitempty"`
	AssetType            string           `protobuf:"bytes,5,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	Action               string           `protobuf:"bytes,6,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}
type AddEventResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type RemoveEventRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type RemoveEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetCryptocurrencyDepositAddressesRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetCryptocurrencyDepositAddressesResponse struct {
	Addresses            map[string]string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}
type GetCryptocurrencyDepositAddressRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Cryptocurrency       string   `protobuf:"bytes,2,opt,name=cryptocurrency,proto3" json:"cryptocurrency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetCryptocurrencyDepositAddressResponse struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type WithdrawFiatRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Currency             string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	BankAccountId        string   `protobuf:"bytes,5,opt,name=bank_account_id,json=bankAccountId,proto3" json:"bank_account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type WithdrawCryptoRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	AddressTag           string   `protobuf:"bytes,3,opt,name=address_tag,json=addressTag,proto3" json:"address_tag,omitempty"`
	Currency             string   `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               float64  `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee                  float64  `protobuf:"fixed64,6,opt,name=fee,proto3" json:"fee,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type WithdrawResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type WithdrawalEventByIDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type WithdrawalEventByIDResponse struct {
	Event                *WithdrawalEventResponse `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}
type WithdrawalEventsByExchangeRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type WithdrawalEventsByDateRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  string   `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type WithdrawalEventsByExchangeResponse struct {
	Event                []*WithdrawalEventResponse `protobuf:"bytes,2,rep,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}
type WithdrawalEventResponse struct {
	Id                   string                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Exchange             *WithdrawlExchangeEvent `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Request              *WithdrawalRequestEvent `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
	CreatedAt            *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp    `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}
type WithdrawlExchangeEvent struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type WithdrawalRequestEvent struct {
	Currency             string                 `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Description          string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount               float64                `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Type                 int32                  `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Fiat                 *FiatWithdrawalEvent   `protobuf:"bytes,6,opt,name=fiat,proto3" json:"fiat,omitempty"`
	Crypto               *CryptoWithdrawalEvent `protobuf:"bytes,7,opt,name=crypto,proto3" json:"crypto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}
type FiatWithdrawalEvent struct {
	BankName             string   `protobuf:"bytes,1,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	AccountName          string   `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	AccountNumber        string   `protobuf:"bytes,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Bsb                  string   `protobuf:"bytes,4,opt,name=bsb,proto3" json:"bsb,omitempty"`
	Swift                string   `protobuf:"bytes,5,opt,name=swift,proto3" json:"swift,omitempty"`
	Iban                 string   `protobuf:"bytes,6,opt,name=iban,proto3" json:"iban,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type CryptoWithdrawalEvent struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AddressTag           string   `protobuf:"bytes,2,opt,name=address_tag,json=addressTag,proto3" json:"address_tag,omitempty"`
	Fee                  float64  `protobuf:"fixed64,3,opt,name=fee,proto3" json:"fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetLoggerDetailsRequest struct {
	Logger               string   `protobuf:"bytes,1,opt,name=logger,proto3" json:"logger,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetLoggerDetailsResponse struct {
	Info                 bool     `protobuf:"varint,1,opt,name=info,proto3" json:"info,omitempty"`
	Debug                bool     `protobuf:"varint,2,opt,name=debug,proto3" json:"debug,omitempty"`
	Warn                 bool     `protobuf:"varint,3,opt,name=warn,proto3" json:"warn,omitempty"`
	Error                bool     `protobuf:"varint,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type SetLoggerDetailsRequest struct {
	Logger               string   `protobuf:"bytes,1,opt,name=logger,proto3" json:"logger,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetExchangePairsRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Asset                string   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetExchangePairsResponse struct {
	SupportedAssets      map[string]*PairsSupported `protobuf:"bytes,1,rep,name=supported_assets,json=supportedAssets,proto3" json:"supported_assets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}
type ExchangePairRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	AssetType            string        `protobuf:"bytes,2,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,3,opt,name=pair,proto3" json:"pair,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type GetOrderbookStreamRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	AssetType            string        `protobuf:"bytes,3,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type GetExchangeOrderbookStreamRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetTickerStreamRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	AssetType            string        `protobuf:"bytes,3,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type GetExchangeTickerStreamRequest struct {
	Exchange             string   `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetAuditEventRequest struct {
	StartDate            string   `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              string   `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	OrderBy              string   `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GetAuditEventResponse struct {
	Events               []*AuditEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type GetHistoricCandlesRequest struct {
	Exchange             string        `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Pair                 *CurrencyPair `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	AssetType            string        `protobuf:"bytes,3,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	Start                int64         `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64         `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	TimeInterval         int64         `protobuf:"varint,6,opt,name=time_interval,json=timeInterval,proto3" json:"time_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}
type GetHistoricCandlesResponse struct {
	Candle               []*Candle `protobuf:"bytes,1,rep,name=candle,proto3" json:"candle,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}
type Candle struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Low                  float64  `protobuf:"fixed64,2,opt,name=low,proto3" json:"low,omitempty"`
	High                 float64  `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	Open                 float64  `protobuf:"fixed64,4,opt,name=open,proto3" json:"open,omitempty"`
	Close                float64  `protobuf:"fixed64,5,opt,name=close,proto3" json:"close,omitempty"`
	Volume               float64  `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type AuditEvent struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GCTScript struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	NextRun              string   `protobuf:"bytes,4,opt,name=next_run,json=nextRun,proto3" json:"next_run,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GCTScriptExecuteRequest struct {
	Script               *GCTScript `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}
type GCTScriptStopRequest struct {
	Script               *GCTScript `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}
type GCTScriptStopAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GCTScriptStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GCTScriptListAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GCTScriptUploadRequest struct {
	ScriptName           string   `protobuf:"bytes,1,opt,name=script_name,json=scriptName,proto3" json:"script_name,omitempty"`
	ScriptData           string   `protobuf:"bytes,2,opt,name=script_data,json=scriptData,proto3" json:"script_data,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Archived             bool     `protobuf:"varint,4,opt,name=archived,proto3" json:"archived,omitempty"`
	Overwrite            bool     `protobuf:"varint,5,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GCTScriptReadScriptRequest struct {
	Script               *GCTScript `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}
type GCTScriptQueryRequest struct {
	Script               *GCTScript `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}
type GCTScriptAutoLoadRequest struct {
	Script               string   `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	Status               bool     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type GCTScriptStatusResponse struct {
	Status               string       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Scripts              []*GCTScript `protobuf:"bytes,2,rep,name=scripts,proto3" json:"scripts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}
type GCTScriptQueryResponse struct {
	Status               string     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Script               *GCTScript `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Data                 string     `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}
type GCTScriptGenericResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type goCryptoTraderClient struct {
	cc grpc.ClientConnInterface
}
type goCryptoTraderGetAccountInfoStreamClient struct {
	grpc.ClientStream
}
type goCryptoTraderGetOrderbookStreamClient struct {
	grpc.ClientStream
}
type goCryptoTraderGetExchangeOrderbookStreamClient struct {
	grpc.ClientStream
}
type goCryptoTraderGetTickerStreamClient struct {
	grpc.ClientStream
}
type goCryptoTraderGetExchangeTickerStreamClient struct {
	grpc.ClientStream
}
type UnimplementedGoCryptoTraderServer struct {
}
type goCryptoTraderGetAccountInfoStreamServer struct {
	grpc.ServerStream
}
type goCryptoTraderGetOrderbookStreamServer struct {
	grpc.ServerStream
}
type goCryptoTraderGetExchangeOrderbookStreamServer struct {
	grpc.ServerStream
}
type goCryptoTraderGetTickerStreamServer struct {
	grpc.ServerStream
}
type goCryptoTraderGetExchangeTickerStreamServer struct {
	grpc.ServerStream
}
