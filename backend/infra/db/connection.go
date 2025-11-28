package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetDBConnectionString() string {
	return "postgres://postgres:mamun1281@localhost:5432/ecommerce?sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	connStr := GetDBConnectionString()
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
