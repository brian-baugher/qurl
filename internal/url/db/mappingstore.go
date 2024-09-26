package db

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

type MappingStore struct {
	Db *sql.DB
}

func NewMappingStore() (*MappingStore, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "mappings",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &MappingStore{Db: db}, nil
}
