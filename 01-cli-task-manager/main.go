package main

import (
	"flag"
	"fmt"
	"os"
)

const dataFile = "tasks.json"

func main() {
	add := flag.String("add", "", "add task")
	list := flag.Bool("list", false, "list tasks")
	complete := flag.Int("complete", -1, "complete task by id")
	del := flag.Int("delete", -1, "delete task by id")
	status := flag.String("status", "", "filter: pending|completed")
	priority := flag.Int("priority", 1, "priority level")
	flag.Parse()

	tasks, err := LoadTasks(dataFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		id := len(tasks) + 1
		tasks = append(tasks, Task{id, *add, false, *priority})
	case *complete > 0:
		for i := range tasks {
			if tasks[i].ID == *complete {
				tasks[i].Completed = true
			}
		}
	case *del > 0:
		var out []Task
		for _, t := range tasks {
			if t.ID != *del {
				out = append(out, t)
			}
		}
		tasks = out
	case *list:
		for _, t := range tasks {
			if *status == "completed" && !t.Completed {
				continue
			}
			if *status == "pending" && t.Completed {
				continue
			}
			fmt.Println(t)
		}
	}

	_ = SaveTasks(dataFile, tasks)
}
