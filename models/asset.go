package models

// Assets struct model for asset api
type Assets struct {
	// Scope is scope of the asset
	Scope string `json:"scope"`

	// Asset is current asset of the scope
	Asset string `json:"asset"`

	// AssetType dynamically allocates & there is no need for passing type to it
	CompanyName string `json:"-"`

	AssetType string `json:"-"`

	// IsLive indicates that current asset is live or not
	IsLive bool `json:"is_live"`
}
