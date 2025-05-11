package handler

import "fmt"

// GetListTask
// retrieves and displays the list of tasks from the JSON file.
func GetListTask(flags *CommandHandlerParam) error {
	tasks, err := readFile()
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}
	filteredStatus := flags.StringFlags["status"]
	for key, task := range tasks {
		if key == 0 {
			fmt.Printf("%-5s %-20s %-20s %-20s %-20s\n", "ID", "Task", "Status", "Created At", "Updated At")
		}
		if filteredStatus == "" || task.Status == filteredStatus {
			fmt.Printf("%-5d %-20s %-20s %-20s %-20s\n", task.ID, task.Task, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
	return nil
}
