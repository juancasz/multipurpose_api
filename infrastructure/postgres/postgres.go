package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Options struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func Init(opts Options) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		opts.User,
		opts.Password,
		opts.Host,
		opts.Port,
		opts.Name,
		"disable",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
