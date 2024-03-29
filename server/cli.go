package server

import (
	"github.com/urfave/cli/v2" // imports as package "cli"
)

func CliStart() {
	app := cli.NewApp()
	app.Name = "pingumail"
	app.Usage = "A simple mail server"
	app.Commands = []*cli.Command{
		{
			Name:  "start",
			Usage: "Start the mail server",
			Action: func(c *cli.Context) error {
				println("Starting mail server...")
				Start()
				return nil
			},
		},
	}
	app.RunAndExitOnError()
}