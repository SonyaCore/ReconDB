package database

import "go.mongodb.org/mongo-driver/mongo"

// Collection returns an instance of mongo.Collection
// it's being used for importing documents in collections
func Collection(collection string) *mongo.Collection {
	client := Client()
	// set the database name
	database := client.Database("ReconDB")

	Collection := database.Collection(collection)
	return Collection
}
