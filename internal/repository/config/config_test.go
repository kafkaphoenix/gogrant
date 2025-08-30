//go:build unit

package config_test

import (
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/kafkaphoenix/gogrant/internal/repository/config"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
	logger *slog.Logger
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func (s *ConfigTestSuite) SetupSuite() {
	config.MockConfigReset()
	s.logger = slog.New(slog.NewTextHandler(io.Discard, nil))
}

func (s *ConfigTestSuite) TearDownSuite() {
	config.MockConfigReset()
}

func (s *ConfigTestSuite) TearDownTest() {
	config.MockConfigReset()
	s.logger = slog.New(slog.NewTextHandler(io.Discard, nil))
}

func (s *ConfigTestSuite) TestLoadConfig_OK() {
	// GIVEN
	os.Setenv("APP_PORT", "8080")
	os.Setenv("APP_DEBUG_ENABLED", "true")
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	os.Setenv("MONGO_DATABASE", "grant")

	// WHEN
	cfg, err := config.Load(s.logger)

	// THEN
	s.NoError(err)
	s.Equal(8080, cfg.App.Port)
	s.Equal(true, cfg.App.DebugEnabled)
	s.Equal("mongodb://localhost:27017", cfg.Mongo.URI)
	s.Equal("grant", cfg.Mongo.Database)
	os.Unsetenv("APP_PORT")
	os.Unsetenv("APP_DEBUG_ENABLED")
	os.Unsetenv("MONGO_URI")
	os.Unsetenv("MONGO_DATABASE")
}

func (s *ConfigTestSuite) TestLoadConfig_UnmarshalError() {
	// GIVEN
	os.Setenv("APP_PORT", "invalid_port") // This should cause an error during unmarshalling

	// WHEN
	cfg, err := config.Load(s.logger)

	// THEN
	s.Error(err)
	s.Nil(cfg)
	s.Contains(err.Error(), "error unmarshalling config")
	os.Unsetenv("APP_PORT")
}
