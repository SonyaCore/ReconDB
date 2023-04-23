package models

// Assets struct model for asset api
// Scope is scope of the asset
// Asset is current asset of the scope
// AssetType dynamically allocates & there is no need for passing type to it
// IsLive indicates that current asset is live or not
type Assets struct {
	Scope       string `json:"scope"`
	Asset       string `json:"asset"`
	CompanyName string `json:"-"`
	AssetType   string `json:"-"`
	IsLive      bool   `json:"is_live"`
}
