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
)

type UrlIncrementAction struct {
	dbClient *db.DbClient
}

func NewUrlIncrementAction(dbClient *db.DbClient) *UrlIncrementAction {
	return &UrlIncrementAction{dbClient: dbClient}
}

func (a *UrlIncrementAction) Handle(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	w.Header().Set("Content-Type", "application/json")

	var counter int
	url := models.NewUrl(r.URL.Path)

	ctx := context.Background()
	filter := repo.NewFilter()
	filter.AddValue("name", url.Name)

	urls, err := a.dbClient.UrlRepository.FindByFilter(ctx, filter)
	if err != nil || len(urls) == 0 {
		err = a.dbClient.UrlRepository.Insert(ctx, url)
		if err != nil {
			log.Printf("Failed to insert url: %v", err)
			counter = 0
		}
		counter = 1
	} else {
		url := urls[0]
		url.Increment()
		err = a.dbClient.UrlRepository.Update(ctx, &url)
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
	json.NewEncoder(w).Encode(map[string]string{"counter": fmt.Sprint(counter)})
}
