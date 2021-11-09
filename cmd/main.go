package main

import (
	"github.com/MihailChapenko/lets-go-chat-app/cmd/server"
	"github.com/MihailChapenko/lets-go-chat-app/internal/config"
)

func main() {
	config.Init("config.yaml")
	//db.Init(config.Get().DB)
	server.Init()
}
