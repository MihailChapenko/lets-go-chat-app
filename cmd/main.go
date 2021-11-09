package main

import (
	"github.com/MihailChapenko/lets-go-chat-app/cmd/server"
	"github.com/MihailChapenko/lets-go-chat-app/internal/config"
	"os"
	"path/filepath"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config.Init(filepath.Join(pwd, "../internal/config/config.yaml"))
	//db.Init(config.Get().DB)
	server.Init()
}
