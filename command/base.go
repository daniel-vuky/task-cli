package command

import "github.com/daniel-vuky/task-cli/handler"

type CommandConfig struct {
	Name        string
	Description string
	StringFlags map[string]FlagDefinition
	IntFlags    map[string]FlagDefinition
	Handler     func(flags *handler.CommandHandlerParam) error
}

type FlagDefinition struct {
	Name       string
	Decription string
	Required   bool
}

// RegisterCommand
// Initialize the command configuration
// and return a map of command names to their configurations.
func RegisterCommand() map[string]CommandConfig {
	commands := make(map[string]CommandConfig)

	addCommand := func(
		name, description string,
		stringFlags map[string]FlagDefinition,
		intFlags map[string]FlagDefinition,
		handler func(flags *handler.CommandHandlerParam) error,
	) {
		commands[name] = CommandConfig{
			Name:        name,
			Description: description,
			StringFlags: stringFlags,
			IntFlags:    intFlags,
			Handler:     handler,
		}
	}

	addCommand(
		"add",
		"Add a new task",
		map[string]FlagDefinition{
			"task": {Name: "task", Decription: "Task to add", Required: true},
		},
		nil,
		handler.AddTask,
	)

	addCommand(
		"update",
		"Update an existing task",
		map[string]FlagDefinition{
			"task": {Name: "task", Decription: "Task to update", Required: true},
		},
		map[string]FlagDefinition{
			"id": {Name: "id", Decription: "ID of the task to update", Required: true},
		},
		handler.UpdateTask,
	)

	addCommand(
		"delete",
		"Delete a task",
		nil,
		map[string]FlagDefinition{
			"id": {Name: "id", Decription: "ID of the task to delete", Required: true},
		},
		handler.DeleteTask,
	)

	addCommand(
		"mark-in-progress",
		"Mark a task as in progress",
		nil,
		map[string]FlagDefinition{
			"id": {Name: "id", Decription: "ID of the task to mark as in progress", Required: true},
		},
		handler.MarkTaskInProgress,
	)

	addCommand(
		"mark-done",
		"Mark a task as done",
		nil,
		map[string]FlagDefinition{
			"id": {Name: "id", Decription: "ID of the task to mark as done", Required: true},
		},
		handler.MarkTaskDone,
	)

	addCommand(
		"list",
		"List all tasks",
		map[string]FlagDefinition{
			"status": {Name: "status", Decription: "Filter tasks by status (all, todo, in-progress, done)", Required: false},
		},
		nil,
		handler.GetListTask,
	)

	return commands
}
