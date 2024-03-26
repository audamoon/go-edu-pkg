package connect

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "db init")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "db ping")
	}

	return db, nil
}
