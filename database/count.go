package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func CountDocuments(collection string, query bson.M) (int64, error) {
	var count int64
	var err error

	if count, err = Collection(collection).CountDocuments(context.Background(), query); err != nil {
		return 0, err
	}
	return count, nil
}
