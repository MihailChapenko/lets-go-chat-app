package user

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	var register = struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		w.Write([]byte("error"))
	}

	payload, err := json.Marshal(register)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Else!\n"))
}
