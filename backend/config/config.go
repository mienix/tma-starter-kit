package config

import (
	"os"
)

type HTTPServerConfig struct {
	Network  string `env:"NETWORK,required"`
	Addr     string `env:"ADDR,required"`
	BotToken string `env:"TG_BOT_TOKEN,required"`
}

func (httpConf *HTTPServerConfig) Load() {
	httpConf.Network = os.Getenv("NETWORK")
	httpConf.Addr = os.Getenv("ADDR")
	httpConf.BotToken = os.Getenv("TG_BOT_TOKEN")
}

type MongoUserDBConfig struct {
	URI            string `env:"MONGODB_URI,required"`
	DBName         string `env:"MONGODB_NAME,required"`
	CollectionName string `env:"MONGODB_COLLECTION_NAME,required"`
}

func (mongoConf *MongoUserDBConfig) Load() {
	mongoConf.URI = os.Getenv("MONGODB_URI")
	mongoConf.DBName = os.Getenv("MONGODB_NAME")
	mongoConf.CollectionName = os.Getenv("MONGODB_COLLECTION_NAME")
}
