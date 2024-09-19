package config

import (
	"github.com/joho/godotenv"
	"log"
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
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	env := os.Getenv("APP_ENV")
	return &Config{
		Port: port,
		Env:  env,
	}
}
