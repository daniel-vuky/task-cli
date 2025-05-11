package command

import (
	"flag"
	"fmt"

	"github.com/daniel-vuky/task-cli/handler"
)

// ExecuteCommand
// Execute the command with the provided arguments
// and return an error if any.
func ExecuteCommand(command CommandConfig, args []string) error {
	cmd := flag.NewFlagSet(command.Name, flag.ExitOnError)

	// Set up string flags for the command
	listStringFlags := make(map[string]*string)
	for flagName, flagDefinition := range command.StringFlags {
		listStringFlags[flagName] = cmd.String(flagName, "", flagDefinition.Decription)
	}

	// Set up int flags for the command
	listIntFlags := make(map[string]*int)
	for flagName, flagDefinition := range command.IntFlags {
		listIntFlags[flagName] = cmd.Int(flagName, 0, flagDefinition.Decription)
	}

	// Parse the command line arguments
	cmd.Parse(args)

	// Set up list string value for the command
	listStringFlagValues := make(map[string]string)
	for flagName, flagDefinition := range command.StringFlags {
		value := *listStringFlags[flagName]
		if flagDefinition.Required && value == "" {
			cmd.Usage()
			return fmt.Errorf("missing required flag: %s", flagName)
		}
		listStringFlagValues[flagName] = value
	}

	listIntValues := make(map[string]int)
	for flagName, flagDefinition := range command.IntFlags {
		value := *listIntFlags[flagName]
		if flagDefinition.Required && value == 0 {
			cmd.Usage()
			return fmt.Errorf("missing required flag: %s", flagName)
		}
		listIntValues[flagName] = value
	}

	return command.Handler(&handler.CommandHandlerParam{
		StringFlags: listStringFlagValues,
		IntFlags:    listIntValues,
	})
}
