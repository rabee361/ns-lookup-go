package main

import (
    "fmt"
    "log"
    "os"
    "context"
	"net"
	"time"

    "github.com/urfave/cli/v3"
)

func main() {
    cmd := &cli.Command{
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:  "type",
                Value: "A",
                Usage: "custom tool to make dns checkup on domains",
            },
        },
        Action: func(ctx context.Context, cmd *cli.Command) error {
            domain := "alnoor-hajj.com"
            record_type := "A"
            if cmd.NArg() > 0 {
                domain = cmd.Args().Get(0)
            }
			if cmd.String("type") == "TXT" {
				record_type = "TXT"
			 }

			if cmd.String("type") == "NS" {
				record_type = "NS"
			 }
			if cmd.String("type") == "CNAME" {
				record_type = "CNAME"
			 }
			result ,err := dns_query(domain, "8.8.8.8", record_type)

			if err != nil {
				fmt.Print("Error")
				return err
			}
			fmt.Printf("the ip for %s is: %s", domain, result)
			return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}

func dns_query(domain, dnsServerIP string, record_type string) (string, error) {
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
		LocalAddr: nil, 
	}

	resolver := &net.Resolver{
		PreferGo: true, 
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.DialContext(ctx, "udp", dnsServerIP+":53") 
		},
	} 

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := ""
	var err error

	if record_type == "A" {
		result, err = QueryA(ctx, domain, resolver)
	} else if record_type == "CNAME" {
		result, err = QueryCNAME(ctx, domain, resolver)
	} else if record_type == "NS" {
		result, err = QueryNS(ctx, domain, resolver)
	} else if record_type == "TXT" {
		result, err = QueryTXT(ctx, domain, resolver)
	} else if err != nil {
		return "", fmt.Errorf("DNS query for %s failed: %w", domain, err)
	}
	return result, err
}

func QueryCNAME(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupCNAME(ctx, domain)
	if err != nil {
		return "", err
	}
	return result, err
}

func QueryA(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupIPAddr(ctx, domain)
	if err != nil {
		return "", err
	}
	return result[0].String(), err
}

func QueryTXT(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupTXT(ctx, domain)
	if err != nil {
		return "", err
	}
	return result[0], err
}

func QueryNS(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	
	result, err := resolver.LookupNS(ctx, domain)
	if err != nil {
		return "", err
	}
	return result[0].Host, err
}