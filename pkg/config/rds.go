package config

import (
	"github.com/caarlos0/env/v6"
)

func init() {
	if err := env.Parse(&RDS); err != nil {
		panic(err)
	}
}

var RDS rds

type rds struct {
	Driver   string `env:"DB_DRIVER"`
	User     string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Database string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}
