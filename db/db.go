package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection struct {
	Session *sqlx.DB
}

func NewConnection(dbURL string) (*Connection, error) {
	conn, err := sqlx.Connect("postgres", dbURL+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &Connection{conn}, nil
}
