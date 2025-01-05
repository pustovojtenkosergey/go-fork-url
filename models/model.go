package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model interface {
	GetID() primitive.ObjectID
	UpdateBSON() bson.M
}

type ModelAgg struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *ModelAgg) GetID() primitive.ObjectID {
	return m.ID
}

func (m *ModelAgg) InitDate() {
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now
}

func (m *ModelAgg) UpdateDate() {
	m.UpdatedAt = time.Now()
}
