package types

import "time"

const WebHookTokenMaxSize = 128

type WebHookEventType uint8

const (
	WebHookHardBounce WebHookEventType = 1
	WebHookSoftBounce WebHookEventType = 2
)

type WebHook interface {
	Email() string
	Time() time.Time
	EventType() WebHookEventType
}

type WebHookType string

const (
	WebHookTypeSarbacane WebHookType = "sarbacane"
)

type WebHookInput interface {
	Body() []byte
	Time() time.Time
}
