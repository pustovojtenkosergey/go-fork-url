package actions

import (
	"net/http"
)

type Action interface {
	Handle(w http.ResponseWriter, r *http.Request, vars map[string]string)
}
