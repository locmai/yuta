package jetstream

import (
	"strings"

	"github.com/locmai/yuta/common"
	natsclient "github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func Prepare(cfg *common.JetStream) natsclient.JetStreamContext {
	// check if we need an in-process NATS Server
	nc, err := natsclient.Connect(strings.Join(cfg.Addresses, ","))
	if err != nil {
		logrus.WithError(err).Panic("Unable to connect to NATS")
		return nil
	}
	s, err := nc.JetStream()
	if err != nil {
		logrus.WithError(err).Panic("Unable to get JetStream context")
		return nil
	}

	for _, stream := range streams {
		name := cfg.TopicFor(stream.Name)
		info, err := s.StreamInfo(name)
		if err != nil && err != natsclient.ErrStreamNotFound {
			logrus.WithError(err).Fatal("Unable to get stream info")
		}
		if info == nil {
			stream.Subjects = []string{name}

			// Namespace the streams without modifying the original streams
			// array, otherwise we end up with namespaces on namespaces.
			namespaced := *stream
			namespaced.Name = name
			if _, err = s.AddStream(&namespaced); err != nil {
				logrus.WithError(err).WithField("stream", name).Fatal("Unable to add stream")
			}
		}
	}
	return s
}
