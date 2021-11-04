package server

import (
	"github.com/MihailChapenko/lets-go-chat-app/internal/api/v1/meta"
	"github.com/gorilla/mux"
	"net/http"
)

//NewRouter initialize router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	metaHandler := meta.NewHandler()
	metaHandler.InitRoutes(r)

	http.Handle("/", r)

	return r
}
