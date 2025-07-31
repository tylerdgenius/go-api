package http

import (
	"api-template/internal/constants"
	utils "api-template/internal/utils"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PORT        string `env:"PORT" env-default:"8080"`
	TIMEOUT     string `env:"TIMEOUT" env-default:"30"`
	SERVER_HOST string `env:"SERVER_HOST" env-default:"localhost"`
}

func LoadConfig() (*Config, error) {

	var err error
	var cfg Config

	if utils.EnvExists(constants.ENV_FILE) {
		err = cleanenv.ReadConfig(constants.ENV_FILE, &cfg)
	} else {
		err = cleanenv.ReadEnv(&cfg)
	}

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
