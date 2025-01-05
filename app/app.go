package app

import (
	"log"
	"myserv/actions"
	"myserv/db"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Port     string
	DbClient *db.DbClient
}

func NewApp(config *Config) *App {
	dbClient := db.NewDbClient(config.MongoUri)
	return &App{
		Port:     config.Port,
		DbClient: dbClient,
	}
}

func (a *App) routeMap() map[string]actions.Action {
	return map[string]actions.Action{
		"/stats":     actions.NewUrlStatsAction(a.DbClient),
		"/find/{id}": actions.NewUrlFindAction(a.DbClient),
	}
}

func (a *App) Start() {
	r := mux.NewRouter()

	for path, handler := range a.routeMap() {
		r.HandleFunc(path, adaptHandler(handler))
	}

	r.PathPrefix("/").HandlerFunc(adaptHandler(actions.NewUrlIncrementAction(a.DbClient)))

	http.Handle("/favicon.ico", http.FileServer(http.Dir("./static")))
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+a.Port, nil))
}

func adaptHandler(a actions.Action) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		a.Handle(w, r, vars)
	}
}
