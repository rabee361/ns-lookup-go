package cmd

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
)

func ScanPort(ctx context.Context, cmd *cli.Command) error {
    domain := ""
    port := ""
    if cmd.NArg() > 0 {
        domain = cmd.Args().Get(0)
    }
    if cmd.NArg() > 1 {
        port = cmd.Args().Get(1)
    }

    result , _ := scan_port(domain, port)

    fmt.Printf("the port %s for %s is: %s", port, domain, result)
    return nil
}

func scan_port(domain, port string) (string, error) {
    conn, err :=  net.DialTimeout("tcp", domain + ":" + port, 10*time.Second)
    if err != nil {
        if strings.Contains(err.Error(), "timeout"){
            return "error timeout", err
        }
        return "error connecting", err
    }
    defer conn.Close()
    
    return "open", nil
}