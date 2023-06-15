package asset

import (
	"ReconDB/database"
	"ReconDB/models"
	"ReconDB/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindCompanyName(Asset models.Assets) (string, error) {
	// Find the matching scope for the given Asset
	var scopeCollection = []string{"Scopes", "OutofScopes"}
	var err error

	var scopeResult struct {
		CompanyName string `bson:"companyname"`
	}

	Asset.Scope = utils.WildCardToRegex(Asset.Scope)

	for _, scope := range scopeCollection {
		collectionScope := database.Collection(scope)

		// find company name based on scope
		scopeQuery := bson.M{
			//"scopetype": Asset.AssetType,
			"scope": primitive.Regex{
				Pattern: "^" + Asset.Scope + "$",
				Options: "i",
			},
		}

		opts := options.FindOne().SetProjection(bson.M{"companyname": 1})
		collectionScope.FindOne(context.Background(), scopeQuery, opts).Decode(&scopeResult)

		if len(scopeResult.CompanyName) < 1 {
			continue
		}

		return scopeResult.CompanyName, nil
	}
	return "", err
}
