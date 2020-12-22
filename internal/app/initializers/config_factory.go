package initializers

import (
	"errors"

	"github.com/devopsfaith/krakend/config"
	"github.com/rs/zerolog/log"
)

const (
	defaultCfgPath              = "./cfg/backend_cfg_default.json"
	defaultServiceConfigVersion = 2
	defaultServiceName          = "ApiProxy"
	defaultServicePort          = 80
	deafultServiceCacheTTL      = 0
	defaultServiceTimeout       = 1500000
	defaultServiceIdleTimeout   = 1500000
	defaultOutputEncoding       = "json"
	loadEndpointsTimeout        = 10
)

var (
	ErrUnmarshalFile                = errors.New("could not unmarshall file")
	ErrOpenFile                     = errors.New("could not open file")
	ErrIncorrectDefaultConfigsPaths = errors.New("incorrect default configs paths")
)

//TODO implement config builder
/*
1. Initialize default service config with default JSON file/values
2. Try to load service config from spf13/viper
*/
func ConfigFactory() (*config.ServiceConfig, error) {
	cfg := &config.ServiceConfig{}

	log.Debug().Interface("Config", cfg).Msg("compiled krakend config")

	if err := cfg.Init(); err != nil {
		return nil, err
	}

	return cfg, nil
}
