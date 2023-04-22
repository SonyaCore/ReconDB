package models

type Assets struct {
	Scope     string `json:"scope"`
	Asset     string `json:"asset"`
	AssetType string `json:"-"`
	IsLive    bool   `json:"is_live"`
}
