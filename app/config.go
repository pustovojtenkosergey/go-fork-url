package app

import "os"

type Config struct {
	MongoUri string
	Port    string
}

func NewConfig() *Config {
	return &Config{
		MongoUri: os.Getenv("mongodb_uri"),
		Port:    os.Getenv("port"),
	}
}