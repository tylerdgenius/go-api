package http

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PORT        string `env:"PORT" env-default:"8080"`
	TIMEOUT     string `env:"TIMEOUT" env-default:"30"`
	SERVER_HOST string `env:"SERVER_HOST" env-default:"localhost"`
}

const envFile = ".env"

func LoadConfig() (*Config, error) {

	var err error
	var cfg Config

	if envExists(envFile) {
		err = cleanenv.ReadConfig(envFile, &cfg)
	} else {
		err = cleanenv.ReadEnv(&cfg)
	}

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func envExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
