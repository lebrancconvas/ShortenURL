package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) InitDB(url string) error {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return fmt.Errorf("Failed to connect to Database: %v", err)
	}

	d.db = db
	return nil
}

func (d *Database) CloseDB() error {
	return d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
