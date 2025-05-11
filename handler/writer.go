package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// AddTask
// adds a new task to the JSON file.
func AddTask(flags *CommandHandlerParam) error {
	tasks, err := readFile()
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	newTask := Task{
		ID:        tasks[len(tasks)-1].ID + 1,
		Status:    "todo",
		Task:      flags.StringFlags["task"],
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	tasks = append(tasks, newTask)
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("Error marshalling tasks: %v", err)
	}
	err = os.WriteFile(filePath, tasksJSON, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	return nil
}

func UpdateTask(flags *CommandHandlerParam) error {
	tasks, err := readFile()
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	id := flags.IntFlags["id"]
	task := flags.StringFlags["task"]
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Task = task
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			break
		}
	}
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("Error marshalling tasks: %v", err)
	}
	err = os.WriteFile(filePath, tasksJSON, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	return nil
}

func DeleteTask(flags *CommandHandlerParam) error {
	tasks, err := readFile()
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	id := flags.IntFlags["id"]
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("Error marshalling tasks: %v", err)
	}
	err = os.WriteFile(filePath, tasksJSON, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	return nil
}

func MarkTaskInProgress(flags *CommandHandlerParam) error {
	tasks, err := readFile()
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	id := flags.IntFlags["id"]
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = "in_progress"
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			break
		}
	}
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("Error marshalling tasks: %v", err)
	}
	err = os.WriteFile(filePath, tasksJSON, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	return nil
}

func MarkTaskDone(flags *CommandHandlerParam) error {
	tasks, err := readFile()
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	id := flags.IntFlags["id"]
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = "done"
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			break
		}
	}
	tasksJSON, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("Error marshalling tasks: %v", err)
	}
	err = os.WriteFile(filePath, tasksJSON, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	return nil
}
