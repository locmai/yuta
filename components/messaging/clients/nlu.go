package clients

type NluClient interface {
	DetectIntentText(sessionID, text string) (string, string, error)
}
