package config

import "fmt"

type ConfigError struct {
	Message string
	Err     error
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}
