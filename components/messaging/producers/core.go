package producers

import (
	"encoding/json"

	"github.com/locmai/yuta/common/utils"
	"github.com/nats-io/nats.go"
)

// CoreProducer produces events for the Core server to consume
type CoreProducer struct {
	Topic     string
	JetStream nats.JetStreamContext
}

// SendActionData sends action data to the core server
func (p *CoreProducer) SendActionData(inputData utils.KubeopsActionData) error {
	m := &nats.Msg{
		Subject: p.Topic,
		Header:  nats.Header{},
	}

	data := inputData

	var err error
	m.Data, err = json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = p.JetStream.PublishMsg(m)
	return err
}
