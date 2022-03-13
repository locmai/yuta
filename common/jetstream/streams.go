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
)

var streams = []*nats.StreamConfig{
	{
		Name:        ActionableItemEvent,
		Retention:   nats.WorkQueuePolicy,
		Description: "Actionale events that Core could pick up and process",
		Storage:     nats.FileStorage,
	}}
