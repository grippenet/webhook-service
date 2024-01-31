package webhook

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/martinlindhe/base36"

	"github.com/grippenet/webhook-service/pkg/config"
	"github.com/grippenet/webhook-service/pkg/types"
)

type WebHookHandler struct {
	hookChannel <-chan types.WebHookInput
	conf        *config.AppConfig
}

func NewWebHookHandler(conf *config.AppConfig, hookChannel <-chan types.WebHookInput) *WebHookHandler {
	return &WebHookHandler{
		hookChannel: hookChannel,
		conf:        conf,
	}
}

func (wh *WebHookHandler) Start(ctx context.Context) {
	go wh.Handle(ctx)
}

func (wh *WebHookHandler) Handle(ctx context.Context) {
	log.Println("Starting Handler...")
	for {
		select {
		case <-ctx.Done():
			log.Println("Handler: Done received")

		case result := <-wh.hookChannel:
			r := rand.Uint64()
			t := uint64(result.Time().UnixMicro())

			fn := strings.ToLower(fmt.Sprintf("%s-%s.json", base36.Encode(t), base36.Encode(r)))

			err := os.WriteFile(fmt.Sprintf("%s/%s", wh.conf.Path, fn), result.Body(), 0644)
			if err != nil {
				log.Printf("Unable to write file %s : %s", fn, err)
			}
		}
	}
}
