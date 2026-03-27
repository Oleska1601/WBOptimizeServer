package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Name string `yaml:"name" env:"APP_NAME"`
}

type ServerConfig struct {
	Host            string        `yaml:"host" env:"SERVER_HOST"`
	PortV1          int           `yaml:"port_v1" env:"SERVER_PORT_V1"`
	PortV2          int           `yaml:"port_v2" env:"SERVER_PORT_V2"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SERVER_SHUTDOWN_TIMEOUT"`
}

type LoggerConfig struct {
	Level string `yaml:"level" env:"LOGGER_LEVEL"`
}

type GinConfig struct {
	Mode string `yaml:"mode" env:"GIN_MODE"`
}

type Config struct {
	App    AppConfig    `yaml:"app"`
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
	Gin    GinConfig    `yaml:"gin"`
}

func New() (*Config, error) {
	var cfg Config
	path := "./config/config.yaml"
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return &cfg, nil
}
