package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/minhnghia2k3/personal-blog/internal/config"
	"os"
	"time"
)

// ConnectDB connects to the dsn and using provided config
func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	// Add settings
	idleTime, err := time.ParseDuration(cfg.DB.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.DB.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConnections)
	db.SetConnMaxIdleTime(idleTime)

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
