package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "cliProject",
		Usage: "This is a sample CLI project.",
		Action: func(c *cli.Context) error {
			logrus.Info("Hello World")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}
