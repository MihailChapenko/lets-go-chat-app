package user

import (
	"encoding/json"
	"github.com/MihailChapenko/lets-go-chat-app/internal/models"
	"log"
	"net/http"
)

var users = make(map[string]string)

//Register new user registration
func (h handler) Register(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("error"))
	}
	store(user)

	payload, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

//Login login user
func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Write([]byte("error"))
	}

	if _, ok := users[user.UserName]; !ok {
		w.Write([]byte("no such user"))
	}

	payload, err := json.Marshal(users[user.UserName])
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func store(user models.User) {
	users[user.UserName] = user.Password
}
