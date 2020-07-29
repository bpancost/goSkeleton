package main

import (
	"github.com/urfave/cli"
	"goSkeleton/internal/logging"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "cliProject",
		Usage: "This is a sample CLI project.",
		Action: func(c *cli.Context) error {
			logging.Info("Hello World")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logging.Fatal(err)
	}
}
