package config

type Config struct {
	MongoURI    string
	MongoDBName string
}

var MongoURI = "mongodb://localhost:27017"
var MongoDBName = "mercurio-web-crawler"

func GetConfig() Config {
	return Config{MongoURI: MongoURI, MongoDBName: MongoDBName}
}
