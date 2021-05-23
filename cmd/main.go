package main

import (
	"fmt"
	"github.com/jpbede/loki-auth-proxy/internal/commands"
	"github.com/urfave/cli/v2"
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
		Version: fmt.Sprintf("%s-%s", version, commit),
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
