package habits

import (
    "database/sql"
    "os"
    "testing"

    dbpkg "github.com/jleikdra/cli-habit-tracker/internal/db"
    _ "github.com/mattn/go-sqlite3"
)

// openTestDB creates a temp sqlite file, initializes schema and returns *sql.DB.
func openTestDB(t *testing.T) *sql.DB {
    t.Helper()
    f, err := os.CreateTemp("", "habits_test_*.db")
    if err != nil {
        t.Fatalf("temp file: %v", err)
    }
    path := f.Name()
    f.Close()
    t.Cleanup(func() { os.Remove(path) })

    db, err := sql.Open("sqlite3", path)
    if err != nil {
        t.Fatalf("open db: %v", err)
    }
    t.Cleanup(func() { db.Close() })

    if _, err := db.Exec(dbpkg.Schema); err != nil {
        t.Fatalf("init schema: %v", err)
    }
    return db
}

func TestAddDoListRemove(t *testing.T) {
    db := openTestDB(t)

    // Add a habit
    if err := Add("reading", db); err != nil {
        t.Fatalf("Add: %v", err)
    }

    // List and verify
    list, err := List(db)
    if err != nil {
        t.Fatalf("List: %v", err)
    }
    if len(list) != 1 {
        t.Fatalf("expected 1 habit, got %d", len(list))
    }
    h := list[0]
    if h.Name != "reading" {
        t.Fatalf("expected name reading, got %s", h.Name)
    }
    if h.Count != 0 {
        t.Fatalf("expected count 0, got %d", h.Count)
    }

    // Do the habit
    if err := Do("reading", 3, db); err != nil {
        t.Fatalf("Do: %v", err)
    }

    list, err = List(db)
    if err != nil {
        t.Fatalf("List after Do: %v", err)
    }
    if list[0].Count != 3 {
        t.Fatalf("expected count 3, got %d", list[0].Count)
    }

    // Remove it
    if err := Remove("reading", db); err != nil {
        t.Fatalf("Remove: %v", err)
    }
    list, err = List(db)
    if err != nil {
        t.Fatalf("List after Remove: %v", err)
    }
    if len(list) != 0 {
        t.Fatalf("expected 0 habits after remove, got %d", len(list))
    }
}

func TestUpdateAndValidation(t *testing.T) {
    db := openTestDB(t)

    // Update a non-existent habit should return an error
    if err := Update("nope", "new", db); err == nil {
        t.Fatalf("expected error updating non-existent habit")
    }

    // Empty name validations
    if err := Add("", db); err == nil {
        t.Fatalf("expected error when adding empty name")
    }
    if err := Do("", 1, db); err == nil {
        t.Fatalf("expected error when doing empty name")
    }
    if err := Remove("", db); err == nil {
        t.Fatalf("expected error when removing empty name")
    }
}
