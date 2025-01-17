package repo

import (
	"context"
	"myserv/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "urls"
)

type UrlRepository struct {
	Repository
}

func NewUrlRepository(db *mongo.Database) *UrlRepository {
	repo := &UrlRepository{}
	repo.collection = db.Collection(collectionName)
	return repo
}

func (r *UrlRepository) FindByFilter(ctx context.Context, filter *Filter) ([]models.Url, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, filter.Filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var urls []models.Url
	for cursor.Next(ctx) {
		var url models.Url
		if err := cursor.Decode(&url); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}
