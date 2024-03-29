package main

import (
	"client"
	"os"
	"server"

	"github.com/joho/godotenv"
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

				to := c.String("to")
				body := c.String("body")

				client.SendMail(to, body)

				return nil
			},
		},
		{
			Name:    "version",
			Aliases: []string{"ver", "v"},
			Usage:   "Show the version of the mail server",
			Action: func(c *cli.Context) error {
				println("Pingumail Version :", os.Getenv("pinguVersion"))
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
					Name:    "show",
					Usage:   "Show the environment variables",
					Aliases: []string{"s"},
					Action: func(c *cli.Context) error {

						envMap, _ := godotenv.Read(".env")

						println("Showing environment variables...")
						println("Username", envMap["pinguUserName"])
						println("Pingumail IP", envMap["pinguServerIP"])
						println("Pingumail version", envMap["pingVersion"])
						return nil
					},
				},
				{
					Name:    "modify",
					Usage:   "Modify some environment variables",
					Aliases: []string{"m"},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "var",
							Usage: "environment variable to modify",
						},
						&cli.StringFlag{
							Name:  "value",
							Usage: "new value of the environment variable",
						},
					},
					Action: func(c *cli.Context) error {

						v := c.String("var")
						val := c.String("value")

						err := os.Setenv(v, val)
						if err != nil {
							println("Error setting environment variable", v, ":", val)
						}

						envList, _ := godotenv.Read(".env")
						envList[v] = val

						godotenv.Write(envList, ".env")
						_ = godotenv.Load(".env")
						return nil
					},
				},
			},
		},
	}
	app.RunAndExitOnError()

}
