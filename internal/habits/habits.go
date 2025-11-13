package habits

import (
	"database/sql"
	"fmt"
	"time"
)

// Habit represents a tracked habit
type Habit struct {
	ID        int64
	Name      string
	Count     int
	CreatedAt time.Time
}


// print command menu
func PrintCommands() {
		fmt.Println("Commands:")
		fmt.Printf("  %-20s %s\n", "habit add [name]", "Add a new habit")
		fmt.Printf("  %-20s %s\n", "habit do  [name]", "Mark a habit as done")
		fmt.Printf("  %-20s %s\n", "habit upd [old] [new]", "Update a habit name")
		fmt.Printf("  %-20s %s\n", "habit rm  [name]", "Remove a habit")
		fmt.Printf("  %-20s %s\n", "habit ls", "List all habits")
}

// Add inserts a new habit into the provided database.
func Add(name string, db *sql.DB) error {
	if name == "" {
		return fmt.Errorf("habit name cannot be empty")
	}

	const query = `INSERT INTO habits (name, count, created_at) VALUES (?, ?, ?)`
	createdAt := time.Now().UTC().Format(time.RFC3339)

	res, err := db.Exec(query, name, 0, createdAt)
	if err != nil {
		return fmt.Errorf("failed to add habit: %w", err)
	}

	if id, err := res.LastInsertId(); err == nil {
		_ = id // optionally log or return this value in the future
	}

	return nil
}

// Do increments the count for a habit by the provided amount.
func Do(name string, increment int, db *sql.DB) error {
	if name == "" {
		return fmt.Errorf("habit name cannot be empty")
	}

	const query = `UPDATE habits SET count = count + ? WHERE name = ?`
	_, err := db.Exec(query, increment, name)
	return err
}

// Remove deletes a habit by name.
func Remove(name string, db *sql.DB) error {
	if name == "" {
		return fmt.Errorf("habit name cannot be empty")
	}
	const query = `DELETE FROM habits WHERE name = ?`
	_, err := db.Exec(query, name)
	return err
}

// Update changes the name of a habit from oldName to newName.
func Update(oldName, newName string, db *sql.DB) error {
	if oldName == "" || newName == "" {
		return fmt.Errorf("both old and new habit names are required")
	}
	const query = `UPDATE habits SET name = ? WHERE name = ?`
	res, err := db.Exec(query, newName, oldName)
	if err != nil {
		return fmt.Errorf("failed to update habit: %w", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err == nil && rowsAffected == 0 {
		return fmt.Errorf("no habit found with the name '%s'", oldName)
	}
	return err
}

// List returns all habits from the database.
func List(db *sql.DB) ([]Habit, error) {
	const query = `SELECT id, name, count, created_at FROM habits`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Habit
	var created string
	for rows.Next() {
		var h Habit
		if err := rows.Scan(&h.ID, &h.Name, &h.Count, &created); err != nil {
			return nil, err
		}
		if t, err := time.Parse(time.RFC3339, created); err == nil {
			h.CreatedAt = t
		}
		out = append(out, h)
	}
	return out, rows.Err()
}
