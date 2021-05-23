package commands

import "github.com/urfave/cli/v2"

var commands []*cli.Command

// Get returns all registered commands for the cli
func Get() []*cli.Command {
	return commands
}

// RegisterCommand adds a command to the registry
func RegisterCommand(command *cli.Command) bool {
	commands = append(commands, command)
	return true
}
