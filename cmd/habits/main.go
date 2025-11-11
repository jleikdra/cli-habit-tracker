package main

import (
	"fmt"
	"os"

	dbpkg "github.com/jleikdra/cli-habit-tracker/internal/db"
	habitpkg "github.com/jleikdra/cli-habit-tracker/internal/habits"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Type 'habits help' for instructions.")
		return
	}

	// Open or create the SQLite database in the working directory.
	dbPath := "habits.db"
	db, err := dbpkg.OpenDB(dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open database: %v\n", err)
		os.Exit(1)
	}
	// find out more about defer: https://go.dev/ref/spec#Defer_statements
	defer db.Close()

	// Ensure schema exists.
	if _, err := db.Exec(dbpkg.Schema); err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize schema: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "help":
		fmt.Println("Commands:")
		fmt.Printf("  %-20s %s\n", "habit add [name]", "Add a new habit")
		fmt.Printf("  %-20s %s\n", "habit do  [name]", "Mark a habit as done")
		fmt.Printf("  %-20s %s\n", "habit rm  [name]", "Remove a habit")
		fmt.Printf("  %-20s %s\n", "habit upd [old] [new]", "Update a habit name")
		fmt.Printf("  %-20s %s\n", "habit ls", "List all habits")

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error. Habit name required.")
			os.Exit(1)
		}
		name := os.Args[2]
		if err := habits.Add(name, db); err != nil {
			fmt.Fprintf(os.Stderr, "failed to add habit: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Habit added")

	case "do":
		if len(os.Args) < 3 {
			fmt.Println("Error. Habit name required.")
			os.Exit(1)
		}
		name := os.Args[2]
		if err := habits.Do(name, 1, db); err != nil {
			fmt.Fprintf(os.Stderr, "failed to mark habit done: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Habit marked as done")

	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("Error. Habit name required.")
			os.Exit(1)
		}
		name := os.Args[2]
		if err := habits.Remove(name, db); err != nil {
			fmt.Fprintf(os.Stderr, "failed to remove habit: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Habit removed")

	case "upd":
		if len(os.Args) < 4 {
			fmt.Println("Error. Old and new habit names required.")
			os.Exit(1)
		}
		oldName := os.Args[2]
		newName := os.Args[3]
		if err := habits.Update(oldName, newName, db); err != nil {
			fmt.Fprintf(os.Stderr, "failed to update habit: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Habit updated")

	case "ls":
		list, err := habits.List(db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to list habits: %v\n", err)
			os.Exit(1)
		}
		if len(list) == 0 {
			fmt.Println("No habits found")
			return
		}
		for _, h := range list {
			fmt.Printf("%d. %s (count: %d)\n", h.ID, h.Name, h.Count)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run habit help for instructions.")
	}

}
