package config

import (
	"log/slog"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *Config   //nolint:gochecknoglobals // singleton pattern
	once sync.Once //nolint:gochecknoglobals // singleton pattern
)

type Config struct {
	logger *slog.Logger
	Mongo  struct {
		URI      string `mapstructure:"uri"`
		Database string `mapstructure:"database"`
	} `mapstructure:"mongo"`
	App struct {
		Port         int  `mapstructure:"port"`
		DebugEnabled bool `mapstructure:"debug_enabled"`
	} `mapstructure:"app"`
}

// Load reads the configuration from the environment. It will load the configuration only once
// and return the cached configuration on subsequent calls.
func Load(logger *slog.Logger) (*Config, error) {
	// viper precedence order:
	// 1. explicit call to Set
	// 2. flag
	// 3. env
	// 4. config file
	// 5. key/value store
	// 6. default
	var err error

	once.Do(func() {
		// enable BindStruct to allow unmarshal env into a nested struct
		// https://github.com/spf13/viper/pull/1429
		viper.SetOptions(viper.ExperimentalBindStruct())
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		cfg = &Config{logger: logger.With("component", "config")}
		if unmarshallErr := viper.Unmarshal(cfg); unmarshallErr != nil {
			err = &ConfigError{
				Message: "error unmarshalling config",
				Err:     unmarshallErr,
			}

			return
		}
	})

	if err != nil {
		return nil, err
	}

	cfg.logger.Info("Configuration loaded", "debug_enabled", cfg.App.DebugEnabled)

	return cfg, nil
}
