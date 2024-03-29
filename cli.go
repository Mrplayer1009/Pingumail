package main

import (
	"client"
	"server"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

func CliStart() {

	app := cli.NewApp()
	app.Name = "pingumail"
	app.Usage = "A simple mail server"
	app.Commands = []*cli.Command{
		{
			Name:    "start",
			Aliases: []string{"run", "up"},
			Usage:   "Start the mail server",
			Action: func(c *cli.Context) error {
				println("Starting mail server...")
				server.Start()
				return nil
			},
		},
		{
			Name:    "reload",
			Aliases: []string{"r"},
			Usage:   "Reload the unread mails",
			Action: func(c *cli.Context) error {
				println("Reloading unread mails...")
				// Reload()
				return nil
			},
		},
		{
			Name:    "send",
			Aliases: []string{"s", "mail", "m"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "f",
					Aliases: []string{"from"},
					Usage:   "Sender of the mail",
				},
				&cli.StringFlag{
					Name:    "t",
					Aliases: []string{"to"},
					Usage:   "Receiver of the mail",
				},
				&cli.StringFlag{
					Name:    "b",
					Aliases: []string{"body"},
					Usage:   "Body of the mail",
				},
			},
			Usage: "Send a mail",
			Action: func(c *cli.Context) error {

				println("Sending mail...")

				from := c.String("from")
				to := c.String("to")
				body := c.String("body")

				client.SendMail(from, to, body)

				return nil
			},
		},
		{
			Name:    "history",
			Aliases: []string{"hist"},
			Usage:   "Show the history of mails",
			Action: func(c *cli.Context) error {
				println("Showing mail history...")
				// History()
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "Stop the mail server",
			Action: func(c *cli.Context) error {
				println("Stopping mail server...")
				// Stop()
				return nil
			},
		},
		{
			Name:  "status",
			Usage: "Show the status of the mail server",
			Action: func(c *cli.Context) error {
				println("Showing mail server status...")
				// Status()
				return nil
			},
		},
		{
			Name:    "version",
			Aliases: []string{"ver", "v"},
			Usage:   "Show the version of the mail server",
			Action: func(c *cli.Context) error {
				println("Showing mail server version...")
				// Version()
				return nil
			},
		},
		{
			Name:    "config",
			Aliases: []string{"conf"},
			Usage:   "Manage the mail server configuration",
			Subcommands: []*cli.Command{
				{
					Name:    "add",
					Usage:   "Add a configuration",
					Aliases: []string{"a"},
					Action: func(c *cli.Context) error {
						println("Adding configuration...")
						// AddConfig()
						return nil
					},
				},
				{
					Name:    "remove",
					Usage:   "Remove a configuration",
					Aliases: []string{"r"},
					Action: func(c *cli.Context) error {
						println("Removing configuration...")
						// RemoveConfig()
						return nil
					},
				},
				{
					Name:    "show",
					Usage:   "Show the configurations",
					Aliases: []string{"s"},
					Action: func(c *cli.Context) error {
						println("Showing configurations...")
						// ShowConfig()
						return nil
					},
				},
			},
		},
		{
			Name:  "env",
			Usage: "Manage the environment variables",
			Subcommands: []*cli.Command{
				{
					Name:    "add",
					Usage:   "Add an environment variable",
					Aliases: []string{"a"},
					Action: func(c *cli.Context) error {
						println("Adding environment variable...")
						// AddEnv()
						return nil
					},
				},
				{
					Name:    "remove",
					Usage:   "Remove an environment variable",
					Aliases: []string{"r"},
					Action: func(c *cli.Context) error {
						println("Removing environment variable...")
						// RemoveEnv()
						return nil
					},
				},
				{
					Name:    "show",
					Usage:   "Show the environment variables",
					Aliases: []string{"s"},
					Action: func(c *cli.Context) error {
						println("Showing environment variables...")
						// ShowEnv()
						return nil
					},
				},
			},
		},
	}
	app.RunAndExitOnError()

}
