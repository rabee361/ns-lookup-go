package main

import (
	"context"
	"log"
	"ns/cmd"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v3"
)

func main() {
	rootCmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name: "lookup",
				Aliases: []string{"l"},
				Usage:   color.BlueString("custom tool to make dns checkup on domains"),
				Action:  cmd.Lookup,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "type",
						Value: "A",
						Usage: "type of dns record (A, MX, NS, etc)",
					},
				},
			},
			{
				Name: "scan",
				Aliases: []string{"s"},
				Usage:   color.BlueString("custom tool to make port checkup on domains"),
				Action:  cmd.ScanPort,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "port",
						Value: "all",
						Usage: "port to check (default: all available)",
					},
				},
			},
			// {
			// 	Name: "http",
			// 	Usage:   color.BlueString("health check tool for a website"),
			// 	Action:  cmd.Check,
			// 	Flags: []cli.Flag{
			// 		&cli.StringFlag{
			// 			Name:  "url",
			// 			Value: "all",
			// 			Usage: "url to check (default: all available)",
			// 		},
			// 	},
			// },
			{
				Name: "ssl",
				Usage: color.BlueString("custom tool to return info on the ssl cert of a website"),
				Action: cmd.CheckSSL,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "domain",
						Value: "all",
						Usage: "domain to check (default: all available)",
					},
				},
			},
		},
	}
	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
