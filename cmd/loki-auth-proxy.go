package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.bnck.me/loki-auth-proxy/internal/commands"
	"log"
	"os"

	_ "go.bnck.me/loki-auth-proxy/pkg/authenticators/all"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := &cli.App{
		Name:     "loki-auth-proxy",
		Usage:    "Grafana Loki authentication proxy",
		Version:  fmt.Sprintf("%s-%s published at %s", version, commit, date),
		Commands: commands.Get(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				DefaultText: "/etc/loki-auth-proxy.yaml",
			},
		},
	}

	// run app
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
