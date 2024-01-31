package main

import (
	"context"
	"log"

	"github.com/grippenet/webhook-service/pkg/config"
	"github.com/grippenet/webhook-service/pkg/server"
	"github.com/grippenet/webhook-service/pkg/types"
	"github.com/grippenet/webhook-service/pkg/webhook"
)

// Config is the structure that holds all global configuration data

func main() {
	conf := config.LoadConfig()

	hookChannel := make(chan types.WebHookInput, conf.HookChannelSize)

	ctx := context.Background()

	handler := webhook.NewWebHookHandler(conf, hookChannel)
	handler.Start(ctx)

	httpServer := server.NewHttpServer(conf.GetHttpConfig(), hookChannel)
	err := httpServer.Start()
	if err != nil {
		log.Fatalf("Http Server stopped %s", err)
	}
}
