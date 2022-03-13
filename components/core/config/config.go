package config

import (
	"io/ioutil"

	"github.com/locmai/yuta/common"
	"gopkg.in/yaml.v2"
)

type AppServiceType string

const (
	KubeopsAppService    AppServiceType = "kubeops"
	PrometheusAppService AppServiceType = "prometheus"
)

type AppService struct {
	AppServiceType AppServiceType `yaml:"type"`
}

type CoreConfig struct {
	common.CommonConfig
	AppServices []AppService
}

// Load the ConfigFile
func Load(configPath string) (*CoreConfig, error) {
	var config CoreConfig

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
