package actions

import (
	"context"
	"encoding/json"
	"myserv/db"
	"myserv/db/repo"
	"net/http"
)

type UrlStatsAction struct {
	UrlRepository *repo.UrlRepository
}

func NewUrlStatsAction(dbClient *db.DbClient) *UrlStatsAction {
	return &UrlStatsAction{
		UrlRepository: dbClient.UrlRepository,
	}
}

func (a *UrlStatsAction) Handle(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	ctx := context.Background()
	urls, err := a.UrlRepository.FindByFilter(ctx, repo.NewFilter())

	if err != nil {
		http.Error(w, "Failed to fetch urls", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(urls); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
