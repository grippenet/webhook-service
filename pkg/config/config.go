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
	Port            uint16
	Hooks           map[string]struct{}
	HookChannelSize int
	Path            string
}

func GetEnvInt(name string, def int) int {
	p := os.Getenv(name)
	if p == "" {
		return def
	} else {
		value, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalf("Invalid value for %s must be an integer: %s", name, err)
		}
		return value
	}
}

func LoadConfig() *AppConfig {
	conf := AppConfig{}

	port := GetEnvInt("WEBHOOK_SERVICE_LISTEN_PORT", 3253)
	if port > 65535 || port <= 0 {
		log.Fatalf("Invalid port number must be between 1 and 65535")
	}
	conf.Port = uint16(port)

	h, err := parseHooks(os.Getenv("WEBHOOK_SERVICE_HOOKS"))
	if err != nil {
		log.Fatalf("Error in WEBHOOK_SERVICE_HOOKS : %s", err)
	}
	conf.Hooks = h

	path := os.Getenv("WEBHOOK_SERVICE_STORAGE_PATH")

	if path != "" {
		f, err := os.Open(path)
		if err != nil {
			log.Fatalf("Unable to read path %s : %s", path, err)
		}
		fileInfo, err := f.Stat()
		if err != nil {
			log.Fatalf("Unable to file info for path %s : %s", path, err)
		}
		if !fileInfo.IsDir() {
			log.Fatalf("Path must be a directory, found: %s", path)
		}
		conf.Path = path
	}

	hookSize := GetEnvInt("WEBHOOK_SERVICE_CHANNEL_SIZE", 100)
	conf.HookChannelSize = hookSize
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
		if len(h) > types.WebHookTokenMaxSize {
			return nil, fmt.Errorf("hook %d : Must be 128 chars maximum", idx)
		}
		hooks[h] = struct{}{}
	}
	return hooks, nil
}
