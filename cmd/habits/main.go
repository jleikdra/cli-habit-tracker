package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"habit-tracker/internal/db"
	"habit-tracker/internal/habits"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: habit <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  add <name>    Add a new habit")
		os.Exit(1)
	}

	// Get database path (store in current directory for now)
	dbPath := filepath.Join(".", "habits.db")

	// Open database
	database, err := db.Open(dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer database.Close()

	// Parse command
	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("Usage: habit add <name>")
		}
		name := os.Args[2]

		if err := habits.Add(database, name); err != nil {
			log.Fatalf("Error: %v", err)
		}

		fmt.Printf("✓ Added habit: %s\n", name)

	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
