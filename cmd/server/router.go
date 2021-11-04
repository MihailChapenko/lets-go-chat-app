package server

import (
	"github.com/MihailChapenko/lets-go-chat-app/internal/api/v1/user"
	"github.com/gorilla/mux"
	"net/http"
)

//NewRouter initialize router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1").Subrouter()
	user.NewHandler().InitRoutes(v1)

	http.Handle("/", r)

	return r
}
