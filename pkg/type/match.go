package _type

import (
	"ReconDB/models"
	"ReconDB/pkg/check"
	"ReconDB/pkg/domain"
	"ReconDB/pkg/parser"
	"ReconDB/pkg/wildcard"
	"fmt"
)

// FindAssetType matches the type of asset with a valid type
// if type of the asset not found it returs asset type not found error.
func FindAssetType(Asset models.Assets) (string, error) {

	// domain
	if err := domain.CheckDomain(Asset.Asset); err == nil {
		return "single", nil
	}

	// ip
	if err := check.IpAddress(Asset.Asset); err == nil {
		return "ip", nil
	}

	// cidr
	if _, _, err := parser.ParseCidr(Asset.Asset); err == nil {
		return "cidr", fmt.Errorf("cidr is not allowed for asset")
	}

	// wildcard
	if wildcard.Match(Asset.Scope, Asset.Asset) {
		return "wildcard", fmt.Errorf("wildcard are not allowed for asset")
	}

	// err
	return "", fmt.Errorf("asset type not found")

}
