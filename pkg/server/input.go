package server

import (
	"time"
)

type WebHookInput interface {
	Body() []byte
	Time() time.Time
}

type WebHookData struct {
	body []byte
	time time.Time
}

func (d *WebHookData) Body() []byte {
	return d.body
}

func (d *WebHookData) Time() time.Time {
	return d.time
}
