package handler

import (
	"encoding/json"
	"os"
)

type CommandHandlerParam struct {
	StringFlags map[string]string
	IntFlags    map[string]int
}

type Task struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	Task      string `json:"task"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const filePath string = "./source/task.json"

// readFile
// reads the JSON file and unmarshals it into a slice of Task structs.
func readFile() ([]Task, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
