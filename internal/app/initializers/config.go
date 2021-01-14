package initializers

import (
	"time"

	"github.com/devopsfaith/krakend/config"
	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

const (
	ViperRemotePrividerTypeEnv     = "VIPER_REMOTE_PROVIDER_TYPE"
	DefaultViperRemotePrividerType = "etcd"

	ViperRemoteProviderPathEnv     = "VIPER_REMOTE_PROVIDER_PATH"
	DefaultViperRemoteProviderPath = "/config/krakend-proxy.json"

	ViperRemoteProviderEndpointEnv     = "VIPER_REMOTE_PROVIDER_ENDPOINT"
	DefaultViperRemoteProviderEndpoint = "http://127.0.0.1:4001"

	ViperRemoteProviderConfigEnv        = "VIPER_REMOTE_PROVIDER_CONFIG_KEY"
	DefaultViperRemoteProviderConfigKey = "krakend.config"

	ViperConfigType         = "json"
	DefaultTestEndpointHost = "https://google.com"
)

//TODO implement config builder
/*
1. Initialize default service config with default JSON file/values
2. Try to load service config from spf13/viper
*/
func InitializeConfig() (*config.ServiceConfig, error) {
	if err := InitializeViper(); err != nil {
		return nil, err
	}

	cfg := NewDefaultServiceConfig()
	key := envy.Get(ViperRemoteProviderConfigEnv, DefaultViperRemoteProviderConfigKey)

	if err := viper.UnmarshalKey(key, cfg); err != nil {
		return nil, err
	}

	log.Debug().Interface("Config", cfg).Msg("compiled krakend config")

	if err := cfg.Init(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func InitializeViper() error {
	err := viper.AddRemoteProvider(
		envy.Get(ViperRemotePrividerTypeEnv, DefaultViperRemotePrividerType),
		envy.Get(ViperRemoteProviderEndpointEnv, DefaultViperRemoteProviderEndpoint),
		envy.Get(ViperRemoteProviderPathEnv, DefaultViperRemoteProviderPath),
	)

	if err != nil {
		return err
	}

	viper.SetConfigType(ViperConfigType)

	return viper.ReadRemoteConfig()
}

func NewDefaultServiceConfig() *config.ServiceConfig {
	return &config.ServiceConfig{
		Name:              "KrakenD proxy",
		Version:           2,
		CacheTTL:          time.Second * 300,
		Timeout:           time.Millisecond * 3000,
		Port:              3000,
		DisableStrictREST: true,
		Endpoints: []*config.EndpointConfig{
			{
				Endpoint:       "/_test",
				Method:         "GET",
				OutputEncoding: "no-op",
				HeadersToPass:  []string{"*"},
				QueryString:    []string{"*"},
				ExtraConfig: config.ExtraConfig{
					"github.com/devopsfaith/krakend/proxy": map[string]interface{}{
						"sequential": false,
					},
				},
				Backend: []*config.Backend{
					{
						HostSanitizationDisabled: true,
						Host:                     []string{DefaultTestEndpointHost},
						URLPattern:               "/",
						Encoding:                 "no-op",
						SD:                       "static",
						Method:                   "GET",
						ExtraConfig:              config.ExtraConfig{},
					},
				},
			},
		},
	}
}
