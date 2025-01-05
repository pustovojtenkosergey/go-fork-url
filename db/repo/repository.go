package repo

import (
	"context"
	"fmt"
	"myserv/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func (r *Repository) Insert(ctx context.Context, model models.Model) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		return fmt.Errorf("failed to insert model: %w", err)
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, model models.Model) error {
	filter := bson.M{"_id": model.GetID()}
	update := bson.M{"$set": model.UpdateBSON()}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update model: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with ID %s", model.GetID())
	}

	return nil
}

func (r *UrlRepository) Delete(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return r.collection.DeleteOne(ctx, bson.M{"_id": id})
}
