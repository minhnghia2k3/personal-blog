package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"time"
)

type dbConfig struct {
	Dsn                string
	MaxOpenConnections int
	MaxIdleConnections int
	MaxIdleTime        string
}

// ConnectDB connects to the dsn and using provided config
func ConnectDB() (*sql.DB, error) {
	maxIdleConnections, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTIONS"))
	maxOpenConnections, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONNECTIONS"))

	cfg := dbConfig{
		Dsn:                os.Getenv("DATABASE_URL"),
		MaxOpenConnections: maxIdleConnections,
		MaxIdleConnections: maxOpenConnections,
		MaxIdleTime:        os.Getenv("MAX_IDLE_TIME"),
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	// Add settings
	idleTime, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxIdleTime(idleTime)

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
