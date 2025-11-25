package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Schema is the SQL used to create the initial database schema for the
// application. It's exported so callers can initialize the DB on startup.
const Schema = `
	CREATE TABLE IF NOT EXISTS habits (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	count INTEGER DEFAULT 0,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

func OpenDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// ensure the schema exists; execute the Schema SQL
	if _, err := db.Exec(Schema); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
