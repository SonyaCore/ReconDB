package asset

import (
	"ReconDB/database"
	"ReconDB/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindCompanyName(Asset models.Assets) (string, error) {
	// Find the matching scope for the given Asset
	var scopeCollection = "Scopes"
	var err error

	collectionScope := database.Collection(scopeCollection)

	// find company name based on scope
	scopeQuery := bson.M{
		//"scopetype": Asset.AssetType,
		"scope": Asset.Scope,
	}

	opts := options.FindOne().SetProjection(bson.M{"companyname": 1})

	var scopeResult struct {
		CompanyName string `bson:"companyname"`
	}

	if err = collectionScope.FindOne(context.Background(), scopeQuery, opts).Decode(&scopeResult); err != nil {
		return "", err
	}

	return scopeResult.CompanyName, nil
}
