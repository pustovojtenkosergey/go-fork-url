package tests

import (
	"testing"

	"myserv/app"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	config := &app.Config{
		Port:     "8080",
		MongoUri: "mongodb://localhost:27017",
	}
	a := app.NewApp(config)
	assert.NotNil(t, a)
	assert.Equal(t, "8080", a.Port)
}
