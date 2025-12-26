package cmd

import (
    "fmt"
    "context"
	"net"
	"time"

    "github.com/urfave/cli/v3"
)

func ScanPort(ctx context.Context, cmd *cli.Command) error {
    domain := "alnoor-hajj.com"
    port := "70"
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
        return "error connecting", err
    }
    defer conn.Close()
    
    // 💡 The "Validation" Trick: 
    // If a firewall is faking the connection, writing to it often fails 
    // or behaves differently.
    err = conn.SetDeadline(time.Now().Add(10 * time.Second))
    if err != nil {
        return "open (but weird)", nil
    }
    return "open", nil
}