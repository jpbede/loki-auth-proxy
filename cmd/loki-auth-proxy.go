package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.bnck.me/loki-auth-proxy/internal/commands"
	"log"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := &cli.App{
		Name:    "loki-auth-proxy",
		Usage:   "Grafana Loki authentication proxy",
		Version: fmt.Sprintf("%s-%s published at %s", version, commit, date),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				DefaultText: "/etc/loki-auth-proxy.yaml",
			},
		},
	}

	// get app commands
	app.Commands = commands.Get()

	// run app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
