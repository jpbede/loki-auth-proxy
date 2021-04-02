package commands

import "github.com/urfave/cli/v2"

var commands []*cli.Command

func Get() []*cli.Command {
	return commands
}

func RegisterCommand(command *cli.Command) bool {
	commands = append(commands, command)
	return true
}
