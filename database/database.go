package database

import (
	"context"
	"github.com/spf13/viper"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.TODO()

func Client() *mongo.Client {
	var mongoUri = viper.GetString("mongo_uri")

	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client

}

func Ping(c *mongo.Client) {
	err := c.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

}
