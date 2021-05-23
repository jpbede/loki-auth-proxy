package commands

import (
	"errors"
	"github.com/jpbede/loki-auth-proxy/config"
	"github.com/jpbede/loki-auth-proxy/pkg/authenticator"
	"github.com/jpbede/loki-auth-proxy/pkg/proxy"
	"github.com/urfave/cli/v2"
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
		Config:        cfg,
		Authenticator: &authenticator.File{},
	}

	return p.Run()
}
