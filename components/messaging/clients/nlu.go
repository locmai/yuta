package clients

import structpb "github.com/golang/protobuf/ptypes/struct"

type NluClient interface {
	DetectIntentText(sessionID, text string) (string, map[string]*structpb.Value, string, error)
}
