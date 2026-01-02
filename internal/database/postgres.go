package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/PranavJoshi2893/boilerplate/internal/config"
)

func NewPostgres(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", cfg.DBUser, cfg.DBName, cfg.DBSSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
