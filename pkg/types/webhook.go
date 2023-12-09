package types

import "time"

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
