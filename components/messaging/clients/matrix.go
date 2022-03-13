package clients

import (
	"fmt"

	"github.com/locmai/yuta/components/messaging/config"
	"github.com/locmai/yuta/components/messaging/producers"
	"github.com/sirupsen/logrus"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
)

type MatrixClient struct {
	mautrix.Client
	NluClient    NluClient
	StartTime    int64
	CoreProducer producers.CoreProducer
}

// Start listening on client /sync streams
func (c *MatrixClient) StartSyncer() {
	syncer := c.Client.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, evt *event.Event) {
		if evt.Sender == "@locmai:dendrite.maibaloc.com" && evt.Timestamp > c.StartTime {
			logrus.Printf("<%[1]s> %[4]s (%[2]s/%[3]s)\n", evt.Sender, evt.Type.String(), evt.ID, evt.Content.AsMessage().Body)
			c.SendText(evt.RoomID, evt.Content.AsMessage().Body)

			_, action, err := c.NluClient.DetectIntentText("test", evt.Content.AsMessage().Body)
			if err != nil {
				panic(err)
			}
			if action == "scale" {
				c.CoreProducer.SendActionData("test-name", "test-namespace", action)
			}
			c.SendText(evt.RoomID, fmt.Sprintf("Action detected: %s", action))
		}
	})
	go c.Sync()
}

func NewMatrixClient(c config.ChatClientConfig, startTime int64, nluClient NluClient, coreProducer producers.CoreProducer) (*MatrixClient, error) {
	client, err := mautrix.NewClient(c.HomeserverURL, "", "")
	if err != nil {
		panic(err)
	}

	authType := mautrix.AuthTypePassword

	matrixClient := MatrixClient{
		NluClient:    nluClient,
		StartTime:    startTime,
		CoreProducer: coreProducer,
	}

	if c.Password == "" {
		authType = mautrix.AuthTypeToken
		_, err := client.Login(&mautrix.ReqLogin{
			Type:             authType,
			Identifier:       mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: c.Username},
			Token:            c.Token,
			StoreCredentials: true,
		})
		if err != nil {
			panic(err)
		}
		matrixClient.Client = *client
		return &matrixClient, nil
	}

	_, err = client.Login(&mautrix.ReqLogin{
		Type:             authType,
		Identifier:       mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: c.Username},
		Password:         c.Password,
		StoreCredentials: true,
	})
	if err != nil {
		panic(err)
	}

	matrixClient.Client = *client

	return &matrixClient, nil
}
