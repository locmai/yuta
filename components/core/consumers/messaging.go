package consumers

import (
	"context"
	"encoding/json"

	"github.com/locmai/yuta/common"
	"github.com/locmai/yuta/common/jetstream"
	"github.com/locmai/yuta/common/utils"
	"github.com/locmai/yuta/components/core/appservices"
	"github.com/locmai/yuta/components/core/config"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

type ActionableItemEventConsumer struct {
	ctx               common.ProcessContext
	JetStream         nats.JetStreamContext
	Topic             string
	Durable           string
	KubeopsAppService appservices.KubeopsAppService
}

// NewOutputRoomEventConsumer creates a new OutputRoomEventConsumer. Call Start() to begin consuming from room servers.
func NewActionableItemEventConsumer(
	ctx common.ProcessContext,
	cfg *config.CoreConfig,
	js nats.JetStreamContext,
) *ActionableItemEventConsumer {
	return &ActionableItemEventConsumer{
		ctx:       ctx,
		JetStream: js,
		Topic:     cfg.JetStream.TopicFor(jetstream.ActionableItemEvent),
		Durable:   cfg.JetStream.Durable(jetstream.ActionableItemEvent),
	}
}

// Start consuming from room servers
func (s *ActionableItemEventConsumer) Start() error {
	return jetstream.JetStreamConsumer(
		s.ctx, s.JetStream, s.Topic, s.Durable, s.onMessage,
		nats.DeliverAll(), nats.ManualAck(),
	)
}

func (s *ActionableItemEventConsumer) onMessage(ctx context.Context, msg *nats.Msg) bool {
	var actionableItem utils.KubeopsActionData
	if err := json.Unmarshal(msg.Data, &actionableItem); err != nil {
		// If the message was invalid, log it and move on to the next message in the stream
		logrus.WithError(err).Errorf("Core server output log: message parse failure")
		return true
	}

	s.KubeopsAppService.Act(appservices.Action(actionableItem.Action), appservices.ObjectMeta{
		Name:      actionableItem.Name,
		Namespace: actionableItem.Namespace,
		Value:     actionableItem.Value,
	})

	logrus.Printf("Ack actionableItem.Action %s", actionableItem.Action)
	return true
}
