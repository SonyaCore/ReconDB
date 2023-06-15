package typeassert

import (
	"ReconDB/models"
	"ReconDB/pkg/host"
	"fmt"
)

// FindAssetType matches the typeassert of asset with a valid typeassert
// if typeassert of the asset not found it returs asset typeassert not found error.
func FindAssetType(Asset models.Assets) (string, error) {

	// domain
	if err := host.CheckDomain(Asset.Asset); err == nil {
		return "single", nil
	}

	// ip
	if err := host.IpAddress(Asset.Asset); err == nil {
		return "ip", nil
	}

	// cidr
	if _, _, err := host.ParseCidr(Asset.Asset); err == nil {
		return "cidr", fmt.Errorf("cidr is not allowed for asset")
	}

	// wildcard
	if host.MatchWildcard(Asset.Scope, Asset.Asset) {
		return "wildcard", fmt.Errorf("wildcard are not allowed for asset")
	}

	// err
	return "", fmt.Errorf("asset type not found")

}
