package commands

import (
	"errors"
	"github.com/urfave/cli/v2"
	"go.bnck.me/loki-auth-proxy/internal/config"
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
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

	authenticator, err := authenticators.GetAuthenticator(cfg.Authenticator.Name, cfg.Authenticator.Config)
	if err != nil {
		return err
	}

	p := proxy.Proxy{
		Backends:      cfg.Backends,
		Authenticator: authenticator,
	}

	return p.Run(cfg.HTTP.Listen)
}
