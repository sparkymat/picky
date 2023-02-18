package main

import (
	"log"
	"os"

	"github.com/sparkymat/pidi/setup"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pidi",
		Usage: "a commandline key-value store",
		Commands: []*cli.Command{
			{
				Name:   "setup",
				Usage:  "configure pidi",
				Action: setup.Configure,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
