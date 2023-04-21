package asset

import (
	"ReconDB/middlewares"
	"ReconDB/models"
)

func FindAssetType(Asset models.Assets) string {

	if middlewares.ValidateDomainName(Asset.Asset) {
		return "single"
	}
	if middlewares.WildCardRegex(Asset.Asset) {
		return "wildcard"
	}
	if err := middlewares.CheckIPAddress(Asset.Asset); err == nil {
		return "ip"
	}

	if _, _, err := middlewares.ParseCidr(Asset.Asset); err == nil {
		return "cidr"
	}

	return ""

}
