package config

import (
	"io/ioutil"

	"github.com/locmai/yuta/common"
	"gopkg.in/yaml.v2"
)

type WebhookConfig struct {
	common.CommonConfig `yaml:"common,inline"`
}

// Load the ConfigFile
func Load(configPath string) (*WebhookConfig, error) {
	var config WebhookConfig

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	// Pass the current working directory and ioutil.ReadFile so that they can
	// be mocked in the tests
	if err = yaml.Unmarshal(configData, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
