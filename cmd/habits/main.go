package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Type 'habits help' for instructions.")
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
		fmt.Println("Habit added")

	case "do":
		if len(os.Args) < 3 {
			fmt.Println("Error. Habit name required.")
			os.Exit(1)
		}
		fmt.Println("Habit marked as done")

	case "rm":
		if len(os.Args) < 3 {
			fmt.Println("Error. Habit name required.")
			os.Exit(1)
		}
		fmt.Println("Habit removed")

	case "upd":
		if len(os.Args) < 4 {
			fmt.Println("Error. Old and new habit names required.")
			os.Exit(1)
		}
		fmt.Println("Habit updated")

	case "ls":
		fmt.Println("Habit list")

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run habit help for instructions.")
	}

}
