package db_test

import (
	"database/sql"
	"path/filepath"
	"testing"

	dbpkg "github.com/jleikdra/cli-habit-tracker/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func TestOpenDBCreatesSchema(t *testing.T) {
	tmp := t.TempDir()
	dbPath := filepath.Join(tmp, "test.db")
	db, err := dbpkg.OpenDB(dbPath)
	if err != nil {
		t.Fatalf("OpenDB failed: %v", err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='habits'").Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			t.Fatalf("habits table not found")
		}
		t.Fatalf("query failed: %v", err)
	}
	if name != "habits" {
		t.Fatalf("unexpected table name: %s", name)
	}
}
