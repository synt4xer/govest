package config

import (
	"fmt"
	"time"
)

type Config struct {
	Server struct {
		Port  string `env:"SERVER_PORT"`
		GoEnv string `env:"GO_ENV"`
	}
	Database struct {
		URI                  string        `env:"DATABASE_URI"`
		DatabaseName         string        `env:"DATABASE_NAME"`
		MaxIdleConns         int           `env:"DATABASE_MAX_IDLE_CONNS"`
		MaxOpenConns         int           `env:"DATABASE_MAX_OPEN_CONNS"`
		MaxIdleConnsLifetime time.Duration `env:"DATABASE_MAX_IDLE_CONNS_LIFETIME"`
	}
}

func ProvideConfig() (*Config, error) {
	cfg := &Config{}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	if cfg.Server.Port == "" {
		cfg.Server.Port = ":8080"
	}

	if cfg.Server.GoEnv == "" {
		cfg.Server.GoEnv = "development"
	}

	if cfg.Database.URI == "" {
		return fmt.Errorf("DATABASE_URI is required")
	}

	if cfg.Database.DatabaseName == "" {
		cfg.Database.DatabaseName = "postgres"
	}

	if cfg.Database.MaxIdleConns == 0 {
		cfg.Database.MaxIdleConns = 5
	}

	if cfg.Database.MaxOpenConns == 0 {
		cfg.Database.MaxOpenConns = 95
	}

	if cfg.Database.MaxIdleConnsLifetime == 0 {
		cfg.Database.MaxIdleConnsLifetime = 5 * time.Minute
	}

	return nil
}
