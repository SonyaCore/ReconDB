package config

// Config types
// PORT is used to set the gin webserver port
// MongoURI is uri for mongodb database
type Config struct {
	PORT     string `json:"port"`
	MongoURI string `json:"mongo_uri"`
}
