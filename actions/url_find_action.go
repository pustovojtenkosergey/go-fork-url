package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"myserv/db"
	"myserv/db/repo"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlFindAction struct {
	dbClient *db.DbClient
}

func NewUrlFindAction(dbClient *db.DbClient) *UrlFindAction {
	return &UrlFindAction{dbClient: dbClient}
}

func (a *UrlFindAction) Handle(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	var ID primitive.ObjectID
	var err error

	if id, ok := vars["id"]; ok {
		ID, err = primitive.ObjectIDFromHex(id)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
			return
		}
	} else {
		json.NewEncoder(w).Encode(map[string]string{"error": "No ID found"})
		return
	}

	filter := repo.NewFilter()
	filter.AddValue("_id", ID)

	urls, err := a.dbClient.UrlRepository.FindByFilter(context.Background(), filter)

	if err != nil || len(urls) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"error": "Url not found"})
		return
	}

	url := urls[0]
	json.NewEncoder(w).Encode(map[string]string{"name": url.Name, "counter": fmt.Sprint(url.Counter)})
}
