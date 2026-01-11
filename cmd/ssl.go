package cmd

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func CheckSSL(ctx context.Context, cmd *cli.Command) error {
	domain := "alnoor-hajj.com"
	recordType := "A"
	if cmd.NArg() > 0 {
		domain = cmd.Args().Get(0)
	}

	// Override recordType if flag is set
	if typeFlag := cmd.String("type"); typeFlag != "" {
		recordType = typeFlag
	}

	result, err := QueryDNS(domain, "8.8.8.8", recordType)
	if err != nil {
		return err
	}
	fmt.Printf("the ip for %s is: %s", domain, result)
	return nil
}
