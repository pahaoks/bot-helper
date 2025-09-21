package config

import (
	"github.com/koding/multiconfig"
)

// AppConfig - generic config for all apps
type AppConfig struct {
	Prefix      string `default:"app-prefix"`
	Name        string `default:"app-name"`
	Environment string `default:"dev"` // 'dev' or 'prod' or 'preprod'
}

// IsProd - is production environment
func (c AppConfig) IsProd() bool {
	return c.Environment == "prod"
}

// Load config, parameter must be pointer
func Load(c any) {
	LoadEnv()

	loader := &multiconfig.DefaultLoader{
		Loader: multiconfig.MultiLoader(
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		),
		Validator: &multiconfig.RequiredValidator{},
	}
	loader.MustLoad(c)
}
