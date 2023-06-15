package database

import (
	"ReconDB/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.TODO()

// Client
// first it get the uri with mongo_uri value in configuration file
// then it connect to database with mongo.Connect
func Client(config config.Config) *mongo.Client {
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// Ping pinging mongodb connection
func Ping(c *mongo.Client) {
	err := c.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

}
