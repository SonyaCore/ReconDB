package models

type Assets struct {
	Asset     string `json:"asset"`
	AssetType string `json:"-"`
	IsLive    bool   `json:"is_live"`
}
