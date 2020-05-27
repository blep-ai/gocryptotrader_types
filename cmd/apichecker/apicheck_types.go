package main

type ShaResponse struct {
	ShaResp string `json:"sha"`
}
type ExchangeInfo struct {
	Name      string
	CheckType string
	Data      *CheckData `json:",omitempty"`
	Disabled  bool
}
type CheckData struct {
	HTMLData   *HTMLScrapingData `json:",omitempty"`
	GitHubData *GithubData       `json:",omitempty"`
}
type HTMLScrapingData struct {
	TokenData     string `json:",omitempty"`
	Key           string `json:",omitempty"`
	Val           string `json:",omitempty"`
	TokenDataEnd  string `json:",omitempty"`
	TextTokenData string `json:",omitempty"`
	DateFormat    string `json:",omitempty"`
	RegExp        string `json:",omitempty"`
	CheckString   string `json:",omitempty"`
	Path          string `json:",omitempty"`
}
type GithubData struct {
	Repo string `json:",omitempty"`
	Sha  string `json:",omitempty"`
}
type ListData struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IDBoard string `json:"idBoard"`
}
type CardFill struct {
	Name      string
	Desc      string
	Pos       string
	Due       string
	ListID    string
	MembersID string
	LabelsID  string
	URLSource string
}
type ItemData struct {
	State    string `json:"state"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position int64  `json:"pos"`
}
type ChecklistItemData struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	CheckItems []ItemData `json:"checkItems"`
}
type MembersData struct {
	Name    string `json:"name"`
	ShortID string `json:"shortlink"`
	ID      string `json:"id"`
}
type Config struct {
	CardID              string `json:"CardID"`
	ChecklistID         string `json:"ChecklistID"`
	ListID              string `json:"ListID"`
	BoardID             string `json:"BoardID"`
	Key                 string `json:"Key"`
	Token               string `json:"Token"`
	CreateCardName      string
	CreateListName      string
	CreateChecklistName string
	Exchanges           []ExchangeInfo `json:"Exchanges"`
}
type TrelloData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
