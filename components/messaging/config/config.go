package config

import (
	"io/ioutil"

	"github.com/locmai/yuta/common"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type ChatClientType string

const (
	MatrixType ChatClientType = "matrix"
	SlackType  ChatClientType = "slack"
)

type ChatClientConfig struct {
	// The chat platform's username to connect with.
	Username string `yaml:"username"`
	// The password or the token to authenticate the requests with.
	Password string `yaml:"password"`
	// The token to authenticate the requests with.
	Token string `yaml:"token"`
	// A URL with the host and port of the matrix server.
	HomeserverURL string `yaml:"homeserverurl"`
	// The desired display name for this client.
	DisplayName string `yaml:"displayname"`
	// The type of this client.
	ChatClientType ChatClientType `yaml:"type"`
}

type NluClientType string

const (
	DiaglogflowClientType NluClientType = "diaglogflow"
	LuisClientType        NluClientType = "luis"
)

type NluClientConfig struct {
	// ProjectID
	ProjectID string `yaml:"projectid"`
	// LanguageCode
	LanguageCode string `yaml:"languagecode"`
	// Type
	Type NluClientType `yaml:"type"`
}

type MessagingConfig struct {
	common.CommonConfig `yaml:"common,inline"`
	ChatClients         []ChatClientConfig `yaml:"chatClients"`
	NluClients          []NluClientConfig  `yaml:"nluClients"`
}

// Load the ConfigFile
func Load(configPath string) (*MessagingConfig, error) {
	var config MessagingConfig

	configData, err := ioutil.ReadFile(configPath)

	if err != nil {
		return nil, err
	}
	// Pass the current working directory and ioutil.ReadFile so that they can
	// be mocked in the tests
	if err = yaml.Unmarshal(configData, &config); err != nil {
		return nil, err
	}

	logrus.Info(config.ChatClients[0].ChatClientType)
	logrus.Info(config.Metrics.Enabled)
	return &config, nil
}
