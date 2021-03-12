package econfig

import (
	"go.uber.org/config"
)

type environment string

const (
	EnvironmentBase        environment = "base.yaml"
	EnvironmentDevelopment environment = "development.yaml"
	EnvironmentProduction  environment = "production.yaml"
	EnvironmentTest        environment = "test.yaml"
)

// Populate the target struct from YAML configs.
func Populate(target interface{}, dir string, env environment) error {
	var opts []config.YAMLOption

	if env != EnvironmentBase {
		opts = append(opts, config.File(dir+string(EnvironmentBase)))
	}

	opts = append(opts, config.File(dir+string(env)))

	provider, err := config.NewYAML(opts...)
	if err != nil {
		return err
	}

	return provider.Get("").Populate(target)
}
