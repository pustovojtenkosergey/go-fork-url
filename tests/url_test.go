package tests

import (
	"myserv/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateBSON(t *testing.T) {
	url := models.NewUrl("test-url")
	url.Counter = 5
	url.UpdatedAt = time.Now()

	expectedBSON := bson.M{
		"counter":    5,
		"updated_at": url.UpdatedAt,
	}

	result := url.UpdateBSON()
	assert.Equal(t, expectedBSON, result)
}
