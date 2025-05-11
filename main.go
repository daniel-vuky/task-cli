package main

import (
	"fmt"
	"os"

	"github.com/daniel-vuky/task-cli/command"
)

func main() {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	listCommand := command.RegisterCommand()
	cmd, existed := listCommand[os.Args[1]]
	if !existed {
		PrintUsage()
		os.Exit(1)
	}
	err := command.ExecuteCommand(cmd, os.Args[2:])
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
}

// PrintUsage
// Print the usage message for the command line tool
func PrintUsage() {
	fmt.Println("expected 'add', 'update', 'delete', 'mark-in-progress', 'mark-done' or 'list' subcommands")
}
