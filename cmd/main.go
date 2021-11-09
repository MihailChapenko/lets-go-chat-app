package main

import (
	"github.com/MihailChapenko/lets-go-chat-app/cmd/server"
	"github.com/MihailChapenko/lets-go-chat-app/internal/config"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	config.Init(dir + "/../internal/config/config.yaml")
	//db.Init(config.Get().DB)
	server.Init()
}
