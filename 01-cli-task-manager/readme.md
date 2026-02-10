# CLI Task Manager (Go)

A simple command-line task management application with JSON persistence.

## Folder Structure
```

01-cli-task-manager/
├── main.go        # CLI entry point, command parsing
├── task.go        # Task struct definition
├── storage.go    # File I/O and JSON handling
├── tasks.json    # Data file (auto-created)
└── go.mod

````

## Files Overview

### main.go
- Parses CLI arguments using `flag`
- Handles commands: add, list, complete, delete
- Calls storage functions to load/save tasks

### task.go
- Defines `Task` struct:
  - `ID` (int)
  - `Title` (string)
  - `Completed` (bool)
  - `Priority` (int)

### storage.go
- Reads tasks from `tasks.json`
- Writes tasks to `tasks.json`
- Uses `encoding/json` and `os`

### tasks.json
- Stores tasks persistently in JSON format
- Created automatically on first run

## Build & Run

### Initialize module
```bash
go mod init cli-task-manager
````

### Build binary

```bash
go build -o taskmgr
```

### Run binary

```bash
./taskmgr -add "Buy milk" -priority 2
./taskmgr -list
./taskmgr -complete 1
./taskmgr -status completed
./taskmgr -delete 1
```

### Run without building

```bash
go run . -add "Test task"
```

## CLI Commands

* Add task
  `-add "task name" -priority <number>`

* List tasks
  `-list`

* Filter tasks
  `-status pending`
  `-status completed`

* Complete task
  `-complete <task_id>`

* Delete task
  `-delete <task_id>`

## Tech Stack

* Go (standard library only)
* `flag`
* `encoding/json`
* `os`

## Learning Outcomes

* Go syntax and basic types
* Structs and methods
* File I/O
* JSON marshaling/unmarshaling
* CLI argument parsing
* Error handling

