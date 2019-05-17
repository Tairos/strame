package command

import "net/http"

type Command interface {
	ID() string
	Handle(w http.ResponseWriter, r *http.Request)
}
