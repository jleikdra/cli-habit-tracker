package main

import (
	"fmt"
	"os"

	dbpkg "github.com/jleikdra/cli-habit-tracker/internal/db"
	habits "github.com/jleikdra/cli-habit-tracker/internal/habits"
)

func main() {
	os.Exit(run())
}

func run() int {
	// early validation check
	if len(os.Args) < 2 {
		habits.PrintCommands()
		return 0
	}

	// open or create the SQLite database in the working directory.
	dbPath := "habits.db"
	db, err := dbpkg.OpenDB(dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open database: %v\n", err)
		return 1
	}

	// close all connections after exit
	defer db.Close()

	// ensure schema exists.
	if _, err := db.Exec(dbpkg.Schema); err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize schema: %v\n", err)
		return 1
	}

	command := os.Args[1]

	switch command {
	case "help":
		habits.PrintCommands()
		return 0

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error. Habit name required.")
			return 1
		}
		name := os.Args[2]
		if err := habits.Add(name, db); err != nil {
			fmt.Fprintf(os.Stderr, "failed to add habit: %v\n", err)
			return 1
		}
		fmt.Println("Habit added")
		return 0

	case "do":
		if len(os.Args) < 3 {
			fmt.Println("Error. Habit name required.")
			return 1
		}
		name := os.Args[2]
		if err := habits.Do(name, 1, db); err != nil {
			fmt.Fprintf(os.Stderr, "failed to mark habit done: %v\n", err)
			return 1
		}
		fmt.Println("Habit marked as done")
		return 0

	case "upd":
		if len(os.Args) < 4 {
			fmt.Println("Error. Old and new habit names required.")
			return 1
		}
		oldName := os.Args[2]
		newName := os.Args[3]
		if err := habits.Update(oldName, newName, db); err != nil {
			fmt.Fprintf(os.Stderr, "failed to update habit: %v\n", err)
			return 1
		}
		fmt.Println("Habit updated")
		return 0

	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("Error. Habit name required.")
			return 1
		}
		name := os.Args[2]
		if err := habits.Remove(name, db); err != nil {
			fmt.Fprintf(os.Stderr, "failed to remove habit: %v\n", err)
			return 1
		}
		fmt.Println("Habit removed")
		return 0

	case "ls":
		list, err := habits.List(db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to list habits: %v\n", err)
			return 1
		}
		if len(list) == 0 {
			fmt.Println("No habits found")
			return 0
		}
		for _, h := range list {
			fmt.Printf("%d. %s (count: %d)\n", h.ID, h.Name, h.Count)
		}
		return 0

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run habit help for instructions.")
		return 1
	}
}
