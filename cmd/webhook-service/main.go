package main

import (
	"log"

	"github.com/grippenet/webhook-service/pkg/config"
	"github.com/grippenet/webhook-service/pkg/server"
)

// Config is the structure that holds all global configuration data

func main() {
	conf := config.LoadConfig()

	httpServer := server.NewHttpServer(conf.GetHttpConfig())
	err := httpServer.Start()
	if err != nil {
		log.Fatalf("Http Server stopped %s", err)
	}
}
