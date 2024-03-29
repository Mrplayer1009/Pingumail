package main

import (
	"client"
	"os"
	"server"

	"github.com/urfave/cli/v2" // imports as package "cli"

	"golang.org/x/term"
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
				mails := client.Reload()

				var username = os.Getenv("pinguUserName")

				for _, mail := range mails {
					if mail.To == username {
						println("From", mail.From, ":", mail.Body)
					}
				}
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
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Manage the users",
			Subcommands: []*cli.Command{
				{
					Name:    "add",
					Usage:   "Add a user",
					Aliases: []string{"a"},
					Action: func(c *cli.Context) error {

						println("Adding user...")
						if c.NArg() != 1 {
							println("Usage: pingumail user add <username>")
							return nil
						}
						var userName = c.Args().Get(0)

						println("Enter password:")
						password, err := term.ReadPassword(int(os.Stdin.Fd()))
						if err != nil {
							println("Error reading password")
							return nil
						}

						server.AddUser(userName, string(password))
						return nil
					},
				},
				{
					Name:    "remove",
					Usage:   "Remove a user",
					Aliases: []string{"r"},
					Action: func(c *cli.Context) error {
						println("Removing user...")
						// RemoveUser()
						return nil
					},
				},
				{
					Name:    "show",
					Usage:   "Show the users",
					Aliases: []string{"s"},
					Action: func(c *cli.Context) error {
						println("Showing users...")
						// ShowUser()
						return nil
					},
				},
			},
		},
		{
			Name:    "login",
			Usage:   "Login as a user",
			Aliases: []string{"l"},
			Action: func(c *cli.Context) error {
				println("Logging in...")

				if c.NArg() != 1 {
					println("Usage: pingumail login <username>")
					return nil
				}

				server.Login(c.Args().Get(0))
				return nil
			},
		},
	}
	app.RunAndExitOnError()

}
