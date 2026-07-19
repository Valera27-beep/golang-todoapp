package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LOG_LEVEL" required:"true"`
	Folder string `envconfig:"LOG_FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process logger config: %w", err)
	}

	return config, nil
}

func MustConfig() Config {
	config, err := NewConfig()

	if err != nil {
		panic(fmt.Errorf("failed to load logger config: %w", err))
	}

	return config
}