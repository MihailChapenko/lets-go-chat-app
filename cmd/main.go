package main

import (
	"github.com/MihailChapenko/lets-go-chat-app/cmd/server"
	"github.com/MihailChapenko/lets-go-chat-app/internal/config"
)

func main() {
	config.Init("../internal/config/config.yaml")
	server.Init()
}
