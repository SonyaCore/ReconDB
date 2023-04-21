package config

type Config struct {
	PORT     string `json:"port"`
	MongoURI string `json:"mongo_uri"`
}
