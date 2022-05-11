package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var pool *sqlx.DB

func Establish(connectionStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create database object: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// コネクションが確立した
	pool = db
	return pool, nil
}

func Pool() *sqlx.DB {
	return pool
}
