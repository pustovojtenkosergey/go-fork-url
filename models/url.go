package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Url struct {
	ModelAgg `bson:",inline"`
	Name     string `bson:"name"`
	Counter  int    `bson:"counter"`
}

func NewUrl(name string) *Url {
	url := &Url{
		Name:    name,
		Counter: 1,
	}
	url.InitDate()
	return url
}

func (url *Url) GetName() string {
	return url.Name
}

func (url *Url) Increment() {
	url.Counter++
	url.UpdateDate()
}

func (url *Url) UpdateBSON() bson.M {
	return bson.M{
		"counter":    url.Counter,
		"updated_at": url.UpdatedAt,
	}
}
