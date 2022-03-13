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
func (p *CoreProducer) SendActionData(name string, namespace string, action string) error {
	m := &nats.Msg{
		Subject: p.Topic,
		Header:  nats.Header{},
	}

	data := utils.KubeopsActionData{
		Action:    action,
		Name:      name,
		Namespace: namespace,
	}

	var err error
	m.Data, err = json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = p.JetStream.PublishMsg(m)
	return err
}
