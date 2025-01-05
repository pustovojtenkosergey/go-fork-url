package db

import (
	"myserv/db/repo"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database string = "myserv_db"
)

type DbClient struct {
	Client        *mongo.Client
	Database      *mongo.Database
	UrlRepository *repo.UrlRepository
}

func NewDbClient(mongoUri string) *DbClient {
	client := GetClient(mongoUri)
	database := client.Database(database)
	return &DbClient{
		Client:        client,
		Database:      database,
		UrlRepository: repo.NewUrlRepository(database),
	}
}
