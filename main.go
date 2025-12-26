package main

import (
    "github.com/urfave/cli/v3"
	"context"
	"log"
	"ns/cmd"
	"os"
)

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
                Name:    "lookup",
                Aliases: []string{"l"},
                Usage:   "custom tool to make dns checkup on domains",
                Action: cmd.Lookup,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "type",
						Value: "A",
						Usage: "custom tool to make dns checkup on domains",
					},
				},
            },
            {
                Name:    "scan",
                Aliases: []string{"s"},
                Usage:   "custom tool to make port checkup on domains",
				Action: cmd.ScanPort,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "port",
						Value: "80",
						Usage: "custom tool to make port checkup on domains/ips",
					},
				},
            },
		},

    }
    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }

}