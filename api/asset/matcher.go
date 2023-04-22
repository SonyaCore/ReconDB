package asset

import (
	"ReconDB/middlewares/address"
	"ReconDB/models"
	"fmt"
)

func FindAssetType(Asset models.Assets) (string, error) {

	if address.ValidateDomainName(Asset.Asset) {
		return "single", nil
	}
	if address.WildCardRegex(Asset.Asset) {
		return "wildcard", fmt.Errorf("wildcard are not allowed for asset")
	}
	if err := address.CheckIPAddress(Asset.Asset); err == nil {
		return "ip", nil
	}

	if _, _, err := address.ParseCidr(Asset.Asset); err == nil {
		return "cidr", fmt.Errorf("cidr is not allowed for asset")
	}

	return "", fmt.Errorf("asset type not found")

}
