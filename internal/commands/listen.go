package commands

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"go.bnck.me/loki-auth-proxy/internal/config"
	intLog "go.bnck.me/loki-auth-proxy/internal/log"
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
	"go.bnck.me/loki-auth-proxy/pkg/proxy"
	"os"
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

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(intLog.DecodeLogLevel(cfg.Log.Level))

	if cfg.Backend == "" {
		return errors.New("no backend server specified")
	}

	authenticator, err := authenticators.GetAuthenticator(cfg.Authenticator.Name, cfg.Authenticator.Config)
	if err != nil {
		return err
	}

	p := proxy.Proxy{
		Backend:       cfg.Backend,
		Authenticator: authenticator,
	}

	log.Info().Msgf("Listening on %s", cfg.HTTP.Listen)
	return p.Logger(&log.Logger).Run(cfg.HTTP.Listen)
}
