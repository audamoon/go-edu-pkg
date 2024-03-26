package connect

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	SSLMode  string
}

func NewPostgresInstance(c *DBConfig) (*sql.DB, error) {
	connSettings := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.Host,
		c.Port,
		c.DBName,
		c.User,
		c.Password,
		c.SSLMode,
	)

	db, err := sql.Open("postgres", connSettings)
	if err != nil {
		return nil, fmt.Errorf("db init: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}

	return db, nil
}
