package config

// Config types
// PORT is used to set the gin webserver port
// MongoURI is uri for mongodb database
type Config struct {
	Mode          string `json:"gin_mode"`
	Authorization string `json:"authorization"`
	PORT          string `json:"port"`
	MongoURI      string `json:"mongo_uri"`
	DataBaseName  string `json:"database"`
}

type RouterURI struct {
	API        string `json:"api_endpoint"`
	Scope      string `json:"scope"`
	OutofScope string `json:"out_of_scope"`
	Company    string `json:"company"`
	Asset      string `json:"asset"`
}
