package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strconv"
)

var Version string

type Config struct {
	Port int
	Env  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
	}

	// App config
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		slog.Error(err.Error())
	}
	env := os.Getenv("APP_ENV")

	return &Config{
		Port: port,
		Env:  env,
	}
}
