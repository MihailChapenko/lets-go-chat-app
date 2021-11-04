package meta

import (
	"net/http"
)

func (h handler) PrintMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func (h handler) Else(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Else!\n"))
}
