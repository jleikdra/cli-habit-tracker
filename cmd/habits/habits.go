package habits

import (
	"fmt"
	"database/sql"
	"time"
)

// custom habit type
type Habit struct {
	ID        int64
	Name      string
	Count     int
	CreatedAt time.Time
}

// add habit
func Add(name string, db *sql.DB) error {
	// validate inputs
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}
	if name == "" {
		return fmt.Errorf("habit name cannot be empty")
	}

	// insert habit into database
	const query = `INSERT INTO habits (name, count, created_at) VALUES (?)`
	createdAt := time.Now().Format(time.RFC3339)

	//db.Exec returns (Result, error)
	//Result is an interface type
	res, err := db.Exec(query, name, 0, createdAt)
	if err != nil {
		return fmt.Errorf("failed to add habit: %v", err)

	if id, err := res.LastInsertId(); err == nil {
		fmt.Printf("Habit added with ID: %d\n", id)
	}

	return nil
}


// add count
func Do(name string, count int) error {
	// implementation here
	return fmt.Println("Your new count for", name, "is", count)
}

// list habits
func List() ([]string, error) {
	// implementation here
	return nil, nil
}

// remove habit
func Remove(name string) error {
	// implementation here
	return nil
}
