package _type

import (
	"ReconDB/middlewares/address"
	"ReconDB/models"
	"ReconDB/pkg/check"
	"ReconDB/pkg/parser"
	"fmt"
)

func FindAssetType(Asset models.Assets) (string, error) {

	if address.ValidateDomainName(Asset.Asset) {
		return "single", nil
	}

	if check.CheckIPAddress(Asset.Asset) {
		return "ip", nil
	}

	if _, _, err := parser.ParseCidr(Asset.Asset); err == nil {
		return "cidr", fmt.Errorf("cidr is not allowed for asset")
	}

	if check.WildCardRegex(Asset.Asset) {
		return "wildcard", fmt.Errorf("wildcard are not allowed for asset")
	}

	return "", fmt.Errorf("asset type not found")

}
