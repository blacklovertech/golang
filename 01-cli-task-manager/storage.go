package main

import (
	"encoding/json"
	"os"
)

func LoadTasks(path string) ([]Task, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []Task{}, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Handle empty file
	if len(b) == 0 {
		return []Task{}, nil
	}
	var tasks []Task
	err = json.Unmarshal(b, &tasks)
	return tasks, err
}

func SaveTasks(path string, tasks []Task) error {
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}
