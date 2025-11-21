package db

import "testing"

func TestOpenDBInMemory(t *testing.T) {
	db, err := OpenDB(":memory:")
	if err != nil {
		t.Fatalf("OpenDB returned error: %v", err)
	}
	t.Cleanup(func() { db.Close() })

	// Exercise the schema to ensure the connection is usable.
	if _, err := db.Exec(Schema); err != nil {
		t.Fatalf("failed to apply schema: %v", err)
	}

	if _, err := db.Exec(`INSERT INTO habits (name, count) VALUES (?, ?)`, "test", 1); err != nil {
		t.Fatalf("failed to write after schema: %v", err)
	}
}
