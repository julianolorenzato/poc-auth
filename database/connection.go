package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func OpenConnection() (*sql.DB, error) {
	// Open db connection
	conn, err := sql.Open("sqlite3", "dev.db")
	if err != nil {
		return nil, err
	}

	// Test db connection
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	// Get the db driver
	driver, err := sqlite.WithInstance(conn, &sqlite.Config{})
	if err != nil {
		return nil, err
	}

	// Create a new migrator
	migrator, err := migrate.NewWithDatabaseInstance("file://./migrations", "sqlite", driver)
	if err != nil {
		return nil, err
	}

	// Perform an "up" migration
	migrator.Up()

	return conn, nil
}
