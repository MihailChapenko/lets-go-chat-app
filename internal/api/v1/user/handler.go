package user

import "github.com/gorilla/mux"

//Handler describe meta handler interface
type Handler interface {
	InitRoutes(r *mux.Router)
}

type handler struct{}

//NewHandler returns handler instance
func NewHandler() Handler {
	return &handler{}
}

//InitRoutes initialize user routes
func (h handler) InitRoutes(r *mux.Router) {
	v1 := r.PathPrefix("/user").Subrouter()

	v1.HandleFunc("", h.Register).Methods("POST")
	v1.HandleFunc("/login", h.Login).Methods("POST")
}
