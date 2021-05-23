package commands

import (
	"errors"
	"github.com/urfave/cli/v2"
	"go.bnck.me/loki-auth-proxy/internal/config"
	"go.bnck.me/loki-auth-proxy/pkg/authenticator"
	"go.bnck.me/loki-auth-proxy/pkg/proxy"
)

func init() {
	RegisterCommand(&cli.Command{
		Name:    "listen",
		Aliases: []string{"l"},
		Usage:   "Starts the loki auth proxy",
		Action:  runListen,
	})
}

func runListen(c *cli.Context) error {
	cfg := config.Get()
	cfg.Load(c.String("config"))

	if len(cfg.Backends) == 0 {
		return errors.New("no backend server specified")
	}

	p := proxy.Proxy{
		Backends:      cfg.Backends,
		ListenAddress: cfg.HTTP.Listen,
		Authenticator: &authenticator.File{},
	}

	return p.Run()
}
