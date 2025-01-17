package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"myserv/db"
	"myserv/db/repo"
	"myserv/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlIncrementAction struct {
	UrlRepository *repo.UrlRepository
}

func NewUrlIncrementAction(dbClient *db.DbClient) *UrlIncrementAction {
	return &UrlIncrementAction{
		UrlRepository: dbClient.UrlRepository,
	}
}

func (a *UrlIncrementAction) Handle(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	w.Header().Set("Content-Type", "application/json")

	var counter int
	url := models.NewUrl(r.URL.Path)

	ctx := context.Background()
	filter := repo.NewFilter()
	filter.AddValue("name", url.Name)

	var id primitive.ObjectID

	urls, err := a.UrlRepository.FindByFilter(ctx, filter)
	if err != nil || len(urls) == 0 {
		id, err = a.UrlRepository.Insert(ctx, url)
		if err != nil {
			log.Printf("Failed to insert url: %v", err)
			counter = 0
		}
		counter = 1
	} else {
		url := urls[0]
		id = url.GetID()
		url.Increment()
		err = a.UrlRepository.Update(ctx, &url)
		if err != nil {
			log.Printf("Failed to update url: %v", err)
			counter = 0
		}
		counter = url.Counter
	}

	if counter == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to increment url"})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"id":      id.Hex(),
		"counter": fmt.Sprint(counter),
	})
}
