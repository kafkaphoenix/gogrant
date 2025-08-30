//go:build unit

package config

import (
	"sync"

	"github.com/spf13/viper"
)

// MockConfigReset reset the config package for testing purposes.
func MockConfigReset() {
	once = sync.Once{}
	cfg = nil
	viper.Reset()
}
