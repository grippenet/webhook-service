package types

type HttpServerConfig struct {
	Host  string
	Hooks map[string]struct{}
}
