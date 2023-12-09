package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/grippenet/webhook-service/pkg/types"
)

type AppConfig struct {
	Port  uint16
	Hooks map[string]struct{}
}

func LoadConfig() *AppConfig {
	conf := AppConfig{}
	p := os.Getenv("WEBHOOK_SERVICE_LISTEN_PORT")
	if p == "" {
		conf.Port = 3253
	} else {
		port, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("Invalid port number WEBHOOK_SERVICE_LISTEN_PORT must be an integer: %s", err)
		}
		if port > 65535 || port <= 0 {
			log.Fatalf("Invalid port number must be between 1 and 65535")
		}
		conf.Port = uint16(port)
	}

	h, err := parseHooks(os.Getenv("WEBHOOK_SERVICE_HOOKS"))
	if err != nil {
		log.Fatalf("Error in WEBHOOK_SERVICE_HOOKS : %s", err)
	}
	conf.Hooks = h

	//conf.MessageDBConfig = config.GetMessageDBConfig()
	return &conf
}

func (conf *AppConfig) GetHttpConfig() *types.HttpServerConfig {
	h := types.HttpServerConfig{
		Host:  fmt.Sprintf(":%d", conf.Port),
		Hooks: conf.Hooks,
	}
	return &h
}

var regexpHookName = regexp.MustCompile("[A-Z0-9a-z]+")

func parseHooks(s string) (map[string]struct{}, error) {
	if s == "" {
		return nil, errors.New("at least one hook should be defined")
	}
	hh := strings.Split(s, ",")
	hooks := make(map[string]struct{}, len(hh))
	for idx, h := range hh {
		h = strings.TrimSpace(h)
		if !regexpHookName.MatchString(h) {
			return nil, fmt.Errorf("hook %d : Must be alphanumeric string", idx)
		}
		hooks[h] = struct{}{}
	}
	return hooks, nil
}
