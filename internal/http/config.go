package http

import (
	"api-template/internal/constants"
	utils "api-template/internal/utils"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Env struct {
	PORT        string `env:"PORT" env-default:"8080"`
	TIMEOUT     string `env:"TIMEOUT" env-default:"30"`
	SERVER_HOST string `env:"SERVER_HOST" env-default:"localhost"`
	PG_HOST     string `env:"PG_HOST" env-default:"localhost"`
	PG_PORT     string `env:"PG_PORT" env-default:"5432"`
	PG_USER     string `env:"PG_USER" env-default:"user"`
	PG_PASSWORD string `env:"PG_PASSWORD" env-default:"password"`
	PG_DATABASE string `env:"PG_DATABASE" env-default:"mdm_db"`
}

func LoadConfig() (*Env, error) {

	var err error
	var cfg Env

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

func (e *Env) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		e.PG_HOST, e.PG_PORT, e.PG_USER, e.PG_PASSWORD, e.PG_DATABASE)
}

func (e *Env) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		e.PG_USER, e.PG_PASSWORD, e.PG_HOST, e.PG_PORT, e.PG_DATABASE)
}
