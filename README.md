# cli-habit-tracker
CLI habit tracker written in Go with SQLite for storage.

## Requirements
- Go 1.22+
- CGO enabled (required by `github.com/mattn/go-sqlite3`)

## Usage
```
go run ./cmd/habits <command> [args]
```

Commands:
- `add <name>` — add a new habit
- `do <name>` — increment the habit count by 1
- `upd <old> <new>` — rename a habit
- `rm <name>` — remove a habit
- `ls` — list all habits
- `help` — show the command summary

The app stores data in `habits.db` in the working directory. The schema is created automatically on first run.

## Development
- Build: `go build ./cmd/habits`
- Tests: `go test ./...`

## Project layout
- `cmd/habits` — CLI entrypoint
- `internal/db` — SQLite connection and schema definition
- `internal/habits` — habit CRUD operations
