package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strconv"
)

var Version string

type DBConfig struct {
	Dsn                string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxIdleTime        string
}
type Config struct {
	Port int
	Env  string
	DB   DBConfig
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

	// Database  config
	Dsn := os.Getenv("DATABASE_URL")
	MaxOpenConnections, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNECTIONS"))
	if err != nil {
		slog.Error(err.Error())
	}
	MaxIdleConnections, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTIONS"))
	if err != nil {
		slog.Error(err.Error())
	}
	MaxIdleTime := os.Getenv("MAX_IDLE_TIME")

	dbConfig := DBConfig{
		Dsn:                Dsn,
		MaxOpenConnections: MaxOpenConnections,
		MaxIdleConnections: MaxIdleConnections,
		MaxIdleTime:        MaxIdleTime,
	}

	return &Config{
		Port: port,
		Env:  env,
		DB:   dbConfig,
	}
}
