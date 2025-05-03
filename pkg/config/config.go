package config

import (
	"log/slog"

	"github.com/spf13/viper"
)

// ConfigKey represents a viper key used in the configuration.
type ConfigKey string

// Config represents a configuration option.
type Config struct {
	// NameInFile is used to register an alias for the key.
	// It is optional and can be empty.
	//
	// You can use multi level keys like "server.port" or "server.host".
	NameInFile string
	// EnvironmentVar is the name of the environment variable to bind to the key.
	// It is optional and can be empty.
	EnvironmentVar string
	// Key is the viper key used in the configuration.
	// It is required and cannot be empty.
	Key ConfigKey
	// DefaultValue is the default value for the key.
	// It is optional and can be nil.
	DefaultValue any
}

// registerConfigOptions registers the configuration options with viper.
// It binds environment variables, registers aliases, and sets default values.
func registerConfigOptions(configs []Config, viperInstance *viper.Viper) {
	for _, config := range configs {
		if config.EnvironmentVar != "" {
			err := viperInstance.BindEnv(string(config.Key), config.EnvironmentVar)
			if err != nil {
				slog.Error("failed to bind environment variable", "error", err)
			}
		}

		if config.NameInFile != "" {
			viperInstance.RegisterAlias(config.NameInFile, string(config.Key))
		}

		if config.DefaultValue != nil {
			viperInstance.SetDefault(string(config.Key), config.DefaultValue)
		}
	}
}

// registerConfigFilePaths registers the paths where viper will look for configuration files.
func registerConfigFilePaths(viperInstance *viper.Viper) {
	viperInstance.SetConfigType("yaml")
	viperInstance.SetConfigName("config")
	viperInstance.AddConfigPath(".")
	viperInstance.AddConfigPath("/etc/service/")
	viperInstance.AddConfigPath("$HOME/.config/service/")
}

// ConfigureFromEnv configures viper from environment variables.
func ConfigureFromEnv(configs []Config, viperInstance *viper.Viper) {
	registerConfigOptions(configs, viperInstance)

	viperInstance.AutomaticEnv()

	for _, config := range configs {
		if config.EnvironmentVar != "" && viperInstance.IsSet(config.EnvironmentVar) {
			viperInstance.Set(string(config.Key), viperInstance.Get(config.EnvironmentVar))
		}
	}
}

// ConfigureFromConfigFile configures viper from a configuration file.
//
// It reads the YAML configuration file from either the current directory or /etc/app or $HOME/.config/app.
//
// If the configuration file is not found, it logs an error.
// It does not override any values with environment variables.
func ConfigureFromConfigFile(configs []Config, viperInstance *viper.Viper) {
	registerConfigOptions(configs, viperInstance)

	if err := viperInstance.ReadInConfig(); err != nil {
		slog.Error("failed to read config file", "error", err)
	}
}

// AutoConfigure automatically configures viper from both environment variables and a configuration file.
//
// It reads the YAML configuration file from either the current directory or /etc/app or $HOME/.config/app.
//
// It first reads the configuration file and then overrides any values with environment variables.
func AutoConfigure(configs []Config, viperInstance *viper.Viper) {
	registerConfigFilePaths(viperInstance)
	registerConfigOptions(configs, viperInstance)

	viperInstance.AutomaticEnv()

	if err := viperInstance.ReadInConfig(); err != nil {
		slog.Error("failed to read config file", "error", err)
		return
	}

	for _, config := range configs {
		if config.EnvironmentVar != "" && viperInstance.IsSet(config.EnvironmentVar) {
			viperInstance.Set(string(config.Key), viperInstance.Get(config.EnvironmentVar))
		}
	}
}
