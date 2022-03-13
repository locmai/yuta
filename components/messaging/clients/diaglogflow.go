package clients

import (
	"context"
	"fmt"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/locmai/yuta/components/messaging/config"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

type DiaglogflowClient struct {
	dialogflow.SessionsClient
	ProjectID    string
	LanguageCode string
}

func NewDiaglogflowClient(c config.NluClientConfig) (*DiaglogflowClient, error) {
	ctx := context.Background()
	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return nil, err
	}
	return &DiaglogflowClient{
		SessionsClient: *sessionClient,
		ProjectID:      c.ProjectID,
		LanguageCode:   c.LanguageCode,
	}, nil
}

func (d DiaglogflowClient) DetectIntentText(sessionID, text string) (string, string, error) {
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return "", "", err
	}
	defer sessionClient.Close()

	if d.ProjectID == "" || sessionID == "" {
		return "", "", fmt.Errorf("received empty project (%s) or session (%s)", d.ProjectID, sessionID)
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", d.ProjectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: d.LanguageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	actionDetected := queryResult.Action
	return fulfillmentText, actionDetected, nil
}
