package database

import (
	"ReconDB/config"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection returns an instance of mongo.Collection
// it's being used for importing documents in collections
func Collection(collection string) *mongo.Collection {
	config, _ := config.LoadConfig()
	client := Client(config)
	// set the database name
	database := client.Database(config.DataBaseName)
	Collection := database.Collection(collection)
	return Collection
}
