package repo

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Filter struct {
	Filter bson.M
}

func NewFilter() *Filter {
	return &Filter{Filter: make(map[string]interface{})}
}

func (f *Filter) AddValue(key string, value interface{}) {
	f.Filter[key] = value
}
