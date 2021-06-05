package listen

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"go.bnck.me/loki-auth-proxy/internal/commands"
	"go.bnck.me/loki-auth-proxy/internal/config"
	intLog "go.bnck.me/loki-auth-proxy/internal/log"
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
	"go.bnck.me/loki-auth-proxy/pkg/proxy"
	"os"
)

func init() {
	commands.RegisterCommand(&cli.Command{
		Name:    "listen",
		Aliases: []string{"l"},
		Usage:   "Starts the loki auth proxy",
		Action:  runListen,
	})
}

func runListen(c *cli.Context) error {
	cfg := config.Get()
	if err := cfg.Load(c.String("config")); err != nil {
		return err
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(intLog.DecodeLogLevel(cfg.Log.Level))

	if cfg.Backend.Distributor == "" {
		return errors.New("no distributor specified")
	}

	if cfg.Backend.QueryFrontend == "" {
		return errors.New("no query-frontend specified")
	}

	authenticator, err := authenticators.GetAuthenticator(cfg.Authenticator.Name, cfg.Authenticator.Config)
	if err != nil {
		return err
	}

	p := proxy.New(cfg.Backend.Distributor, cfg.Backend.QueryFrontend, cfg.Backend.Querier, authenticator)

	opts := []proxy.Option{proxy.WithLogger(&log.Logger)}
	if cfg.Prometheus {
		opts = append(opts, proxy.WithPrometheus())
	}

	log.Info().Msgf("Listening on %s", cfg.HTTP.Listen)
	return p.Run(cfg.HTTP.Listen, opts...)
}
