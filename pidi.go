package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pidi",
		Usage: "a commandline key-value store",
		Commands: []*cli.Command{
			{
				Name:  "setup",
				Usage: "configure pidi",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("configure it!")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
