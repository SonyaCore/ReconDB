package _type

import (
	"ReconDB/middlewares/address"
	"ReconDB/models"
	"ReconDB/pkg/check"
	"ReconDB/pkg/parser"
	"fmt"
)

// FindAssetType matches the type of asset with a valid type
// if type of the asset not found it returs asset type not found error.
func FindAssetType(Asset models.Assets) (string, error) {

	// domain
	if address.ValidateDomainName(Asset.Asset) {
		return "single", nil
	}

	// ip
	if check.IpAddress(Asset.Asset) {
		return "ip", nil
	}

	// cidr
	if _, _, err := parser.ParseCidr(Asset.Asset); err == nil {
		return "cidr", fmt.Errorf("cidr is not allowed for asset")
	}

	// wildcard
	if check.WildCardRegex(Asset.Asset) {
		return "wildcard", fmt.Errorf("wildcard are not allowed for asset")
	}

	// err
	return "", fmt.Errorf("asset type not found")

}
