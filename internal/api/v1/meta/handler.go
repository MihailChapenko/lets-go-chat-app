package meta

import "github.com/gorilla/mux"

//Handler describe meta handler interface
type Handler interface {
	InitRoutes(r *mux.Router)
}

type handler struct {}

//NewHandler returns handler instance
func NewHandler() Handler {
	return &handler{}
}

//InitRoutes initialize meta routes
func (h handler) InitRoutes(r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter().PathPrefix("/meta").Subrouter()

	v1.HandleFunc("/", h.PrintMessage).Methods("GET").Schemes("http")
	v1.HandleFunc("/else", h.Else).Methods("GET").Schemes("http")
}