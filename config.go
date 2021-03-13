package econfig

import (
	"go.uber.org/config"
)

// Environment is an enum for the location
// of the config file.
type Environment string

const (
	EnvironmentBase        Environment = "base.yaml"
	EnvironmentDevelopment Environment = "development.yaml"
	EnvironmentProduction  Environment = "production.yaml"
	EnvironmentTest        Environment = "test.yaml"
)

// Populate the target struct from YAML configs.
func Populate(target interface{}, dir string, env Environment) error {
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
