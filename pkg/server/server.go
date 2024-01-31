package server

import (
	"log"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/grippenet/webhook-service/pkg/types"
)

type HttpServer struct {
	app         *fiber.App
	config      *types.HttpServerConfig
	instance    string
	start       time.Time
	hookChannel chan<- types.WebHookInput
}

func NewHttpServer(config *types.HttpServerConfig, hookChannel chan<- types.WebHookInput) *HttpServer {
	return &HttpServer{config: config, hookChannel: hookChannel}
}

func (server *HttpServer) WebhookHandler(c *fiber.Ctx) error {
	hookType := c.Params("type")
	if hookType != "sarbacane" {
		return fiber.NewError(fiber.StatusNotFound, "Hook type not handled")
	}
	hookName := c.Params("name")
	found := false
	if len(hookName) > 0 && len(hookName) <= types.WebHookTokenMaxSize {
		_, ok := server.config.Hooks[hookName]
		if ok {
			found = true
		}
	}
	if !found {
		return fiber.NewError(fiber.StatusNotFound, "Bad hook")
	}
	log.Println("Hook received")
	w := &WebHookData{
		time: time.Now(),
		body: c.Body(),
	}
	server.hookChannel <- w
	return nil
}

func (server *HttpServer) HomeHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Status":   "ok",
		"Instance": server.instance,
		"Started":  server.start,
	})
}

func (server *HttpServer) Start() error {

	app := fiber.New()

	fiberlog.SetLevel(fiberlog.LevelInfo)

	server.app = app
	server.instance = uuid.NewString()
	server.start = time.Now()

	app.Get("/", server.HomeHandler)
	app.Post("/hook/:type/:name", server.WebhookHandler)

	return app.Listen(server.config.Host)
}
