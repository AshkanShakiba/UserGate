package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

type MysqlConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"3306"`
	User     string `env:"DB_USER" envDefault:"root"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
	DbName   string `env:"DB_NAME" envDefault:"usergate"`
}

type Configuration struct {
	Env      string `env:"ENVIRONMENT" envDefault:"staging"`
	AppName  string `env:"APP_NAME"  envDefault:"UserGate"`
	Port     string `env:"PORT"      envDefault:"8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	Mysql    MysqlConfig
}

func Configure() (*Configuration, error) {
	cfg := &Configuration{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("parsing configuration error: %w", err)
	}

	return cfg, nil
}
