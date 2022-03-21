package jetstream

import (
	"github.com/nats-io/nats.go"
)

const (
	Name      = "name"
	Namespace = "namespace"
)

var (
	ActionableItemEvent = "ActionableItemEvent"
	MessageToSent       = "TextToSent"
)

var streams = []*nats.StreamConfig{
	{
		Name:        ActionableItemEvent,
		Retention:   nats.WorkQueuePolicy,
		Description: "Actionale events that core server could pick up and process",
		Storage:     nats.FileStorage,
	},
	{
		Name:        MessageToSent,
		Retention:   nats.WorkQueuePolicy,
		Description: "Text message to be sent by the messaging server",
		Storage:     nats.MemoryStorage,
	},
}
