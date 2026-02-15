package main

import "fmt"

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Priority  int    `json:"priority"`
}

func (t Task) String() string {
	status := "[ ]"
	if t.Completed {
		status = "[✓]"
	}
	return fmt.Sprintf("%s [ID:%d] [P:%d] %s", status, t.ID, t.Priority, t.Title)
}
