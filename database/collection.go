package database

import "go.mongodb.org/mongo-driver/mongo"

func Collection(collection string) *mongo.Collection {
	client := Client()
	database := client.Database("ReconDB")

	Collection := database.Collection(collection)
	return Collection
}
