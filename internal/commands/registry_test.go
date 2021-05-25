package commands_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"go.bnck.me/loki-auth-proxy/internal/commands"
	"testing"
)

func TestGet(t *testing.T) {
	cmds := commands.Get()
	assert.NotNil(t, cmds)
	assert.Greater(t, len(cmds), 0)
}

func TestRegisterCommand(t *testing.T) {
	beforeCnt := len(commands.Get())
	commands.RegisterCommand(&cli.Command{
		Name: "test",
	})
	afterCnt := len(commands.Get())

	assert.Greater(t, afterCnt, beforeCnt)
}
