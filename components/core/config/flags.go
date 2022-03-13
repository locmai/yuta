package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/locmai/yuta/common"
	"github.com/sirupsen/logrus"
)

var (
	configPath = flag.String("config", "config.yaml", "The path to the config file. For more information, see the config file in this repository.")
	version    = flag.Bool("version", false, "Shows the current version and exits immediately.")
)

// ParseFlags parses the commandline flags and uses them to create a config.
func ParseFlags() *MessagingConfig {
	flag.Parse()

	if *version {
		fmt.Println(common.VersionString())
		os.Exit(0)
	}

	if *configPath == "" {
		logrus.Fatal("--config must be supplied")
	}

	cfg, err := Load(*configPath)

	if err != nil {
		logrus.Fatalf("Invalid config file: %s", err)
	}

	return cfg
}
